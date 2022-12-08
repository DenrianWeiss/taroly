package web

import (
	"fmt"
	"github.com/DenrianWeiss/taroly/service/db"
	"github.com/gin-gonic/gin"
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
	err := server.Run(":80")
	if err != nil {
		panic(err)
	}

}
