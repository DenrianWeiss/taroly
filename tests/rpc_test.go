package tests

import (
	"github.com/DenrianWeiss/taroly/service/eth_rpc"
	"testing"
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
