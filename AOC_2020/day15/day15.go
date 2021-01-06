package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	firstStint  int = 2020
	secondStint int = 30000000
)

var (
	turn         int = 1
	lastSpoken   int = 0
	spokeOnTurns     = map[int][]int{}
)

func Playgame(maxTurns int) {
	for turn <= maxTurns {
		if turns := spokeOnTurns[lastSpoken]; len(turns) == 1 {
			lastSpoken = 0
		} else {
			lastSpoken = turns[len(turns)-1] - turns[len(turns)-2]
		}
		spokeOnTurns[lastSpoken] = append(spokeOnTurns[lastSpoken], turn)
		turn++
	}
}

func main() {
	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		panic(file)
	}

	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		startingNumbers := strings.Split(scanner.Text(), ",")
		for _, number := range startingNumbers {
			lastSpoken, _ = strconv.Atoi(number)
			spokeOnTurns[lastSpoken] = []int{turn}
			turn++
		}

	}

	Playgame(firstStint)
	fmt.Println("On turn", firstStint, "the last word spoken is", lastSpoken)
	Playgame(secondStint)
	fmt.Println("On turn", secondStint, "the last word spoken is", lastSpoken)
}
