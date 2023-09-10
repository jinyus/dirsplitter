package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/i582/cfmt/cmd/cfmt"
)

func SplitDir(dir string, maxSize int64) {
	tracker := make(map[int]int64)
	currentPart := 1
	filesMoved := 0
	failedOps := 0
	skippedFiles := 0

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		parent := filepath.Dir(path)
		if isPartDir(parent) {
			skippedFiles += 1
			return nil
		}

		tracker[currentPart] += info.Size()

		if tracker[currentPart] >= maxSize {
			currentPart += 1
		}

		partDir := filepath.Join(dir, "part"+strconv.Itoa(currentPart))

		os.MkdirAll(partDir, os.ModePerm)

		relativeFilePath, err := filepath.Rel(dir, path)

		if err != nil {
			cfmt.Printf("{{Failed to get relative path:}}::red %s\n", path)
			fmt.Println(err)
			failedOps += 1
			return nil
		}

		newPath := filepath.Join(partDir, relativeFilePath)

		err = os.Rename(path, newPath)

		if err != nil {
			cfmt.Printf("{{Failed to move file:}}::red %s\n", path)
			fmt.Println(err)
			failedOps += 1
			return nil
		}

		filesMoved += 1

		return nil
	})

	cfmt.Println("\n{{Results:}}::yellow")
	cfmt.Printf("{{Files moved:}}::green %d\n", filesMoved)
	cfmt.Printf("{{Skipped files in part directory:}}::cyan %d\n", skippedFiles)
	cfmt.Printf("{{Failed operations:}}::red %d\n\n", failedOps)

	for part, size := range tracker {
		cfmt.Printf("{{part%d:}}::magenta %dMB\n", part, size/1024/1024)
	}

	cfmt.Println("\n{{Tar Command:}}::yellow")
	if currentPart == 1 {
		fmt.Printf(`tar -cf "part1.tar" "part1"; done` + "\n")
	} else {
		fmt.Printf(`for n in {1..%d}; do tar -cf "part$n.tar" "part$n"; done`+"\n", currentPart)
	}
}
