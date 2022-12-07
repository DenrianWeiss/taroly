package rpc

import (
	"encoding/json"
	"os"
)

var rpcMap = make(map[string]string)

func init() {
	res, v := os.LookupEnv("TAROLY_RPC_URL")
	if !v {
		panic("TAROLY_RPC_URL not set")
	}
	err := json.Unmarshal([]byte(res), &rpcMap)
	if err != nil {
		panic(err)
	}

}

func SetRpcMap(chain string, url string) {
	rpcMap[chain] = url
}

func GetRpcUrl(chain string) string {
	chainUrl, ok := rpcMap[chain]
	if !ok {
		return ""
	}
	return chainUrl
}

func ListChains() []string {
	var chains []string
	for k := range rpcMap {
		chains = append(chains, k)
	}
	return chains
}
