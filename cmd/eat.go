package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var eatCmd = &cobra.Command{
	Use:     "eat",
	Aliases: []string{"delete", "e", "del", "rm", "remove"},
	Short:   "eat (remove/delete) a crumb",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("eat called")
	},
}

func init() {
	rootCmd.AddCommand(eatCmd)
}
