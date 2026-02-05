package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var fCmd = &cobra.Command{
	Use:     "f",
	Aliases: []string{"finish", "followed"},
	Short:   "sets a breadcrumb as followed (finished/completed) if applicable",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("f called")
	},
}

func init() {
	rootCmd.AddCommand(fCmd)
}
