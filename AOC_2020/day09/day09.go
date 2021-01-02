package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const (
	preamble int = 25
)

var (
	numbers       = []int{}
	invalidNumber int
)

func main() {
	file, err := os.Open("input.txt")
	defer file.Close()

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		numbers = append(numbers, number)
	}

	for i := 0; i < len(numbers)-preamble; i++ {
		if !isValid(numbers[i+preamble], i, i+preamble) {
			invalidNumber = numbers[i+preamble]
			fmt.Println("First invalid number found:", invalidNumber)
			break
		}
	}

	setStart, setEnd, setFound := findContiguousSet(invalidNumber)

	if !setFound {
		fmt.Println("No continuous set found")
	} else {
		setNumbers := numbers[setStart:setEnd]
		sort.Ints(setNumbers)
		weak := setNumbers[0] + setNumbers[len(setNumbers)-1]
		fmt.Println("Encryption weakness number:", weak)
	}
}

func findContiguousSet(number int) (int, int, bool) {
	sum := 0

	for i := 0; i < len(numbers); i++ {
		sum = numbers[i]
		for j := i + 1; j < len(numbers); j++ {
			sum = sum + numbers[j]

			if sum == invalidNumber {
				return i, j + 1, true
			}
			if sum > invalidNumber {
				break
			}
		}
	}

	return 0, 0, false
}
func isValid(number int, start, end int) bool {
	eval := false

	for i := start; i < end; i++ {
		for j := i + 1; j < end; j++ {
			if number == numbers[i]+numbers[j] {
				return true
			}
		}
	}

	return eval
}
