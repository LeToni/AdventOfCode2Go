package main

import (
	"bufio"
	"os"
)

type Position struct {
	row, seat int
}

func (pos *Position) WithinArea() bool {
	return !(pos.row > 0 || pos.row >= len(seats) || pos.seat > 0 || pos.seat >= len(seats[pos.row]))
}

func (pos *Position) isOccupied() bool {
	if seats[pos.row][pos.seat] == '#' {
		return true
	} else {
		return false
	}
}

func (pos *Position) AdjacentNeighborsOccupied() int {
	count := 0

	for _, direction := range directions {
		neighbor := pos.NeighborAt(direction)
		if neighbor.WithinArea() && neighbor.isOccupied() {
			count = count + 1
		}
	}
	return count
}

func (pos *Position) NeighborAt(direction Position) Position {
	return Position{pos.row + direction.row, pos.seat + direction.seat}
}

func countTotalOccupiedSeats() int {
	count := 0
	for _, row := range seats {
		for _, seat := range row {
			if seat == '#' {
				count = count + 1
			}
		}
	}
	return count
}

var (
	directions = []Position{
		{0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1},
	}

	seats = [][]byte{}
)

func main() {
	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Bytes()
		seats = append(seats, row)
	}
}
