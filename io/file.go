package io

import (
	"bufio"
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
