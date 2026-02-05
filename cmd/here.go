package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var hereCmd = &cobra.Command{
	Use:     "here",
	Aliases: []string{"h"},
	Short:   "creates a crumb for the current directory",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("here called")
	},
}

func init() {
	rootCmd.AddCommand(hereCmd)
}
