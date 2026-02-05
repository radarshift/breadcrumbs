package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var renameCmd = &cobra.Command{
	Use:     "rename",
	Aliases: []string{"re", "r"},
	Short:   "did a misinput? use this! renames a crumb",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rename called")
	},
}

func init() {
	rootCmd.AddCommand(renameCmd)
}
