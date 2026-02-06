package repo

import (
	"fmt"
	"log"
	"os"
)

func CrumbRepoExists(path string) bool {
	entries, err := os.ReadDir(path)

	if err != nil {
		log.Fatalf("Could not read directory: %v\n", err)
	}

	fmt.Printf("Folders in %s:\n", path)

	for _, entry := range entries {
		if entry.IsDir() && entry.Name() == ".crumbs" {
			return true
		}
	}

	return false
}
