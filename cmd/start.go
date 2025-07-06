package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/FerzDevZ/ferzlmp/internal/config"
	"github.com/FerzDevZ/ferzlmp/internal/services"
	"github.com/FerzDevZ/ferzlmp/internal"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start Apache and MySQL services",
	Long:  `Starts Apache and MySQL using the configured or bundled binaries. Checks for port conflicts before starting.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig(filepath.Join("config", "ferzlmp.yaml"))
		if err != nil {
			color.Red("Failed to load config: %v", err)
			return
		}
		if internal.IsPortInUse(cfg.PortApache) {
			color.Red("Port %d is already in use. Please stop the conflicting service or use a different port.", cfg.PortApache)
			return
		}
		if internal.IsPortInUse(cfg.PortMySQL) {
			color.Red("Port %d is already in use. Please stop the conflicting service or use a different port.", cfg.PortMySQL)
			return
		}
		color.Cyan("[FerzLmp] Starting services...")
		// Check if Apache binary exists
		if _, err := os.Stat(cfg.ApachePath); os.IsNotExist(err) {
			color.Red("Apache binary not found at %s. Please install or configure the correct path.", cfg.ApachePath)
			return
		}
		// Check if MySQL binary exists
		if _, err := os.Stat(cfg.MySQLPath); os.IsNotExist(err) {
			color.Red("MySQL binary not found at %s. Please install or configure the correct path.", cfg.MySQLPath)
			return
		}
		// Start Apache
		color.Yellow("Starting Apache...")
		if err := services.StartApache(cfg.ApachePath); err != nil {
			color.Red("Failed to start Apache: %v", err)
			return
		}
		color.Green("Apache started successfully!")
		// Start MySQL
		color.Yellow("Starting MySQL...")
		if err := services.StartMySQL(cfg.MySQLPath); err != nil {
			color.Red("Failed to start MySQL: %v", err)
			return
		}
		color.Green("MySQL started successfully!")
		color.Green("All services started successfully!")
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
