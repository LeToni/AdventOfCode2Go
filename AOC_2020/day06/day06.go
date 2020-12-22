package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type group struct {
	answers []string
}

// Day4 Advent of Code Challenge
func main() {
	// Open the file
	file, err := os.Open("input.txt")
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)

	groups := []group{}
	var currentGroup group
	for scanner.Scan() {
		text := scanner.Text()

		if len(text) == 0 {
			groups = append(groups, currentGroup)
			currentGroup.answers = nil
		} else {
			currentGroup.answers = append(currentGroup.answers, text)
		}

	}
	groups = append(groups, currentGroup)

	sum := 0
	sumAll := 0
	for _, g := range groups {
		sum = sum + groupAnsweredQuestions(g)
		sumAll = sumAll + questionsAnsweredByAllInGroup(g)
	}

	fmt.Println("Number of questions to which anyone answered 'yes': ", sum)
	fmt.Println("Number of questions answered with yes by all in group: ", sumAll)

}

func groupAnsweredQuestions(aGroup group) int {
	letters := make(map[rune]bool)

	for _, answer := range aGroup.answers {
		for _, c := range answer {
			letters[c] = true
		}
	}

	return len(letters)
}

func questionsAnsweredByAllInGroup(aGroup group) int {
	letters := make(map[rune]bool)

	for _, char := range aGroup.answers[0] {
		letters[char] = false
	}

	for i := 1; i < len(aGroup.answers); i++ {
		for char := range letters {
			if !strings.ContainsRune(aGroup.answers[i], char) {
				delete(letters, char)
			}
		}
	}

	return len(letters)
}
