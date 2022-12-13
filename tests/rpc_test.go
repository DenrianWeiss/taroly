package tests

import (
	"encoding/base32"
	json2 "encoding/json"
	"github.com/DenrianWeiss/taroly/model"
	"github.com/DenrianWeiss/taroly/service/cache/user"
	"github.com/DenrianWeiss/taroly/service/eth_rpc"
	"github.com/DenrianWeiss/taroly/service/foundry/anvil"
	"github.com/DenrianWeiss/taroly/service/web"
	"github.com/DenrianWeiss/taroly/utils/hmac"
	"strconv"
	"testing"
	"time"
)

func TestJsonRpc(t *testing.T) {
	impersonate, err := eth_rpc.Impersonate("http://127.0.0.1:8545", "0x4ca6A5cc14216Eacb00a9E71638A38937003EC26")
	if err != nil {
		panic(err)
	}
	t.Log(impersonate)
}

func TestSetBalance(t *testing.T) {
	setBalance, err := eth_rpc.SetBalance("http://127.0.0.1:8545", "0x4ca6A5cc14216Eacb00a9E71638A38937003EC26", "0x10000")
	if err != nil {
		panic(err)
	}
	t.Log(setBalance)
}

func TestNewRpc(t *testing.T) {
	pid, port, _ := anvil.StartFork("eth", 0)
	user.SetUserSimulation("100", pid, port)
	json, _ := json2.Marshal(model.EndPoint{
		Uid:  strconv.Itoa(100),
		Port: strconv.Itoa(port),
	})
	encode := base32.StdEncoding.EncodeToString(json)
	hmacV := hmac.SignWithNonce(string(json))
	t.Log(encode)
	t.Log(hmacV)
	go web.Serve()
	time.Sleep(600 * time.Second)
}
