package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Instruction struct {
	input1, input2 string
	gate           Gate
}

type Gate uint8

const (
	NONE Gate = iota
	AND
	OR
	RSHIFT
	LSHIFT
	NOT
)

var (
	instructionBook = make(map[string]Instruction)
	wires           = make(map[string]uint16)
)

func loadInstructions() {
	file, err := os.Open("input.txt")
	defer file.Close()

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instructionInput := scanner.Text()

		var (
			gate                       Gate
			input1, input2, outputWire string
		)

		if n, _ := fmt.Sscanf(instructionInput, "%s -> %s", &input1, &outputWire); n == 2 {
			gate = NONE
		} else if n, _ := fmt.Sscanf(instructionInput, "%s AND %s -> %s", &input1, &input2, &outputWire); n == 3 {
			gate = AND
		} else if n, _ := fmt.Sscanf(instructionInput, "%s OR %s -> %s", &input1, &input2, &outputWire); n == 3 {
			gate = OR
		} else if n, _ := fmt.Sscanf(instructionInput, "%s RSHIFT %s -> %s", &input1, &input2, &outputWire); n == 3 {
			gate = RSHIFT
		} else if n, _ := fmt.Sscanf(instructionInput, "%s LSHIFT %s -> %s", &input1, &input2, &outputWire); n == 3 {
			gate = LSHIFT
		} else if n, _ := fmt.Sscanf(instructionInput, "NOT %s -> %s", &input1, &outputWire); n == 2 {
			gate = NOT
		} else {
			panic(instructionInput)
		}

		instructionBook[outputWire] = Instruction{input1: input1, input2: input2, gate: gate}
	}
}

func getSignal(wire string) uint16 {

	if valueSignal, err := strconv.ParseUint(wire, 10, 16); err == nil {
		return uint16(valueSignal)
	}
	if signalValue, ok := wires[wire]; ok {
		return signalValue
	}

	instruction := instructionBook[wire]
	var signalValue uint16

	switch instruction.gate {
	case NONE:
		signalValue = getSignal(instruction.input1)
	case AND:
		signalValue = getSignal(instruction.input1) & getSignal(instruction.input2)
	case OR:
		signalValue = getSignal(instruction.input1) | getSignal(instruction.input2)
	case RSHIFT:
		signalValue = getSignal(instruction.input1) >> getSignal(instruction.input2)
	case LSHIFT:
		signalValue = getSignal(instruction.input1) << getSignal(instruction.input2)
	case NOT:
		signalValue = ^getSignal(instruction.input1)
	}

	wires[wire] = signalValue

	return signalValue
}

func main() {
	loadInstructions()
	var signal = getSignal("a")
	fmt.Println("Signal value of wire 'a': ", signal)

	// Reset wires
	for k := range wires {
		delete(wires, k)
	}

	wires["b"] = signal
	signal = getSignal("a")
	fmt.Println("Signal value of wire 'a': ", signal)
}
