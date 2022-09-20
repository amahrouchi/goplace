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

	// Confirmation
	confirmed, err := io.Confirm()
	if err != nil {
		fmt.Printf("\nUser confirmation error: \"%v\"\n", err)
		return
	}

	if !confirmed {
		fmt.Println("\nAbort mission!!!")
		return
	}

	fmt.Println("\nLet's go!")

	// Read the file line by line
	// Replace the string
	// Print the resulted file
	// Ask if the user wants to print in a file
	// If yes print the file path given by the user
}
