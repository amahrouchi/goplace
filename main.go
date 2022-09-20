package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("==============================")
	fmt.Println("======Welcome to GoPlace======")
	fmt.Println("==============================")
	fmt.Println("\nWhat is the file you want to work with?")
	fmt.Print("-> ")

	// Ask for file path
	var path string
	_, err := fmt.Scanln(&path)
	if err != nil {
		fmt.Printf("\nInput error: \"%v\"\n", err)
		return
	}

	// Check file existence
	_, err = os.Stat(path)
	if err != nil {
		fmt.Printf("\nFile error: \"%v\"\n", err)
		return
	}

	// Ask for the string to replace
	fmt.Println("\nWhat is the string you want to replace in the file?")
	fmt.Print("-> ")
	var toReplace string
	_, err = fmt.Scanln(&toReplace)
	if err != nil {
		fmt.Printf("\nInput error: \"%v\"\n", err)
		return
	}

	// Ask for the replacement string
	fmt.Println("\nWhat is the replacement string that you want?")
	fmt.Print("-> ")
	var replacement string
	_, err = fmt.Scanln(&replacement)
	if err != nil {
		fmt.Printf("\nInput error: \"%v\"\n", err)
		return
	}

	// Display a summary of what will be done
	fmt.Println("\n=========== Summary ===========")
	fmt.Println("File: ", path)
	fmt.Println("Replacing: ", toReplace)
	fmt.Println("By: ", replacement)

	// Confirmation
	var answer string
	for strings.ToLower(answer) != "y" && strings.ToLower(answer) != "n" {
		fmt.Println("\nDo you want to proceed to the replacement? (y/n)")
		_, err = fmt.Scanln(&answer)
		if err != nil {
			fmt.Printf("\nInput error: \"%v\"\n", err)
			return
		}

		if strings.ToLower(answer) == "n" {
			fmt.Println("\nAbort mission!!!")
			return
		}
	}

	fmt.Println("\nLet's go!")

	// Read the file line by line
	// Replace the string
	// Print the resulted file
	// Ask if the user wants to print in a file
	// If yes print the file path given by the user
}
