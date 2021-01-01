package main

import (
	"bufio"
	"fmt"
	"os"
)

type Triangle struct {
	sideA, sideB, sideC int
}

func (triangle *Triangle) IsTriangle() bool {
	if triangle.sideA+triangle.sideB > triangle.sideC &&
		triangle.sideA+triangle.sideC > triangle.sideB &&
		triangle.sideB+triangle.sideC > triangle.sideA {
		return true
	} else {
		return false
	}
}

func TaskOne() {
	input_1, err := os.Open("input.txt")
	defer input_1.Close()
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(input_1)
	count := 0
	for scanner.Scan() {
		input := scanner.Text()

		var (
			a, b, c int
		)
		if n, _ := fmt.Sscanf(input, "  %d  %d  %d", &a, &b, &c); n == 3 {
			triangle := Triangle{a, b, c}
			if triangle.IsTriangle() {
				count++
			}
		} else {
			panic(n)
		}
	}
	fmt.Println("Task 1: Number of triangles found: ", count)
}

func TaskTwo() {
	input_2, err := os.Open("input.txt")
	defer input_2.Close()
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(input_2)

	var (
		triangle1, triangle2, triangle3 Triangle
		i                               = 0
		count                           = 0
	)

	for scanner.Scan() {
		input := scanner.Text()
		var (
			side1, side2, side3 int
		)
		if n, _ := fmt.Sscanf(input, "  %d  %d  %d", &side1, &side2, &side3); n == 3 {
			switch i {
			case 0:
				triangle1.sideA = side1
				triangle2.sideA = side2
				triangle3.sideA = side3
				i++
			case 1:
				triangle1.sideB = side1
				triangle2.sideB = side2
				triangle3.sideB = side3
				i++
			case 2:
				triangle1.sideC = side1
				triangle2.sideC = side2
				triangle3.sideC = side3
				i = 0
				if triangle1.IsTriangle() {
					count++
				}
				if triangle2.IsTriangle() {
					count++
				}
				if triangle3.IsTriangle() {
					count++
				}
			}
		} else {
			panic(n)
		}

	}
	fmt.Println("Task 2: Number of triangles found: ", count)
}

func main() {
	TaskOne()
	TaskTwo()
}
