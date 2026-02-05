package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "crumb",
	Short:   "breadcrumbs - repo-local waypoints for unfinished work",
	Version: "v0.1.0",
	Long:    `breadcrumbs - leave navigational “crumbs” in your repo to track unfinished work and jump back to it instantly`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "There was an error using breadcrumbs!\n%s\n", err)
		os.Exit(1)
	}
}
