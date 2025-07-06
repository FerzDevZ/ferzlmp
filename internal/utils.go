package internal

import (
	"fmt"
	"os/exec"
	"runtime"
)

func IsPortInUse(port int) bool {
	// Simple TCP check for port usage
	address := ":" + fmt.Sprint(port)
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("netstat", "-ano")
	} else {
		cmd = exec.Command("lsof", "-i", address)
	}
	out, err := cmd.Output()
	return err == nil && len(out) > 0
}
