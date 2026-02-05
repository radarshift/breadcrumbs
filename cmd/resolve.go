package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var resolveCmd = &cobra.Command{
	Use:     "resolve",
	Aliases: []string{"res"},
	Short:   "gives the full path that a crumb points to",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("resolve called")
	},
}

func init() {
	rootCmd.AddCommand(resolveCmd)
}
