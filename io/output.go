package io

import "fmt"

func DisplaySummary(path, toReplace, replacement string) {
	fmt.Println("\n=========== Summary ===========")
	fmt.Println("File: ", path)
	fmt.Println("Replacing: ", toReplace)
	fmt.Println("By: ", replacement)
}
