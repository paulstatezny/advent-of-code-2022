package main

import (
	"strconv"
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	max := 0
	current := 0

	readFile, err := os.Open("./1.txt")
	check(err)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		text := fileScanner.Text()

		if text == "" {
			current = 0
		} else {
			i, err := strconv.Atoi(text)
			check(err)
			current += i

			if current > max {
				max = current
			}
		}
	}

	fmt.Println("Max:", max)
}
