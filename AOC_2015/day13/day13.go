package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	attendees        []string
	possibleSeatings [][]string
	happyRules       = make(map[string]map[string]int)
)

func main() {
	file, err := os.Open("input.txt")
	defer file.Close()

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		happyRule := scanner.Text()

		var (
			person, personNextTo, feeling string
			amount                        int
		)

		if n, _ := fmt.Sscanf(happyRule, "%s would %s %d happiness units by sitting next to %s.", &person, &feeling, &amount, &personNextTo); n == 4 {
			if feeling == "lose" {
				amount = -1 * amount
			}
			if _, exists := happyRules[person]; !exists {
				happyRules[person] = make(map[string]int)
				attendees = append(attendees, person)
			}
			personNextTo = strings.ReplaceAll(personNextTo, ".", "")
			happyRules[person][personNextTo] = amount
		} else {
			panic(happyRule)
		}
	}

	result := findBestSeating()

	fmt.Println("Result", result)
}

func permutateSeatingArrangement(n int) {
	if n == 1 {
		seatingCopy := make([]string, len(attendees))
		copy(seatingCopy, attendees)
		possibleSeatings = append(possibleSeatings, seatingCopy)
	} else {
		for i := 0; i < n-1; i++ {
			permutateSeatingArrangement(n - 1)
			if n%2 == 0 {
				swap(attendees, i, n-1)
			} else {
				swap(attendees, 0, n-1)
			}
		}
		permutateSeatingArrangement(n - 1)
	}
}

func swap(seating []string, i, j int) {
	seating[i], seating[j] = seating[j], seating[i]
}

func evaluateHappiness(seating []string) int {
	total := 0
	for index := range seating {
		if index == len(seating)-1 {
			total += happyRules[seating[index]][seating[0]]
			total += happyRules[seating[0]][seating[index]]
		} else {
			total += happyRules[seating[index]][seating[index+1]]
			total += happyRules[seating[index+1]][seating[index]]
		}
	}

	return total
}

func findBestSeating() int {
	permutateSeatingArrangement(len(attendees))
	bestSeating := 0

	for _, seating := range possibleSeatings {
		value := evaluateHappiness(seating)

		if value > bestSeating {
			bestSeating = value
		}
	}

	return bestSeating
}
