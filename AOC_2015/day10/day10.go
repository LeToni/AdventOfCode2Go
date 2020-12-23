package main

import (
	"fmt"
	"strconv"
)

var (
	numbers []byte
)

func main() {
	numbers = []byte("3113322113")

	maxProcess := 50

	for process := 0; process < maxProcess; process++ {
		var nextNumber []byte
		for i := 0; i < len(numbers); i++ {
			number := numbers[i]
			count := 1
			for i < len(numbers)-1 && numbers[i+1] == number {
				i++
				count++
			}

			nextNumber = append(nextNumber, strconv.Itoa(count)[0])
			nextNumber = append(nextNumber, number)
		}

		numbers = nextNumber

		if process+1 == 40 {
			fmt.Println("Results after 40 iteration: ", len(numbers))
		}
	}

	fmt.Println("Endresult:", len(numbers))
}
