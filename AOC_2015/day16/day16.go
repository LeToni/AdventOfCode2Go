package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var (
	sueDnaInfo = map[string]string{
		"children":    "3",
		"cats":        "7",
		"samoyeds":    "2",
		"pomeranians": "3",
		"akitas":      "0",
		"vizslas":     "0",
		"goldfish":    "5",
		"trees":       "3",
		"cars":        "2",
		"perfumes":    "1",
	}

	regexFilter = map[string]*regexp.Regexp{
		"number":      regexp.MustCompile(`Sue (\d+)`),
		"children":    regexp.MustCompile(`children: (\d+)`),
		"cats":        regexp.MustCompile(`cats: (\d+)`),
		"samoyeds":    regexp.MustCompile(`samoyeds: (\d+)`),
		"pomeranians": regexp.MustCompile(`pomeranians: (\d+)`),
		"akitas":      regexp.MustCompile(`akitas: (\d+)`),
		"vizslas":     regexp.MustCompile(`vizslas: (\d+)`),
		"goldfish":    regexp.MustCompile(`goldfish: (\d+)`),
		"trees":       regexp.MustCompile(`trees: (\d+)`),
		"cars":        regexp.MustCompile(`cars: (\d+)`),
		"perfumes":    regexp.MustCompile(`perfumes: (\d+)?`),
	}

	auntSueNr string
)

func main() {

	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		info := scanner.Text()
		sue := true
		for dnaType, amount := range sueDnaInfo {
			match := regexFilter[dnaType].FindStringSubmatch(info)

			if len(match) > 0 && amount != match[1] {
				sue = false
			}
		}

		if sue == true {
			auntSueNr = regexFilter["number"].FindStringSubmatch(info)[1]
			fmt.Println("It should be Aunt Sue ", auntSueNr)
			break
		}
	}

}
