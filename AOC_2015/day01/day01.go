package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, readErr := os.Open("input.txt")
	defer file.Close()

	if readErr != nil {
		panic(readErr)
	}

	floor := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		floor += takeToFloor(text)
		whichInstruction(text)
	}

	fmt.Println("Instructions take Santa to floor ", floor)

}

func takeToFloor(text string) int {
	up := strings.Count(text, "(")
	down := strings.Count(text, ")")
	floor := up - down

	return floor
}

func whichInstruction(text string) {
	floor := 0
	for number, instruction := range text {
		if instruction == '(' {
			floor++
		} else {
			floor--
		}
		if floor < 0 {
			fmt.Println("Instruction ", number+1, "takes Santa to floor -1")
			return
		}
	}
}
