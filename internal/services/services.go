package services

import (
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
