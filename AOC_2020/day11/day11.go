package main

import (
	"bufio"
	"fmt"
	"os"
)

type Position struct {
	row, seat int
}

func (pos *Position) WithinArea() bool {
	return !(pos.row < 0 || pos.row >= len(seatArea) || pos.seat < 0 || pos.seat >= len(seatArea[pos.row]))
}

func (pos *Position) isOccupied() bool {
	if seatArea[pos.row][pos.seat] == '#' {
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

func (pos *Position) SeatWithinSightOccupied(direction Position) bool {
	for i := 1; ; i++ {
		posSight := Position{pos.row + direction.row*i, pos.seat + direction.seat*i}
		if !posSight.WithinArea() {
			return false
		} else if posSight.WithinArea() && seatArea[posSight.row][posSight.seat] == '#' {
			return true
		} else if posSight.WithinArea() && seatArea[posSight.row][posSight.seat] == 'L' {
			return false
		}
	}
}

func (pos *Position) NeighborsWithinSightOccupied() int {
	count := 0

	for _, direction := range directions {
		if pos.SeatWithinSightOccupied(direction) {
			count = count + 1
		}
	}
	return count
}

func (pos *Position) NeighborAt(direction Position) *Position {
	return &Position{pos.row + direction.row, pos.seat + direction.seat}
}

func countTotalOccupiedSeats() int {
	count := 0
	for _, row := range seatArea {
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
		{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1},
	}

	seatArea = [][]rune{}
)

func main() {
	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := []rune{}
		for _, r := range scanner.Text() {
			row = append(row, r)
		}
		seatArea = append(seatArea, row)
	}

	seatsChanged := true
	for seatsChanged {
		seatsChanged = false
		newSeatArea := [][]rune{}
		for row := range seatArea {
			newRow := []rune{}
			for seat, seatStatus := range seatArea[row] {
				newSeatstatus := seatStatus
				pos := Position{row, seat}
				if occupiedSeats := pos.NeighborsWithinSightOccupied(); seatStatus == '#' && occupiedSeats >= 5 {
					newSeatstatus = 'L'
					seatsChanged = true
				} else if seatStatus == 'L' && occupiedSeats == 0 {
					newSeatstatus = '#'
					seatsChanged = true
				}
				newRow = append(newRow, newSeatstatus)
			}
			newSeatArea = append(newSeatArea, newRow)
		}
		seatArea = newSeatArea
	}
	fmt.Println("Task 1 -> Total amount of seats:", countTotalOccupiedSeats())
}
