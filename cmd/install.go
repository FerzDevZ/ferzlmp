package cmd

import (
	"fmt"
	"github.com/FerzDevZ/ferzlmp/internal/download"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var installCmd = &cobra.Command{
	Use:   "install [module] [version]",
	Short: "Install PHP, MySQL, or Apache binaries",
	Long:  `Downloads and installs the specified module (php/mysql/apache) and version.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		module := args[0]
		version := args[1]
		color.Cyan("Installing %s version %s...", module, version)
		var url, destDir string
		switch module {
		case "php":
			url = fmt.Sprintf("https://github.com/FerzDevZ/ferzlmp/releases/download/php-%s/php-%s.zip", version, version)
			destDir = filepath.Join("modules", "php")
		case "mysql":
			url = fmt.Sprintf("https://github.com/FerzDevZ/ferzlmp/releases/download/mysql-%s/mysql-%s.zip", version, version)
			destDir = filepath.Join("modules", "mysql")
		case "apache":
			url = fmt.Sprintf("https://github.com/FerzDevZ/ferzlmp/releases/download/apache-%s/apache-%s.zip", version, version)
			destDir = filepath.Join("modules", "apache")
		default:
			color.Red("Unknown module: %s. Supported: php, mysql, apache", module)
			return
		}
		os.MkdirAll(destDir, 0755)
		archivePath := filepath.Join(destDir, module+"-"+version+".zip")
		color.Yellow("Downloading from %s...", url)
		if err := download.DownloadFile(url, archivePath); err != nil {
			color.Red("Download failed: %v", err)
			return
		}
		color.Yellow("Extracting to %s...", destDir)
		if err := download.Unzip(archivePath, destDir); err != nil {
			color.Red("Extraction failed: %v", err)
			return
		}
		color.Green("%s %s installed successfully!", module, version)
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
