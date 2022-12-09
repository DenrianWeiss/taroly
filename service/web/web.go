package web

import (
	"encoding/base32"
	"encoding/json"
	"fmt"
	"github.com/DenrianWeiss/taroly/model"
	"github.com/DenrianWeiss/taroly/service/cache/user"
	"github.com/DenrianWeiss/taroly/service/db"
	"github.com/DenrianWeiss/taroly/utils/hmac"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func Serve() {
	server := gin.Default()
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	server.GET("/trace/:txId", func(c *gin.Context) {
		txId := c.Param("txId")
		result, err := db.Get(db.GetDb(), []byte("trace"+txId))
		if err != nil || result == nil || len(result) == 0 {
			c.Data(200, "text/plain", []byte("No trace result found for "+txId))
			return
		}
		c.Data(200, "text/html", []byte(fmt.Sprintf("<html><head><title>Trace result for:%s</title></head>"+
			"<body style=\"font: monospace;\">%s</body></html>", txId, result)))
	})
	server.Any("/rpc/:instance/:hash", RpcProxy)
	err := server.Run(":80")
	if err != nil {
		panic(err)
	}

}

func RpcProxy(c *gin.Context) {
	endPointInfo := c.Param("instance")
	sig := c.Param("hash")
	if !hmac.Validate(endPointInfo, sig) {
		log.Println("Invalid signature:", sig)
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid signature",
		})
		return
	}

	var result []byte
	// Decode Base 32
	_, err := base32.StdEncoding.Decode(result, []byte(endPointInfo))
	if err != nil {
		c.JSON(500, gin.H{
			"error":  "invalid endpoint",
			"detail": err.Error(),
		})
		return
	}
	rpcInfo := model.EndPoint{}
	err = json.Unmarshal(result, &rpcInfo)
	if err != nil {
		c.JSON(500, gin.H{
			"error":  "invalid endpoint",
			"detail": err.Error(),
		})
		return
	}
	// Get User ID and port
	s := user.GetUserStatus(rpcInfo.Uid)
	if strconv.Itoa(s.SimulationPort) != rpcInfo.Port {
		c.JSON(500, gin.H{
			"error":  "invalid endpoint",
			"detail": "port mismatch",
		})
		return
	}
	// Forward request to simulation
	req, err := http.NewRequest(c.Request.Method, "http://127.0.0.1:"+rpcInfo.Port, c.Request.Body)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.DataFromReader(resp.StatusCode, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, nil)
	return
}
