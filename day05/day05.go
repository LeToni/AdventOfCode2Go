package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Day4 Advent of Code Challenge
func main() {
	// Open the file
	file, err := os.Open("input.txt")
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	max := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		code := scanner.Text()

		row, col := encode(code)
		seatID := calculateSeatID(row, col)

		if max < seatID {
			max = seatID
		}

	}
	fmt.Println("Highest SeatID on the list: ", max)
}

func calculateSeatID(row int, col int) int {
	return row*8 + col
}

func encode(code string) (int, int) {
	codeRow := code[:len(code)-3]
	codeCol := code[len(code)-3:]

	codeRow = strings.ReplaceAll(codeRow, "F", "0")
	codeRow = strings.ReplaceAll(codeRow, "B", "1")

	codeCol = strings.ReplaceAll(codeCol, "L", "0")
	codeCol = strings.ReplaceAll(codeCol, "R", "1")

	binRow, errR := strconv.ParseInt(codeRow, 2, 64)
	binCol, errC := strconv.ParseInt(codeCol, 2, 64)

	if errR != nil || errC != nil {
		fmt.Println("Not able to convert code")
		if errR == nil {
			panic(errR)
		} else {
			panic(errC)
		}
	}

	return int(binRow), int(binCol)
}
