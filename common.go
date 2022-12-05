package main

import (
	"bufio"
	"fmt"
	"os"
)

// Return a bufio.Scanner to iterate through lines of a text file
func LineScanner(path string) (*os.File, *bufio.Scanner) {
	file, err := os.Open(path)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	return file, scanner
}

func CloseFile(f *os.File) {
	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

// Get a single line of text as a []rune from a bufio.Scanner
func GetLine(scanner *bufio.Scanner) []rune {
	return []rune(scanner.Text())
}

