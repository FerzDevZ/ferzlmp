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
		dirs := []string{
			"projects",
			"modules/apache/bin",
			"modules/apache/conf/vhosts",
			"modules/php/bin",
			"modules/mysql/bin",
			"config",
		}
		for _, d := range dirs {
			if err := os.MkdirAll(d, 0755); err != nil {
				color.Red("Failed to create %s: %v", d, err)
				return
			}
		}
		configPath := filepath.Join("config", "ferzlmp.yaml")
		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			defaultConfig := `apache_path: modules/apache/bin/httpd
mysql_path: modules/mysql/bin/mysqld
php_path: modules/php/bin/php
projects_dir: projects
vhost_dir: modules/apache/conf/vhosts
`
			if err := os.WriteFile(configPath, []byte(defaultConfig), 0644); err != nil {
				color.Red("Failed to write config: %v", err)
				return
			}
			color.Green("Created default config at %s", configPath)
		} else {
			color.Yellow("Config already exists at %s", configPath)
		}
		color.Green("FerzLmp folder structure initialized!")
		color.Cyan("\nSelanjutnya:\n- Letakkan/copy binary Apache, PHP, MySQL ke modules/[apache|php|mysql]/bin/\n- Atau, edit config/ferzlmp.yaml untuk pakai binary sistem\n- Buat project baru: ferzlmp new [laravel|wordpress] [namaproject]\n- Atau, pindahkan project lama ke folder projects/\n- Jalankan: ferzlmp start\n- Akses project di http://[namaproject].test\n")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
