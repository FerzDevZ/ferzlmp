package services

import (
	"fmt"
	"net"
	"os/exec"
	"runtime"
)

func StartApache(apachePath string) error {
	bin := apachePath
	if runtime.GOOS == "windows" {
		bin = apachePath + ".exe"
	}
	cmd := exec.Command(bin)
	return cmd.Start()
}

func StartMySQL(mysqlPath string) error {
	bin := mysqlPath
	if runtime.GOOS == "windows" {
		bin = mysqlPath + ".exe"
	}
	cmd := exec.Command(bin)
	return cmd.Start()
}

func IsPortInUse(port int) bool {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return true // Port is in use
	}
	ln.Close()
	return false // Port is free
}
