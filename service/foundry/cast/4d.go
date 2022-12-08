package cast

import (
	"github.com/DenrianWeiss/taroly/service/env"
	"github.com/DenrianWeiss/taroly/utils/proc"
	"os/exec"
	"strings"
)

func EncodeCall(sig string, args ...string) string {
	p := []string{"calldata", sig}
	if len(args) > 0 {
		p = append(p, args...)
	}
	cmd := exec.Command(env.GetCastPath(), p...)
	go proc.KillAfter(10, cmd)
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	r := strings.TrimRight(string(output), "\n\r")
	return r
}

func DecodeCall(payload string) string {
	cmd := exec.Command(env.GetCastPath(), "4byte-decode", payload)
	go proc.KillAfter(10, cmd)
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	return string(output)
}

func SigCall(payload string) string {
	cmd := exec.Command(env.GetCastPath(), "4byte", payload)
	go proc.KillAfter(10, cmd)
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	return string(output)
}

func RunCall(rpc, txId string) string {
	cmd := exec.Command(env.GetCastPath(), "run", "--rpc-url", rpc, txId)
	go proc.KillAfter(480, cmd)
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	return strings.TrimLeft(strings.TrimPrefix(string(output), "Executing previous transactions from the block."), "\r\n")
}
