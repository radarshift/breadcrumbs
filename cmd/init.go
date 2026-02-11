package cmd

import (
	"log"

	"github.com/radarshift/breadcrumbs/pkg/repo"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:     "init",
	Aliases: []string{"i"},
	Short:   "creates a new crumb repo",
	Run: func(cmd *cobra.Command, args []string) {
		/*
		* Want to first check if .crumb exists
		* If it does, we should just return and not do anything (maybe print a message saying "crumb repo already exists" ?)
		* If it doesn't exists, we should create it and print a message saying the repo was created successfully
		* For now (?) Dont care if subdirectories have crumbs/crumb repos
		* Printing error for now, in future, return fmt.Errorf and handle it in root.go
		 */

		// Process:
		// Check if .crumb exists
		// True: Print message, return
		// False: Create .crumb
		// Then Check created successfully
		// True: Print success message
		// False: Print error message

		dir, err := repo.GetUserDirPath()

		if err != nil {
			log.Printf("There was an error getting the user directory: %s\n", err)
		}

		// Check if .crumbs exists
		if !repo.CheckCrumbDirExists(dir) {
			// Create .crumbs
			err := repo.CreateCrumbDirectory(dir)

			// Check if error creating
			if err != nil {
				log.Printf("There was an error creating the crumb repo: %s\n", err)
			}

		} else { // .crumbs exists
			log.Println("Crumb repo already exists in this directory. No action taken.")
		}

	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
