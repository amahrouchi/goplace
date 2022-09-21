package io

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func FindReplaceFile(src, old, new string) (newLines string, occ int, lines []int, err error) {
	// Open the file to read
	file, err := os.Open(src)
	if err != nil {
		return "", 0, []int{}, err
	}
	defer file.Close()

	// Go through the file and replace the string
	var totalCount int
	currLine := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currLine++

		// Get the curr line
		line := scanner.Text()

		// Replace occurrences
		newLines += processLine(line, old, new) + "\n"

		// Count the nb of occurrences
		count := strings.Count(line, old)
		if count == 0 {
			continue
		}

		totalCount += count
		lines = append(lines, currLine)
	}

	return newLines, totalCount, lines, nil
}

func processLine(line, old, new string) string {
	return strings.ReplaceAll(line, old, new)
}

func WriteFile(path, content string) error {
	// Check file existence
	_, err := os.Stat(path)

	// Create/open the file
	var f *os.File
	if err == nil {
		f, err = os.OpenFile(path, os.O_WRONLY, 0666)
		defer f.Close()
		if err != nil {
			return err
		}
	} else {
		f, err = os.Create(path)
		defer f.Close()
		if err != nil {
			return err
		}

		fmt.Printf("\nCreating the file: \"%v\"...\n", path)
	}

	_, err = f.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}
