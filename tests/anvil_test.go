package tests

import (
	"github.com/DenrianWeiss/taroly/service/cache/rpc"
	"github.com/DenrianWeiss/taroly/service/foundry/anvil"
	"github.com/DenrianWeiss/taroly/service/foundry/cast"
	"testing"
)

func TestCastEncode(t *testing.T) {
	r := cast.EncodeCall("transfer(address,uint256)", "0x4ca6A5cc14216Eacb00a9E71638A38937003EC26", "0")
	t.Log(r)
}

func TestCastDecode(t *testing.T) {
	r := cast.DecodeCall("0xa9059cbb0000000000000000000000004ca6a5cc14216eacb00a9e71638a38937003ec260000000000000000000000000000000000000000000000000000000000000000")
	t.Log(r)
}

func TestCast4Byte(t *testing.T) {
	r := cast.SigCall("0xa9059cbb")
	t.Log(r)
}

func TestCastRun(t *testing.T) {
	r := cast.RunCall("https://rpc.ankr.com/eth", "0x3350ea2b4c77db04db33a909bb1196ff4aa171200cbe413d5de5c908602fb6d5")
	t.Log(r)
}

func TestStartAnvil(t *testing.T) {
	rpc.SetRpcMap("eth", "https://rpc.ankr.com/eth")
	fork, i, err := anvil.StartFork("eth", 0)
	if err != nil {
		t.Error(err)
	}
	t.Log(fork, i)
}
