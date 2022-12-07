package proc

import (
	"os/exec"
	"time"
)

func KillAfter(sec time.Duration, process *exec.Cmd) {
	time.Sleep(sec * time.Second)
	_ = process.Process.Kill()
}
