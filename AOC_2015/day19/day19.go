package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type rule struct {
	leftSide, rightSide string
}

var (
	rules     []*rule
	molecules = make(map[string]bool)
	molecule  string
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
		if input == "" {
			continue
		}

		var inputLeft, inputRight string
		if n, _ := fmt.Sscanf(input, "%s => %s", &inputLeft, &inputRight); n == 2 {
			inputRule := rule{leftSide: inputLeft, rightSide: inputRight}
			rules = append(rules, &inputRule)
		} else {
			molecule = input
		}
	}

	for _, r := range rules {
		regex := regexp.MustCompile(r.leftSide)

		if !regex.MatchString(molecule) {
			continue
		}
		particles := regex.FindAllStringIndex(molecule, -1)

		for _, particle := range particles {
			newMolecule := make([]byte, len(molecule))
			copy(newMolecule, molecule)
			newMolecule = append(newMolecule[:particle[0]], append([]byte(r.rightSide), newMolecule[particle[1]:]...)...)
			// test := append([]byte(r.rightSide), newMolecule[particle[0]:]...)
			// // newMolecule = append(newMolecule[:particle[0]])
			// print(string(test))
			molecules[string(newMolecule)] = true
		}
	}

	fmt.Println(len(molecules))
}
