package cmd

import (
	"fmt"
	"github.com/ferzdev/ferzlmp/internal/config"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

var useCmd = &cobra.Command{
	Use:   "use [module] [version]",
	Short: "Switch active PHP or MySQL version",
	Long:  `Sets the active version of PHP or MySQL globally or per project.",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		module := args[0]
		version := args[1]
		cfgPath := filepath.Join("config", "ferzlmp.yaml")
		cfg, err := config.LoadConfig(cfgPath)
		if err != nil {
			color.Red("Failed to load config: %v", err)
			return
		}
		var newPath string
		switch module {
		case "php":
			newPath = filepath.Join("modules", "php", version, "bin", "php")
			cfg.PHPPath = newPath
		case "mysql":
			newPath = filepath.Join("modules", "mysql", version, "bin", "mysqld")
			cfg.MySQLPath = newPath
		default:
			color.Red("Unknown module: %s. Supported: php, mysql", module)
			return
		}
		// Save updated config
		f, err := os.OpenFile(cfgPath, os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			color.Red("Failed to open config for writing: %v", err)
			return
		}
		defer f.Close()
		enc := yaml.NewEncoder(f)
		if err := enc.Encode(cfg); err != nil {
			color.Red("Failed to write config: %v", err)
			return
		}
		color.Green("Active %s version set to %s!", module, version)
	},
}

func init() {
	rootCmd.AddCommand(useCmd)
}
