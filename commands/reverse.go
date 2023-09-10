package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/i582/cfmt/cmd/cfmt"
)

func ReverseSplitDir(dir string) {
	foldersToRemove := make(map[string]bool)

	files, err := os.ReadDir(dir)

	if err != nil {
		cfmt.Printf("{{Failed to read directory:}}::red %s\n", dir)
		fmt.Println(err)
		os.Exit(1)
	}

	for _, file := range files {
		if file.IsDir() && isPartDir(file.Name()) {

			fullPath := filepath.Join(dir, file.Name())

			err := filepath.Walk(fullPath, func(path string, f os.FileInfo, err error) error {
				if f.IsDir() {
					return nil
				}

				relativeFilePath, err := filepath.Rel(fullPath, path)

				if err != nil {
					cfmt.Printf("{{Failed to get relative path:}}::red %s\n", path)
					fmt.Println(err)
					return err
				}

				newPath := filepath.Join(dir, relativeFilePath)

				err = os.Rename(path, newPath)

				if err != nil {
					cfmt.Printf("{{Failed to move file:}}::red %s\n", path)
					fmt.Println(err)
					return err
				}

				foldersToRemove[fullPath] = true

				return nil
			})

			if err != nil {
				cfmt.Printf("{{Failed to walk directory:}}::red %s\n", fullPath)
				fmt.Println(err)
				os.Exit(1)
			}

		}
	}

	cfmt.Printf("{{Done:}}::green Reverse split %d parts.\n", len(foldersToRemove))

	for folder := range foldersToRemove {
		os.RemoveAll(folder)
	}
}
