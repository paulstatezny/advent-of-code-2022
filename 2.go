package main

import (
	"bufio"
	"fmt"
	"os"
)

type Choice int
type Outcome int

const (
	Rock Choice = iota
	Paper
	Scissors
)

const (
	Win Outcome = iota
	Tie
	Loss
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	readFile, err := os.Open("2.txt")

	check(err)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	score := 0

	for fileScanner.Scan() {
		text := []rune(fileScanner.Text())
		self := choice(text[2])
		opponent := choice(text[0])
		score += roundScore(self, opponent)
	}

	fmt.Println("Total score:", score)
}

func choice(r rune) Choice {
	switch r {
	case 'A', 'X':
		return Rock
	case 'B', 'Y':
		return Paper
	case 'C', 'Z':
		return Scissors
	default:
		panic("invalid input")
	}
}

func roundScore(self Choice, opponent Choice) int {
	score := 0

	switch self {
	case Rock:
		score += 1

	case Paper:
		score += 2

	case Scissors:
		score += 3
	}

	switch outcome(self, opponent) {
	case Win:
		score += 6
	case Tie:
		score += 3
	case Loss:
		score += 0
	}

	return score
}

func outcome(a Choice, b Choice) Outcome {
	if a == b {
		return Tie
	}

	if (a-b+3)%3 == 1 {
		return Win
	} else {
		return Loss
	}
}
