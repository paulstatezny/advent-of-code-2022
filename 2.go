package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
 * SEMANTIC TYPES
 *
 * To help make the code easier to reason about.
 */

type Choice int
type Outcome int

// A type to represent a Choice that a player makes in a Rock, Paper, Scissors game
const (
	Rock Choice = iota
	Paper
	Scissors
)

// A type to represent the Outcome of a Rock, Paper, Scissors game
const (
	Win Outcome = iota
	Tie
	Loss
)

/*
 * READING AND PARSING THE TEXT FILE
 */

// Deserialize a "Choice" rune into a Choice type
func parseChoice(r rune) Choice {
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

// Deserialize an "Outcome" rune into an Outcome type
func parseOutcome(r rune) Outcome {
	switch r {
	case 'X':
		return Loss
	case 'Y':
		return Tie
	case 'Z':
		return Win
	default:
		panic("invalid input")
	}
}

// Return a bufio.Scanner to iterate through lines of a text file
func lineScanner() (*os.File, *bufio.Scanner) {
	file, err := os.Open("2.txt")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	return file, scanner
}

func closeFile(f *os.File) {
	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

// Get a single line of text as a []rune from a bufio.Scanner
func getLine(scanner *bufio.Scanner) []rune {
	return []rune(scanner.Text())
}

/*
 * LOGIC UTILITIES
 *
 * Simple functions that return atomic facts or decisions about a scenario in a
 * Rock, Paper, Scissors game.
 */

// Given the choices of the self and opponent, return the Outcome
// from the perspective of the self. ("Did I win, lose, or tie?")
func outcomeOfRound(self Choice, opponent Choice) Outcome {
	if self == opponent {
		return Tie
	}

	if (self-opponent+3)%3 == 1 {
		return Win
	} else {
		return Loss
	}
}

// Given a Choice, return how many points should be tallied
func choiceScore(choice Choice) int {
	switch choice {
	case Rock:
		return 1

	case Paper:
		return 2

	default:
		return 3
	}
}

// Given an Outcome, return how many points should be tallied
func outcomeScore(outcome Outcome) int {
	switch outcome {
	case Win:
		return 6
	case Tie:
		return 3
	default:
		return 0
	}
}

// Given an opponent Choice and a desired Outcome, return
// the Choice that the self should make in order to achieve
// the Outcome
func selfChoice(opponent Choice, outcome Outcome) Choice {
	switch outcome {
	case Win:
		return (opponent + 1) % 3
	case Tie:
		return opponent
	default:
		return (opponent - 1 + 3) % 3
	}
}

/*
 * PROGRAM ENTRY POINTS
 *
 * For parts 1+2 of day 2 of Advent of Code 2022.
 */

// Part 1
func Part1Main() {
	file, scanner := lineScanner()
	defer closeFile(file)
	score := 0

	for scanner.Scan() {
		line := getLine(scanner)
		self := parseChoice(line[2])
		opponent := parseChoice(line[0])
		score += choiceScore(self) + outcomeScore(outcomeOfRound(self, opponent))
	}

	fmt.Println("Total score:", score)
}

// Part 2
func main2() {
	file, scanner := lineScanner()
	defer closeFile(file)
	score := 0

	for scanner.Scan() {
		line := getLine(scanner)
		opponent := parseChoice(line[0])
		outcome := parseOutcome(line[2])
		self := selfChoice(opponent, outcome)
		score += choiceScore(self) + outcomeScore(outcome)
	}

	fmt.Println("Total score:", score)
}
