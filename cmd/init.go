package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize FerzLmp folders and config",
	Long:  `Creates the default folder structure and config file for FerzLmp. Run this after install!`,
	Run: func(cmd *cobra.Command, args []string) {
		dirs := []string{"projects", "modules/apache/conf/vhosts", "modules/mysql", "modules/php", "config"}
		for _, d := range dirs {
			if err := os.MkdirAll(d, 0755); err != nil {
				color.Red("Failed to create %s: %v", d, err)
				return
			}
		}
		configPath := filepath.Join("config", "ferzlmp.yaml")
		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			defaultConfig := """apache_path: modules/apache/bin/httpd\nmysql_path: modules/mysql/bin/mysqld\nphp_path: modules/php/bin/php\nprojects_dir: projects\nvhost_dir: modules/apache/conf/vhosts\n"""
			if err := os.WriteFile(configPath, []byte(defaultConfig), 0644); err != nil {
				color.Red("Failed to write config: %v", err)
				return
			}
			color.Green("Created default config at %s", configPath)
		} else {
			color.Yellow("Config already exists at %s", configPath)
		}
		color.Green("FerzLmp folder structure initialized!")
		color.Cyan("Next: ferzlmp doctor && ferzlmp new [laravel|wordpress] [name]")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
