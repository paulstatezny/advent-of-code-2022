package main

import (
	"fmt"
)

type RuckSack struct {
	// compartment1 is stored as a map (effectively a set) to keep searching O(n)
	// instead of O(n^2)
	compartment1 map[rune]bool
	compartment2 []rune
}

func main() {
	file, scanner := LineScanner("3.txt")
	defer CloseFile(file)

	totalPriority := 0

	for scanner.Scan() {
		line := GetLine(scanner)
		ruckSack := ruckSack(line)
		commonItem := commonItem(ruckSack)
		priority := priority(commonItem)
		totalPriority += priority
	}

	fmt.Println("Sum of priorities:", totalPriority)
}

// Get a RuckSack from a line of text
func ruckSack(line []rune) RuckSack {
	compartmentSize := len(line) / 2
	var compartment1 []rune = line[:compartmentSize]
	compartment2 := line[compartmentSize:]
	compartment1Map := make(map[rune]bool)

	for _, item := range compartment1 {
		compartment1Map[item] = true
	}

	return RuckSack{compartment1Map, compartment2}
}

// Find the common item between the two compartments of a ruck sack
func commonItem(sack RuckSack) rune {
	for _, item := range sack.compartment2 {
		if sack.compartment1[item] == true {
			return item
		}
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
