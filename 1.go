package main

import (
	"strconv"
	"bufio"
	"fmt"
	"os"
	"sort"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	totals := []int{}
	current := 0

	readFile, err := os.Open("./1.txt")
	check(err)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		text := fileScanner.Text()

		if text == "" {
			totals = append(totals, current)
			current = 0
		} else {
			i, err := strconv.Atoi(text)
			check(err)
			current += i
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(totals)))

	fmt.Println("Sum of first 3:", totals[0] + totals[1] + totals[2])
}
