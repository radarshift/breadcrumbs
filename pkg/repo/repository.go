package repo

import (
	"fmt"
	"os"
)

func CreateCrumbRepo(path string) {
	crumbpath := path + "/.crumbs/cont"

	err := os.MkdirAll(crumbpath, 0755)
	if err != nil {
		fmt.Println("Error creating .crumbs directory:", err)
		os.Exit(1)
	}
}
