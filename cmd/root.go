package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ferzlmp",
	Short: "FerzLmp - CLI-based Web Server Manager (Apache, PHP, MySQL)",
	Long: `FerzLmp is a CLI tool to manage a local PHP development server.
It provides an all-in-one web environment (Apache, PHP, MySQL) that works without GUI.

EXAMPLES:
  ferzlmp start
  ferzlmp stop
  ferzlmp install php 8.2
  ferzlmp use php 8.2
  ferzlmp new laravel blog
  ferzlmp vhost add blog.test ./projects/blog
  ferzlmp doctor
`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
