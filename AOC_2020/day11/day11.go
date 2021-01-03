package main

type Position struct {
	x, y int
}

func (position *Position) withinSeatingArea() bool {
	return !(position.x < 0 || position.x >= len(seats) || position.y < len(seats[position.x]))
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

var (
	directions = []Position{
		{0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1},
	}
	seats [][]byte
)
