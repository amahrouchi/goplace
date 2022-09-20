package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("==============================")
	fmt.Println("======Welcome to GoPlace======")
	fmt.Println("==============================")
	fmt.Println("")
	fmt.Println("What is the file you want to work with?")
	fmt.Print("-> ")

	// Ask for file path
	var path string
	_, err := fmt.Scanln(&path)
	if err != nil {
		fmt.Println("Error while reading your input sorry...", err)
		return
	}

	// Check file existence
	_, err = os.Stat(path)
	if err != nil {
		fmt.Printf("File error: \"%v\"\n", err)
		return
	}

	fmt.Println(path)

	// Ask for the string to replace
	// Ask for the replacement string
	// Display a summary of what will be done
	// Read the file line by line
	// Replace the string
	// Print the resulted file
	// Ask if the user wants to print in a file
	// If yes print the file path given by the user
}
