package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	adapters = []int{0}
	device   int
)

func main() {

	file, err := os.Open("input.txt")
	defer file.Close()

	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		n, _ := strconv.Atoi(line)
		adapters = append(adapters, n)
	}
	sort.Ints(adapters)
	device = adapters[len(adapters)-1] + 3

	min, max := calculateTaskOne()
	result1 := min * max
	fmt.Println("There are", min, "diferences of 1 jolt")
	fmt.Println("There are", max, "diferences of 3 jolt")
	fmt.Println(" => Final Result 1:", result1)

	combinations := calculateCombinations()
	fmt.Println("Possible number of combinations: ", combinations)
}

func calculateCombinations() int {
	memory := make(map[int]int)
	memory[0] = 1

	for i := 1; i <= device; i++ {
		if contains(adapters, i) || i == device {
			j := 0

			if val, ok := memory[i-3]; ok {
				j = j + val
			}
			if val, ok := memory[i-2]; ok {
				j = j + val
			}
			if val, ok := memory[i-1]; ok {
				j = j + val
			}
			memory[i] = j
		}
	}

	return memory[device]
}

func contains(list []int, value int) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}

	return false
}

func calculateTaskOne() (int, int) {
	cache := [4]int{0, 0, 0, 0}

	if adapters[0] > 3 {
		return 0, 3
	}

	cache[adapters[0]] = cache[adapters[0]] + 1
	for i := 0; i+1 < len(adapters); i++ {
		diff := adapters[i+1] - adapters[i]

		if diff > 3 {
			break
		}
		cache[diff] = cache[diff] + 1
	}
	cache[3]++
	return cache[1], cache[3]
}
