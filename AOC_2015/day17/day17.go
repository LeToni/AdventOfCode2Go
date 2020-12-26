package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const (
	capacity = 150
)

var (
	containers = []int{}
)

func main() {
	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		container, err := strconv.Atoi(input)
		if err != nil {
			panic(err)
		}
		containers = append(containers, container)
	}
	sort.Ints(containers)
	total := amountPossibleCombinations(containers, capacity)
	fmt.Println("Amount of possible combos:", total)

	total = amountUniqueCombinations()
	fmt.Println("Amount of unique combos:", total)
}

func amountPossibleCombinations(containers []int, target int) int {
	total := 0

	if len(containers) == 2 {
		if containers[0]+containers[1] == target {

			return 1
		} else {
			return 0
		}
	}

	for i := 1; i < len(containers); i++ {
		if containers[0]+containers[i] == target {
			total = total + 1
		}
	}

	return total + amountPossibleCombinations(containers[1:], target) + amountPossibleCombinations(containers[1:], target-containers[0])
}

func amountUniqueCombinations() int {
	combinations := map[int]int{}
	checkCombination := func(combo []int) {
		sum := 0
		for _, container := range combo {
			sum += container
		}
		if sum == capacity {
			combinations[len(combo)]++
		}
	}

	for num := 1; num <= len(containers); num++ {
		combo := make([]int, num)

		var nextCombination func(int, int)
		last := len(combo) - 1
		nextCombination = func(i, from int) {
			for j := from; j < len(containers); j++ {
				combo[i] = containers[j]
				if i == last {
					checkCombination(combo)
				} else {
					nextCombination(i+1, j+1)
				}
			}
		}

		nextCombination(0, 0)

		if len(combinations) > 0 {
			return combinations[num]
		}
	}
	return 0
}
