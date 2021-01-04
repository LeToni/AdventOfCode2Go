package main

import (
	"bufio"
	"fmt"
	"os"
)

type Position struct {
	row, seat int
}

func (pos *Position) withinArea() bool {
	return !(pos.row < 0 || pos.row >= len(seatArea) || pos.seat < 0 || pos.seat >= len(seatArea[pos.row]))
}

func (pos *Position) isOccupied() bool {
	if seatArea[pos.row][pos.seat] == '#' {
		return true
	} else {
		return false
	}
}

func (pos *Position) adjacentNeighborsOccupied() int {
	count := 0

	for _, direction := range directions {
		neighbor := pos.neighborAt(direction)
		if neighbor.withinArea() && neighbor.isOccupied() {
			count = count + 1
		}
	}
	return count
}

func (pos *Position) seatWithinSightOccupied(direction Position) bool {
	for i := 1; ; i++ {
		posSight := Position{pos.row + direction.row*i, pos.seat + direction.seat*i}
		if !posSight.withinArea() {
			return false
		} else if posSight.withinArea() && seatArea[posSight.row][posSight.seat] == '#' {
			return true
		} else if posSight.withinArea() && seatArea[posSight.row][posSight.seat] == 'L' {
			return false
		}
	}
}

func (pos *Position) neighborsWithinSightOccupied() int {
	count := 0

	for _, direction := range directions {
		if pos.seatWithinSightOccupied(direction) {
			count = count + 1
		}
	}
	return count
}

func (pos *Position) neighborAt(direction Position) *Position {
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

func TaskOne(area [][]rune) {
	seatArea = area
	seatsChanged := true
	for seatsChanged {
		seatsChanged = false
		newSeatArea := [][]rune{}
		for row := range seatArea {
			newRow := []rune{}
			for seat, seatStatus := range seatArea[row] {
				newSeatstatus := seatStatus
				pos := Position{row, seat}
				if occupiedSeats := pos.adjacentNeighborsOccupied(); seatStatus == '#' && occupiedSeats >= 4 {
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

func TaskTwo(area [][]rune) {
	seatArea = area
	seatsChanged := true
	for seatsChanged {
		seatsChanged = false
		newSeatArea := [][]rune{}
		for row := range seatArea {
			newRow := []rune{}
			for seat, seatStatus := range seatArea[row] {
				newSeatstatus := seatStatus
				pos := Position{row, seat}
				if occupiedSeats := pos.neighborsWithinSightOccupied(); seatStatus == '#' && occupiedSeats >= 5 {
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
	fmt.Println("Task 2 -> Total amount of seats:", countTotalOccupiedSeats())
}

func main() {
	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	scannedArea := [][]rune{}
	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		row := []rune{}
		for _, r := range scanner.Text() {
			row = append(row, r)
		}
		scannedArea = append(scannedArea, row)
	}
	TaskOne(scannedArea)
	TaskTwo(scannedArea)
}
