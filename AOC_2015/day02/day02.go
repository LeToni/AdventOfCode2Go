package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func main() {
	file, readErr := os.Open("input.txt")
	defer file.Close()

	if readErr != nil {
		panic(readErr)
	}

	dimensionRegex := regexp.MustCompile(`^(\d+)x(\d+)x(\d+)$`)
	totalWrapper := 0
	totalRibbon := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		length, _ := strconv.Atoi(dimensionRegex.ReplaceAllString(text, "$1"))
		width, _ := strconv.Atoi(dimensionRegex.ReplaceAllString(text, "$2"))
		height, _ := strconv.Atoi(dimensionRegex.ReplaceAllString(text, "$3"))
		totalWrapper = totalWrapper + (2 * length * width) + (2 * width * height) + (2 * height * length)

		sidesDimension := []int{}
		sidesDimension = append(sidesDimension, length*width)
		sidesDimension = append(sidesDimension, width*height)
		sidesDimension = append(sidesDimension, height*length)
		sort.Ints(sidesDimension)
		totalWrapper = totalWrapper + sidesDimension[0]

		sides := []int{}
		sides = append(sides, width)
		sides = append(sides, length)
		sides = append(sides, height)
		sort.Ints(sides)

		totalRibbon += 2*sides[0] + 2*sides[1]
		totalRibbon += sides[0] * sides[1] * sides[2]

	}

	fmt.Println("total square feet of wrapping paper required: ", totalWrapper)
	fmt.Println("total feet of ribbon paper required: ", totalRibbon)
}
