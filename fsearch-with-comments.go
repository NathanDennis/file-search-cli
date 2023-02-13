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
	// Get the search term from the command line arguments
	searchTerm := os.Args[1]

	// Initialize a slice to store the matching files
	var files []string

	// Walk the file tree, starting from the current directory
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		// Skip directories
		if info.IsDir() {
			return nil
		}
		// If the file name contains the search term, add it to the list of files
		if strings.Contains(path, searchTerm) {
			files = append(files, path)
		}
		return nil
	})

	// If no files were found, print a message and return
	if len(files) == 0 {
		fmt.Println("No files found")
		return
	}

	// Create a prompt to select a file from the list of matching files
	prompt := promptui.Select{
		Label: "Select file",
		Items: files,
	}

	// Run the prompt and get the selected file
	_, file, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	// Launch the selected file in the Goland editor, with the --new-instance option
	err = exec.Command("open", "-a", "Goland", "--new-instance", file).Start()
	if err != nil {
		fmt.Printf("Failed to open file %v\n", err)
	}
}
