package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var linkCmd = &cobra.Command{
	Use:     "link",
	Aliases: []string{"l"},
	Short:   "links a crumb to a given file",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("link called")
	},
}

func init() {
	rootCmd.AddCommand(linkCmd)
}
