package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Welcome message
	fmt.Println("Welcome to the File Renamer!")
	fmt.Println("This tool will help you rename files in a directory by replacing characters in their names.")

	// Prompt for directory
	fmt.Print("Enter the directory path where files are located: ")
	directory, _ := reader.ReadString('\n')
	directory = strings.ReplaceAll(directory, "'", "")
	directory = strings.TrimSpace(directory)

	// Prompt for the character to replace
	fmt.Print("What's the first character you'd like to replace (you can include spaces): ")
	charToReplace, _ := reader.ReadString('\n')
	charToReplace = strings.TrimSuffix(charToReplace, "\n")

	// Prompt for the replacement character
	fmt.Print("What do you want it to be replaced with (you can include spaces): ")
	replacementChar, _ := reader.ReadString('\n')
	replacementChar = strings.TrimSuffix(replacementChar, "\n")

	// Prompt for file type/extension
	fmt.Print("What file extension to target (use * for all files): ")
	fileExtension, _ := reader.ReadString('\n')
	fileExtension = strings.TrimSpace(fileExtension)

	// Prompt for trimming the first x characters
	fmt.Print("Would you like to trim the first x characters from the file names? (yes/no): ")
	trimResponse, _ := reader.ReadString('\n')
	trimResponse = strings.TrimSpace(strings.ToLower(trimResponse))

	var trimLength int
	if trimResponse == "yes" {
		fmt.Print("How many characters would you like to trim from the start of the file names? ")
		trimLengthStr, _ := reader.ReadString('\n')
		trimLengthStr = strings.TrimSpace(trimLengthStr)

		trimLength, _ = strconv.Atoi(trimLengthStr) // Convert to integer
	}

	// Walk through the directory
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// If the file matches the given extension (or all files if '*')
		if !info.IsDir() && (fileExtension == "*" || strings.HasSuffix(info.Name(), fileExtension)) {
			originalName := info.Name()

			// Optionally trim the first x characters
			if trimLength > 0 && len(originalName) > trimLength {
				originalName = originalName[trimLength:]
			}

			// Replace the character in the file name
			newName := strings.ReplaceAll(originalName, charToReplace, replacementChar)

			if newName != info.Name() {
				// Construct the new file path
				newPath := filepath.Join(filepath.Dir(path), newName)

				// Rename the file
				err := os.Rename(path, newPath)
				if err != nil {
					return err
				}
				fmt.Printf("Renamed: %s -> %s\n", path, newPath)
			}
		}
		return nil
	})

	// Handle errors in walking the directory
	if err != nil {
		fmt.Printf("Error walking through the directory: %v\n", err)
	} else {
		fmt.Println("File renaming completed successfully.")
	}
}
