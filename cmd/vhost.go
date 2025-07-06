package cmd

import (
	"fmt"
	"github.com/FerzDevZ/ferzlmp/internal/config"
	"github.com/FerzDevZ/ferzlmp/internal/vhost"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var vhostCmd = &cobra.Command{
	Use:   "vhost [add|remove] [domain] [path]",
	Short: "Manage Apache virtualhosts",
	Long:  `Adds or removes Apache virtualhosts and updates the local hosts file.",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		action := args[0]
		domain := args[1]
		cfg, err := config.LoadConfig(filepath.Join("config", "ferzlmp.yaml"))
		if err != nil {
			color.Red("Failed to load config: %v", err)
			return
		}
		apacheConfPath := filepath.Join("modules", "apache", "conf", "vhosts")
		if err := os.MkdirAll(apacheConfPath, 0755); err != nil {
			color.Red("Failed to create vhost config dir: %v", err)
			return
		}
		switch action {
		case "add":
			if len(args) < 3 {
				color.Red("Usage: ferzlmp vhost add <domain> <project_path>")
				return
			}
			projectPath := args[2]
			if err := vhost.AddVHost(domain, projectPath, apacheConfPath); err != nil {
				color.Red("Failed to add vhost: %v", err)
				return
			}
			if err := addHostEntry(domain); err != nil {
				color.Red("Failed to update hosts file: %v", err)
				return
			}
			color.Green("Virtualhost %s added for %s!", domain, projectPath)
		case "remove":
			if err := vhost.RemoveVHost(domain, apacheConfPath); err != nil {
				color.Red("Failed to remove vhost: %v", err)
				return
			}
			if err := removeHostEntry(domain); err != nil {
				color.Red("Failed to update hosts file: %v", err)
				return
			}
			color.Green("Virtualhost %s removed!", domain)
		default:
			color.Red("Unknown vhost action: %s. Use add or remove.", action)
		}
	},
}

func addHostEntry(domain string) error {
	hostsPath := "/etc/hosts"
	if os.Getenv("OS") == "Windows_NT" {
		hostsPath = filepath.Join(os.Getenv("SystemRoot"), "System32", "drivers", "etc", "hosts")
	}
	f, err := os.OpenFile(hostsPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	entry := fmt.Sprintf("127.0.0.1\t%s\n", domain)
	_, err = f.WriteString(entry)
	return err
}

func removeHostEntry(domain string) error {
	hostsPath := "/etc/hosts"
	if os.Getenv("OS") == "Windows_NT" {
		hostsPath = filepath.Join(os.Getenv("SystemRoot"), "System32", "drivers", "etc", "hosts")
	}
	data, err := os.ReadFile(hostsPath)
	if err != nil {
		return err
	}
	lines := []string{}
	for _, line := range splitLines(string(data)) {
		if !containsDomain(line, domain) {
			lines = append(lines, line)
		}
	}
	return os.WriteFile(hostsPath, []byte(joinLines(lines)), 0644)
}

func splitLines(s string) []string {
	return []string{os.ExpandEnv(s)}
}

func joinLines(lines []string) string {
	return fmt.Sprintln(lines)
}

func containsDomain(line, domain string) bool {
	return len(line) > 0 && (line == fmt.Sprintf("127.0.0.1\t%s", domain) || line == fmt.Sprintf("127.0.0.1 %s", domain))
}

func init() {
	rootCmd.AddCommand(vhostCmd)
}
