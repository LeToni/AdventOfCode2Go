package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	preamble int = 25
)

var (
	numbers = []int{}
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
			fmt.Println("First invalid number found:", numbers[i+preamble])
			break
		}
	}
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
