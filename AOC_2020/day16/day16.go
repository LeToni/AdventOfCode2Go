package main

import (
	"io/ioutil"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	content := strings.Split(string(file), "\n\n")

}
