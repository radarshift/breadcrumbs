package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var catCmd = &cobra.Command{
	Use:     "cat",
	Aliases: []string{"c"},
	Short:   "prints a crumb",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cat called")
	},
}

func init() {
	rootCmd.AddCommand(catCmd)
}
