package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Filter is a filter function applied to a single record
type Filter func(string) bool

// ApplyFilters applies a set of filters to a record list
func ApplyFilters(records []string, filters ...Filter) []string {
	if len(filters) == 0 {
		return records
	}

	filteredRecords := make([]string, 0, len(records))

	for _, record := range records {
		keep := true

		for _, filter := range filters {
			if !filter(record) {
				keep = false
				break
			}
		}

		if keep {
			filteredRecords = append(filteredRecords, record)
		}
	}

	return filteredRecords
}

// FilterForNiceString filters out records which not apply to rules for nice string
func FilterForNiceString(records []string) []string {
	return ApplyFilters(records,
		filterWords3Vowels,
		filterDoppleLetters,
		filterNotContainsString,
	)
}

// FilterForNiceStringAdvanced filters out records for second puzzle
func FilterForNiceStringAdvanced(records []string) []string {
	return ApplyFilters(records,
		filterPairWithBetweenLetter,
		filterTwoPairs,
	)
}

func filterWords3Vowels(record string) bool {
	vowels := []string{"a", "e", "i", "o", "u"}
	count := 0

	for _, vowel := range vowels {
		count = count + strings.Count(record, vowel)
	}

	return count >= 3
}

func filterDoppleLetters(record string) bool {
	for i := 0; i < len(record)-1; i++ {
		if record[i] == record[i+1] {
			return true
		}
	}
	return false
}

func filterNotContainsString(record string) bool {
	return !regexp.MustCompile("ab|cd|pq|xy").MatchString(record)
}

func filterPairWithBetweenLetter(record string) bool {
	for i := 0; i < len(record)-2; i++ {
		if record[i] == record[i+2] {
			return true
		}
	}
	return false
}

func filterTwoPairs(record string) bool {
	for i := 0; i < len(record)-2; i++ {
		if strings.Count(record, record[i:i+2]) >= 2 {
			return true
		}
	}
	return false
}

func main() {
	file, readErr := os.Open("input.txt")
	defer file.Close()

	if readErr != nil {
		panic(readErr)
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	filteredLines := FilterForNiceString(lines)
	fmt.Println("Puzzel 1 => Number of strings that are nice:", len(filteredLines))

	filteredLines = FilterForNiceStringAdvanced(lines)
	fmt.Println("Puzzel 2 => Number of strings that are nice:", len(filteredLines))
}
