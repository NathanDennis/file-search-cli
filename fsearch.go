package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/manifoldco/promptui"
)

func main() {
	searchTerm := os.Args[1]
	var files []string
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if strings.Contains(path, searchTerm) {
			files = append(files, path)
		}
		return nil
	})
	if len(files) == 0 {
		fmt.Println("No files found")
		return
	}

	prompt := promptui.Select{
		Label: "Select file",
		Items: files,
	}

	_, file, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	err = exec.Command("goland", "--new-instance", file).Start()
	if err != nil {
		fmt.Printf("Failed to open file %v\n", err)
	}
}
