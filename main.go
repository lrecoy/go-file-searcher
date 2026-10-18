package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("Go File Searcher")
	fmt.Println("----------------")

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <directory>")
		return
	}

	directory := os.Args[1]

	fmt.Println("Directory to search:", directory)

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		if !info.IsDir() {
			fmt.Println(path)
		}

		return nil
	})

	if err != nil {
		fmt.Println("Eror searching directory:", err)
	}
}
