package main

import (
	"io/ioutil"
	"strings"
)

type Intervall struct {
	lowerBound, upperBound int
}

func (inter *Intervall) withinBoundary(number int) bool {
	if inter.lowerBound <= number && number <= inter.upperBound {
		return true
	} else {
		return false
	}
}

type TicketField struct {
	rule1 Intervall
	rule2 Intervall
}

func (tf *TicketField) isValid(number int) bool {

	if tf.rule1.withinBoundary(number) || tf.rule2.withinBoundary(number) {
		return true
	} else {
		return false
	}
}

func main() {
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	content := strings.Split(string(file), "\n\n")

}
