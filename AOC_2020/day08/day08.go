package main

import (
	"bufio"
	"fmt"
	"os"
)

type Instruction struct {
	operation string
	memory    int
	executed  bool
}

func runInstructionsOneTime() {
	i := 0
	for i >= 0 && i < len(instructions) && instructions[i].executed != true {
		instruction := instructions[i]
		instruction.executed = true
		switch instruction.operation {
		case "nop":
			i++
		case "acc":
			acc = acc + instruction.memory
			i++
		case "jmp":
			i = i + instruction.memory
		}
	}
}

func resetInstructions() {
	for _, instruction := range instructions {
		instruction.executed = false
	}
}

var (
	instructions = []*Instruction{}
	acc          int
)

func main() {

	file, err := os.Open("input.txt")
	defer file.Close()

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()

		var (
			op  string
			val int
		)
		if n, _ := fmt.Sscanf(input, "%s %d", &op, &val); n == 2 {
			instructions = append(instructions, &Instruction{op, val, false})
		} else {
			panic(input)
		}
	}
	runInstructionsOneTime()
	fmt.Println("Value currently in accumulator:", acc)

	resetInstructions()
}
