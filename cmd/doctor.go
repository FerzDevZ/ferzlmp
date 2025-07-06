package cmd

import (
	"fmt"
	"github.com/ferzdev/ferzlmp/internal"
	"github.com/ferzdev/ferzlmp/internal/config"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"os/exec"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Diagnose environment and check for issues",
	Long:  `Checks system ports, required binaries, permissions, and virtualhost config.",
	Run: func(cmd *cobra.Command, args []string) {
		color.Cyan("[FerzLmp] Running diagnostics...")
		cfg, err := config.LoadConfig(filepath.Join("config", "ferzlmp.yaml"))
		if err != nil {
			color.Red("Failed to load config: %v", err)
			return
		}
		// Check ports
		if internal.IsPortInUse(cfg.PortApache) {
			color.Red("Port %d (Apache) is in use!", cfg.PortApache)
		} else {
			color.Green("Port %d (Apache) is free.", cfg.PortApache)
		}
		if internal.IsPortInUse(cfg.PortMySQL) {
			color.Red("Port %d (MySQL) is in use!", cfg.PortMySQL)
		} else {
			color.Green("Port %d (MySQL) is free.", cfg.PortMySQL)
		}
		// Check binaries
		checkBin := func(name, path string) {
			if _, err := os.Stat(path); os.IsNotExist(err) {
				color.Red("%s binary not found at %s", name, path)
			} else {
				color.Green("%s binary found at %s", name, path)
			}
		}
		checkBin("Apache", cfg.ApachePath)
		checkBin("MySQL", cfg.MySQLPath)
		checkBin("PHP", cfg.PHPPath)
		// Check permissions (hosts file)
		hostsPath := "/etc/hosts"
		if os.Getenv("OS") == "Windows_NT" {
			hostsPath = filepath.Join(os.Getenv("SystemRoot"), "System32", "drivers", "etc", "hosts")
		}
		if f, err := os.OpenFile(hostsPath, os.O_WRONLY|os.O_APPEND, 0644); err != nil {
			color.Red("No write access to hosts file: %s", hostsPath)
		} else {
			color.Green("Write access to hosts file: %s", hostsPath)
			f.Close()
		}
		// Check vhost config dir
		vhostDir := filepath.Join("modules", "apache", "conf", "vhosts")
		if _, err := os.Stat(vhostDir); os.IsNotExist(err) {
			color.Red("Apache vhost config dir not found: %s", vhostDir)
		} else {
			color.Green("Apache vhost config dir found: %s", vhostDir)
		}
		// Dependency check
		checkDep := func(name, bin string) {
			_, err := exec.LookPath(bin)
			if err != nil {
				color.Red("Dependency missing: %s (%s)", name, bin)
			} else {
				color.Green("Dependency found: %s", bin)
			}
		}
		checkDep("Composer", "composer")
		checkDep("Curl", "curl")
		checkDep("Unzip", "unzip")
		checkDep("Tar", "tar")
		color.Cyan("[FerzLmp] Diagnostics complete.")
	},
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}
