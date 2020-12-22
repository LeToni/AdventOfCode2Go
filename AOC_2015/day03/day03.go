package main

import (
	"bufio"
	"fmt"
	"os"
)

type coordinates struct {
	x, y int
}

// Santa representation
type Santa struct {
	coord coordinates
}

func main() {
	file, readErr := os.Open("input.txt")
	defer file.Close()

	if readErr != nil {
		panic(readErr)
	}
	houseVisited := 0
	houseVisitWithRobot := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		route := scanner.Text()
		houseVisited += santaOnly(route)
		houseVisitWithRobot += santaWithRobot(route)
	}
	fmt.Println("Santa visited", houseVisited, " houses")
	fmt.Println("Santa and robote: ", houseVisitWithRobot)
}

func santaOnly(route string) (homesDelivered int) {
	santa := Santa{}
	santas := []*Santa{&santa}

	homesDelivered = deliverToHomes(santas, route)
	return homesDelivered
}

func santaWithRobot(route string) (homesDelivered int) {
	santa := Santa{}
	robot := Santa{}
	santas := []*Santa{&santa, &robot}

	homesDelivered = deliverToHomes(santas, route)
	return homesDelivered
}

func deliverToHomes(santas []*Santa, route string) int {
	homes := make(map[coordinates]bool)
	coord := coordinates{0, 0}
	homes[coord] = true

	for index, direction := range route {

		santa := santas[index%len(santas)]
		if direction == '<' {
			santa.moveToWest()
		}
		if direction == '>' {
			santa.moveToEast()
		}
		if direction == '^' {
			santa.moveToNorth()
		}
		if direction == 'v' {
			santa.moveToSouth()
		}
		santa.delivers(homes)
	}
	return len(homes)
}

func (santa *Santa) moveToWest() {
	santa.coord.x = santa.coord.x - 1
}

func (santa *Santa) moveToNorth() {
	santa.coord.y = santa.coord.y + 1
}

func (santa *Santa) moveToEast() {
	santa.coord.x = santa.coord.x + 1
}

func (santa *Santa) moveToSouth() {
	santa.coord.y = santa.coord.y - 1
}

func (santa *Santa) delivers(homes map[coordinates]bool) {
	homes[santa.coord] = true
}
