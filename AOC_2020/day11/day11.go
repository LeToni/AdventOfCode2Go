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
	file, err := os.Open("input_test.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Bytes()
		seats = append(seats, row)
	}

	seatsChanged := true

	for seatsChanged {
		seatsChanged = false
		currentSeats := [][]byte{}
		for i := range seats {
			row := []byte{}

			for j := range seats[i] {
				pos := Position{i, j}
				occupiedSeats := pos.occupiedSeatNeighbors()
				if seats[pos.x][pos.y] == '.' {
					row = append(row, '.')
				} else if seats[pos.x][pos.y] == 'L' {
					if occupiedSeats == 0 {
						row = append(row, '#')
						seatsChanged = true
					} else {
						row = append(row, 'L')
					}
				} else if seats[pos.x][pos.y] == '#' {
					if occupiedSeats >= 4 {
						row = append(row, 'L')
						seatsChanged = true
					} else {
						row = append(row, '#')
					}
				}

			}
			currentSeats = append(currentSeats, row)
		}

		seats = currentSeats
	}

	occupiedSeats := countOccupiedSeats()
	fmt.Println("At the end, number of occupied seats:", occupiedSeats)
}
