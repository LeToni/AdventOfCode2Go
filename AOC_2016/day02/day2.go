package main

import (
	"bufio"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	defer file.Close()

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		directions := scanner.Text()

		for _, direction := range directions {

		}
	}
}
