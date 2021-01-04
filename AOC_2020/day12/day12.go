package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

const (
	N int = iota
	E
	S
	W
)

type Instruction struct {
	action rune
	value  int
}

func (instruct *Instruction) executeOn(ship *Ship) {
	if instruct.action == 'R' || instruct.action == 'L' {
		ship.turn(instruct.action, instruct.value)
	} else if instruct.action == 'F' {
		ship.move(instruct.value)
	} else {
		ship.navigate(instruct.action, instruct.value)
	}
}

type Ship struct {
	facing int
	x, y   int
}

func (ship *Ship) navigate(cardinalDirection rune, units int) {
	switch cardinalDirection {
	case 'N':
		ship.y = ship.y + units
	case 'E':
		ship.x = ship.x + units
	case 'S':
		ship.y = ship.y - units
	case 'W':
		ship.x = ship.x - units
	default:
		err := fmt.Errorf("Not a valid cardinal direction: %d", cardinalDirection)
		panic(err)
	}
}

func (ship *Ship) move(units int) {
	switch ship.facing {
	case E:
		ship.x = ship.x + units
	case S:
		ship.y = ship.y - units
	case W:
		ship.x = ship.x - units
	default:
		ship.y = ship.y + units
	}
}

func (ship *Ship) turn(direction rune, turnRadius int) {
	if direction == 'R' {
		ship.facing = (ship.facing + (turnRadius / 90)) % 4
	} else {
		ship.facing = ((ship.facing - (turnRadius / 90)) + 4) % 4
	}
}

func main() {
	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	regInstructionFilter := regexp.MustCompile(`(\w)(\d+)`)
	instructions := []Instruction{}
	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		input := scanner.Text()
		filteredInput := regInstructionFilter.FindStringSubmatch(input)
		action := []rune(filteredInput[1])[0]
		value, _ := strconv.Atoi(filteredInput[2])
		instruction := Instruction{action, value}
		instructions = append(instructions, instruction)
	}

	ship := &Ship{x: 0, y: 0, facing: E}

	for _, instruction := range instructions {
		instruction.executeOn(ship)
	}

	result := math.Abs(float64(ship.x)) + math.Abs(float64(ship.y))
	fmt.Println("Task1 -> Distance between starting point and ship:", result)
}
