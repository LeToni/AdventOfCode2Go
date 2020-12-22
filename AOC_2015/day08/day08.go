package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	stringCharacters  int
	memoryCharacters  int
	encodedCharacters int
)

func main() {
	file, err := os.Open("input.txt")
	defer file.Close()

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		stringInput := scanner.Text()

		stringCharacters = stringCharacters + len(stringInput)
		unquotedString, err1 := strconv.Unquote(stringInput)
		if err1 != nil {
			panic(err)
		}
		quotedString := strconv.Quote(stringInput)

		memoryCharacters = memoryCharacters + len(unquotedString)
		encodedCharacters = encodedCharacters + len(quotedString)
	}

	fmt.Println("Total difference 1: ", stringCharacters-memoryCharacters)
	fmt.Println("Total difference 2: ", encodedCharacters-stringCharacters)
}
