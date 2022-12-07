package eth_rpc

import (
	"encoding/hex"
	"fmt"
	"github.com/DenrianWeiss/taroly/utils/context"
	"github.com/DenrianWeiss/taroly/utils/hx"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	rpc2 "github.com/ethereum/go-ethereum/rpc"
	"github.com/shopspring/decimal"
	"math/big"
	"time"
)

func BalanceOf(rpc string, address string) (string, error) {
	client, err := ethclient.Dial(rpc)
	defer client.Close()
	if err != nil {
		return "", err
	}
	balance, err := client.BalanceAt(context.ContextWithTimeout(10*time.Second), common.HexToAddress(address), nil)
	deci := decimal.NewFromBigInt(balance, -18)
	return fmt.Sprintf("Balance of account %s is %s, raw %s", address, deci.String(), balance.String()), nil
}

func BlockNumber(rpc string) (int64, error) {
	client, err := ethclient.Dial(rpc)
	defer client.Close()
	if err != nil {
		return 0, err
	}
	blockNumber, err := client.BlockNumber(context.ContextWithTimeout(15 * time.Second))
	if err != nil {
		return 0, err
	}
	return int64(blockNumber), nil
}

func Impersonate(rpc string, address string) (string, error) {
	dial, err := rpc2.Dial(rpc)
	defer dial.Close()
	if err != nil {
		return "", err
	}
	result := ""
	err = dial.Call(&result, "anvil_impersonateAccount", address)
	if err != nil {
		return "", err
	}
	return result, nil
}

func SetBalance(rpc string, address string, amount string) (string, error) {
	dial, err := rpc2.Dial(rpc)
	defer dial.Close()
	if err != nil {
		return "", err
	}
	result := ""
	err = dial.Call(&result, "anvil_setBalance", address, amount)
	if err != nil {
		return "", err
	}
	return result, nil
}

func Call(rpc string, from, contract, payload, value string) (string, error) {
	client, err := ethclient.Dial(rpc)
	defer client.Close()
	if err != nil {
		return "", err
	}
	conc := common.HexToAddress(contract)
	v, _ := big.NewInt(0).SetString(value, 0)
	callContract, err := client.CallContract(context.ContextWithTimeout(15*time.Second), ethereum.CallMsg{
		From:  common.HexToAddress(from),
		To:    &conc,
		Value: v,
		Data:  hx.HexStringToBytes(payload),
	}, nil)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(callContract), nil
}

func Send(rpc string, from, contract, payload, value string) (string, error) {
	client, err := rpc2.Dial(rpc)
	defer client.Close()
	e, err := ethclient.Dial(rpc)
	defer e.Close()
	if err != nil {
		return "", err
	}
	result := ""
	v, _ := big.NewInt(0).SetString(value, 0)

	req := map[string]interface{}{}
	req["from"] = from
	req["to"] = contract
	req["data"] = payload
	req["value"] = (*hexutil.Big)(v)

	// Gas section
	conc := common.HexToAddress(contract)
	gas, err := e.EstimateGas(context.ContextWithTimeout(15*time.Second), ethereum.CallMsg{
		From:  common.HexToAddress(from),
		To:    &conc,
		Value: v,
		Data:  hx.HexStringToBytes(payload),
	})
	if err != nil {
		gas = 1000000
	} else {
		gas = gas * 2
	}

	req["gas"] = gas

	// GasPrice Section
	gasPrice, err := e.SuggestGasPrice(context.ContextWithTimeout(15 * time.Second))
	if err != nil {
		return "", err
	}

	req["gasPrice"] = (*hexutil.Big)(gasPrice)

	err = client.Call(&result, "eth_sendTransaction", req)
	if err != nil {
		return "", err
	}
	return result, nil
}
