package main

import (
	"fmt"
	"os"
	"tests/goplace/io"
)

func main() {
	fmt.Println("==============================")
	fmt.Println("======Welcome to GoPlace======")
	fmt.Println("==============================")

	// Ask for file path
	path, err := io.GetUserInput("What is the file you want to work with?")
	if err != nil {
		fmt.Printf("\nUser input error: \"%v\"\n", err)
		return
	}

	// Check file existence
	_, err = os.Stat(path)
	if err != nil {
		fmt.Printf("\nFile error: \"%v\"\n", err)
		return
	}

	// Ask for the string to replace
	toReplace, err := io.GetUserInput("What is the string you want to replace in the file?")
	if err != nil {
		fmt.Printf("\nUser input error: \"%v\"\n", err)
		return
	}

	// Ask for the replacement string
	replacement, err := io.GetUserInput("What is the replacement string that you want?")
	if err != nil {
		fmt.Printf("\nUser input error: \"%v\"\n", err)
		return
	}

	// Display a summary of what will be done
	io.DisplaySummary(path, toReplace, replacement)

	// Replace
	newLines, occ, lines, err := io.FindReplaceFile(path, toReplace, replacement)
	if err != nil {
		fmt.Printf("\nError while replacing: \"%v\"\n", err)
		return
	}

	fmt.Printf("%d occurrences of \"%s\" where replaced by \"%s\"\n", occ, toReplace, replacement)
	fmt.Printf("Occurrences find on lines: %v\n", lines)

	confirmed, err := io.Confirm("Do you want see the new content?")
	if confirmed {
		fmt.Println(newLines)
	}

	// Write to file
	confirmed, err = io.Confirm("Do you write the content in a new file?")
	if err != nil {
		fmt.Printf("\nUser confirmation error: \"%v\"\n", err)
		return
	}

	if confirmed {
		newPath, err := io.GetUserInput("What is the file you want to write in?")
		if err != nil {
			fmt.Printf("\nUser input error: \"%v\"\n", err)
			return
		}

		err = io.WriteFile(newPath, newLines)
		if err != nil {
			fmt.Printf("\nWriting output error: \"%v\"\n", err)
			return
		}
	}
}
