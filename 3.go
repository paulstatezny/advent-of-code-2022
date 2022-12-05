package main

import (
	"fmt"
)

type RuckSack1 struct {
	// compartment1 is stored as a map (effectively a set) to keep searching O(n)
	// instead of O(n^2)
	compartment1 map[rune]bool
	compartment2 []rune
}

type RuckSack2 struct {
	// stored as a map (effectively a set) to keep searching O(n) instead of O(n^2)
	items map[rune]bool
}

func part1Main() {
	file, scanner := LineScanner("3.txt")
	defer CloseFile(file)

	totalPriority := 0

	for scanner.Scan() {
		line := GetLine(scanner)
		ruckSack := ruckSack1(line)
		commonItem := commonItem1(ruckSack)
		priority := priority(commonItem)
		totalPriority += priority
	}

	fmt.Println("Sum of priorities:", totalPriority)
}

func main() {
	file, scanner := LineScanner("3.txt")
	defer CloseFile(file)

	totalPriority := 0
	group := []RuckSack2{}

	for scanner.Scan() {
		line := GetLine(scanner)
		ruckSack := ruckSack2(line)
		group = append(group, ruckSack)

		if len(group) == 3 {
			commonItem := commonItem2(group)
			totalPriority += priority(commonItem)
			group = make([]RuckSack2, 0)
		}
	}

	fmt.Println("Sum of priorities:", totalPriority)
}

// Get a RuckSack1 from a line of text
func ruckSack1(line []rune) RuckSack1 {
	compartmentSize := len(line) / 2
	var compartment1 []rune = line[:compartmentSize]
	compartment2 := line[compartmentSize:]
	compartment1Map := make(map[rune]bool)

	for _, item := range compartment1 {
		compartment1Map[item] = true
	}

	return RuckSack1{compartment1Map, compartment2}
}

// Get a RuckSack2 from a line of text
func ruckSack2(line []rune) RuckSack2 {
	items := make(map[rune]bool)

	for _, item := range line {
		items[item] = true
	}

	return RuckSack2{items}
}

// Find the common item between the two compartments of a ruck sack
func commonItem1(sack RuckSack1) rune {
	for _, item := range sack.compartment2 {
		if sack.compartment1[item] == true {
			return item
		}
	}

	return '0'
}

// Find the common item between 3 separate RuckSack2s
func commonItem2(sacks []RuckSack2) rune {
	matches := sacks[0].items

	for _, sack := range sacks[1:] {
		for index, _ := range matches {
			if sack.items[index] == false {
				delete(matches, index)
			}
		}
	}

	if len(matches) != 1 {
		panic("More than 1 common item between 3 RuckSacks")
	}

	for item, _ := range matches {
		return item
	}

	return '0'
}

// Return the priority of the given item
func priority(item rune) int {
	if item >= 'a' {
		return 1 + int(item) - int('a')
	} else {
		return 27 + int(item) - int('A')
	}
}
