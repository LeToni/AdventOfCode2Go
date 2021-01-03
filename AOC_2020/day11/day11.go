package main

import (
	"bufio"
	"fmt"
	"os"
)

type Position struct {
	x, y int
}

func (position *Position) withinSeatingArea() bool {
	return !(position.x < 0 || position.x >= len(seats) || position.y < 0 || position.y >= len(seats[position.x]))
}

func (position *Position) isOccupied() bool {
	if position.withinSeatingArea() && seats[position.x][position.y] == '#' {
		return true
	} else {
		return false
	}
}

func (position *Position) occupiedSeatNeighbors() int {
	count := 0
	for _, direction := range directions {
		pos := position.adjacentSeatAt(direction)
		if pos.isOccupied() {
			count++
		}
	}
	return count
}

func (position *Position) adjacentSeatAt(direction Position) *Position {
	return &Position{position.x + direction.x, position.y + direction.y}
}

func countOccupiedSeats() int {
	count := 0
	for _, row := range seats {
		for _, seat := range row {
			if seat == '#' {
				count++
			}
		}
	}
	return count
}

var (
	directions = []Position{
		{0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1},
	}
	seats [][]byte
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

	for steps := 0; steps < len(seats); steps++ {
		currentSeats := make([][]byte, len(seats))
		copy(currentSeats, seats)

		for i := range seats {
			for j := range seats[i] {
				pos := Position{i, j}
				occupiedSeats := pos.occupiedSeatNeighbors()
				if seats[pos.x][pos.y] == 'L' && occupiedSeats == 0 {
					currentSeats[pos.x][pos.y] = '#'
				} else if seats[pos.x][pos.y] == '#' && occupiedSeats >= 4 {
					currentSeats[pos.x][pos.y] = 'L'
				}
			}
		}
		seats = currentSeats
	}

	occupiedSeats := countOccupiedSeats()
	fmt.Println("At the end, number of occupied seats:", occupiedSeats)
}
