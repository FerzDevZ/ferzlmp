package cmd

import (
	"fmt"
	"github.com/FerzDevZ/ferzlmp/internal/download"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"io"
	"os"
	"path/filepath"
	"runtime"
)

var installCmd = &cobra.Command{
	Use:   "install [module|all] [version]",
	Short: "Install PHP, MySQL, or Apache binaries",
	Long:  `Downloads and installs the specified module (php/mysql/apache) and version, or all at once.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		module := args[0]
		if module == "all" {
			versions := map[string]string{"php": "8.2", "mysql": "5.7", "apache": "2.4"}
			if len(args) > 1 {
				for i, m := range []string{"php", "mysql", "apache"} {
					if len(args) > i+1 {
						versions[m] = args[i+1]
					}
				}
			}
			for m, v := range versions {
				color.Cyan("Installing %s version %s...", m, v)
				if !installModule(m, v) {
					return
				}
			}
			color.Green("All modules installed!")
			return
		}
		if len(args) < 2 {
			color.Red("Usage: ferzlmp install [php|mysql|apache|all] [version]")
			return
		}
		if !installModule(module, args[1]) {
			return
		}
	},
}

func installModule(module, version string) bool {
	var url, destDir, localFile, ext, osTag string
	if runtime.GOOS == "windows" {
		ext = ".zip"
		osTag = "windows"
	} else {
		ext = ".tar.gz"
		osTag = "linux"
	}
	switch module {
	case "php":
		url = fmt.Sprintf("https://github.com/FerzDevZ/ferzlmp/releases/download/php-%s/php-%s-%s%s", version, version, osTag, ext)
		destDir = filepath.Join("modules", "php")
		localFile = filepath.Join("internal", "download", fmt.Sprintf("php-%s-%s%s", version, osTag, ext))
	case "mysql":
		url = fmt.Sprintf("https://github.com/FerzDevZ/ferzlmp/releases/download/mysql-%s/mysql-%s-%s%s", version, version, osTag, ext)
		destDir = filepath.Join("modules", "mysql")
		localFile = filepath.Join("internal", "download", fmt.Sprintf("mysql-%s-%s%s", version, osTag, ext))
	case "apache":
		url = fmt.Sprintf("https://github.com/FerzDevZ/ferzlmp/releases/download/apache-%s/apache-%s-%s%s", version, version, osTag, ext)
		destDir = filepath.Join("modules", "apache")
		localFile = filepath.Join("internal", "download", fmt.Sprintf("apache-%s-%s%s", version, osTag, ext))
	default:
		color.Red("Unknown module: %s. Supported: php, mysql, apache", module)
		return false
	}
	os.MkdirAll(destDir, 0755)
	archivePath := filepath.Join(destDir, module+"-"+version+ext)
	if _, err := os.Stat(localFile); err == nil {
		color.Yellow("Using local file %s...", localFile)
		if err := copyFile(localFile, archivePath); err != nil {
			color.Red("Failed to copy local file: %v", err)
			return false
		}
	} else {
		color.Yellow("Downloading from %s...", url)
		if err := download.DownloadFile(url, archivePath); err != nil {
			color.Red("Download failed: %v", err)
			return false
		}
	}
	color.Yellow("Extracting to %s...", destDir)
	err := download.Unzip(archivePath, destDir)
	if err != nil {
		color.Red("Extraction failed: %v", err)
		color.Yellow("Coba cek apakah file arsip valid, atau gunakan versi lain.")
		return false
	}
	color.Green("%s %s installed successfully!", module, version)
	return true
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, in)
	return err
}

func init() {
	rootCmd.AddCommand(installCmd)
}
