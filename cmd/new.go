package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path/filepath"
)

var newCmd = &cobra.Command{
	Use:   "new [type] [name]",
	Short: "Create a new Laravel or WordPress project",
	Long:  `Downloads starter files and sets up vhost for Laravel or WordPress projects.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		ptype := args[0]
		name := args[1]
		projectPath := filepath.Join("projects", name)
		if _, err := os.Stat(projectPath); err == nil {
			color.Red("Project folder %s already exists!", projectPath)
			return
		}
		os.MkdirAll(projectPath, 0755)
		switch ptype {
		case "laravel":
			color.Cyan("Downloading Laravel installer...")
			cmd := exec.Command("composer", "create-project", "laravel/laravel", projectPath)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				color.Red("Failed to create Laravel project: %v", err)
				return
			}
		case "wordpress":
			color.Cyan("Downloading WordPress...")
			wpZip := filepath.Join(projectPath, "wordpress.zip")
			if err := downloadFile("https://wordpress.org/latest.zip", wpZip); err != nil {
				color.Red("Failed to download WordPress: %v", err)
				return
			}
			if err := unzipFile(wpZip, projectPath); err != nil {
				color.Red("Failed to extract WordPress: %v", err)
				return
			}
			os.Remove(wpZip)
		default:
			color.Red("Unknown project type: %s. Supported: laravel, wordpress", ptype)
			return
		}
		color.Green("Project %s created at %s!", name, projectPath)
		// Auto-create vhost
		domain := name + ".test"
		vhostPath := filepath.Join("modules", "apache", "conf", "vhosts")
		if err := os.MkdirAll(vhostPath, 0755); err == nil {
			if err := exec.Command(os.Args[0], "vhost", "add", domain, projectPath).Run(); err == nil {
				color.Green("Virtualhost %s added!", domain)
			}
		}
	},
}

func downloadFile(url, dest string) error {
	cmd := exec.Command("curl", "-L", "-o", dest, url)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func unzipFile(src, dest string) error {
	cmd := exec.Command("unzip", src, "-d", dest)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func init() {
	rootCmd.AddCommand(newCmd)
}
