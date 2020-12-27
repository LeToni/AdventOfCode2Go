package main

import (
	"bufio"
	"os"
)

var (
	grid = make([][]byte, 100)
)

func main() {

	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		row := scanner.Bytes()
		grid[i] = row
		i++
	}
}
