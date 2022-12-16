package anvil

import (
	"errors"
	"github.com/DenrianWeiss/taroly/service/cache/fork"
	"github.com/DenrianWeiss/taroly/service/cache/rpc"
	"github.com/DenrianWeiss/taroly/service/env"
	"os"
	"os/exec"
	"strconv"
)

func StartFork(chainName string, blocknumber int64) (pid int, port int, err error) {
	// Get RPC url
	ethRpc := rpc.GetRpcUrl(chainName)
	if ethRpc == "" {
		return 0, 0, errors.New("chain not supported")
	} else {
		port := fork.GetPort()
		if port == 0 {
			return 0, 0, errors.New("no port available")
		}
		var cmd *exec.Cmd
		if blocknumber != 0 {
			cmd = exec.Command(env.GetAnvilPath(), "--steps-tracing", "--fork-url", ethRpc, "--fork-block-number", strconv.Itoa(int(blocknumber)), "-p", strconv.Itoa(port))
		} else {
			cmd = exec.Command(env.GetAnvilPath(), "--fork-url", ethRpc, "-p", strconv.Itoa(port))
		}
		err := cmd.Start()
		if err != nil {
			return 0, 0, err
		}
		fork.SetFork(cmd.Process.Pid, cmd)
		fork.SetPidPort(cmd.Process.Pid, port)
		return cmd.Process.Pid, port, nil
	}
}

func StopFork(pid int) error {
	cmd := fork.GetFork(pid)
	if cmd == nil {
		return errors.New("fork not found")
	} else {
		err := cmd.Process.Kill()
		if err != nil {
			if err == os.ErrProcessDone {
				fork.DeleteFork(pid)
				fork.ReturnPort(fork.GetPidPort(pid))
				return nil
			}
			return err
		}
		fork.DeleteFork(pid)
		fork.ReturnPort(fork.GetPidPort(pid))
		return nil
	}
}
