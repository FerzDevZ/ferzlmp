package cmd

import (
	"github.com/FerzDevZ/ferzlmp/internal/config"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop Apache and MySQL services",
	Long:  `Stops Apache and MySQL services if running.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig(filepath.Join("config", "ferzlmp.yaml"))
		if err != nil {
			color.Red("Failed to load config: %v", err)
			return
		}
		color.Cyan("[FerzLmp] Stopping services...")
		stopService("apache", cfg.ApachePath)
		stopService("mysql", cfg.MySQLPath)
	},
}

func stopService(name, binPath string) {
	if runtime.GOOS == "windows" {
		// Windows: taskkill by image name
		exe := filepath.Base(binPath)
		if !strings.HasSuffix(exe, ".exe") {
			exe += ".exe"
		}
		cmd := exec.Command("taskkill", "/F", "/IM", exe)
		if err := cmd.Run(); err != nil {
			color.Red("Failed to stop %s: %v", name, err)
			return
		}
		color.Green("%s stopped!", name)
	} else {
		// Linux: pkill by binary name
		bin := filepath.Base(binPath)
		cmd := exec.Command("pkill", "-f", bin)
		if err := cmd.Run(); err != nil {
			color.Red("Failed to stop %s: %v", name, err)
			return
		}
		color.Green("%s stopped!", name)
	}
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
