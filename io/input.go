package io

import (
	"fmt"
	"strings"
)

func GetUserInput(question string) (string, error) {
	fmt.Println("\n" + question)
	fmt.Print("-> ")

	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		return "", err
	}

	return input, nil
}

func Confirm() (bool, error) {
	var answer string
	for strings.ToLower(answer) != "y" && strings.ToLower(answer) != "n" {

		fmt.Println("\nDo you want to proceed to the replacement? (y/n)")
		_, err := fmt.Scanln(&answer)
		if err != nil {
			return false, err
		}

		if strings.ToLower(answer) == "n" {
			return false, nil
		}
	}

	return true, nil
}
