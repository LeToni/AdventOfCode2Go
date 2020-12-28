package main

import "fmt"

const (
	target int = 33100000
)

var (
	houseVisitOne = make([]int, target)
	houseVisitTwo = make([]int, target)
)

func main() {
	houseNumber := calculateTaskOne()
	fmt.Println("Part 1: Lowest house number: ", houseNumber)
	houseNumber = calculateTaskTwo()
	fmt.Println("Part 2: Lowest house number: ", houseNumber)
}

// reddit post
func calculateTaskOne() int {
	for elf := 1; elf <= target/10; elf++ {
		upto := target / 10 / elf
		for i := 1; i <= upto; i++ {
			houseVisitOne[i*elf] += elf * 10
		}
	}

	for houseNumber := 1; houseNumber < len(houseVisitOne); houseNumber++ {
		presents := houseVisitOne[houseNumber]
		if presents >= target {
			return houseNumber
		}
	}
	return 0
}

func calculateTaskTwo() int {
	for elf := 1; elf <= target/10; elf++ {
		upto := target / 10 / elf
		if upto > 50 {
			upto = 50
		}

		for i := 1; i <= upto; i++ {
			houseVisitOne[i*elf] += elf * 11
		}
	}

	for houseNumber := 1; houseNumber < len(houseVisitOne); houseNumber++ {
		presents := houseVisitOne[houseNumber]
		if presents >= target {
			return houseNumber
		}
	}
	return 0
}
