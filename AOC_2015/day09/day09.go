package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Path struct {
	fromCity string
	toCity   string
}

var (
	infinity = math.Inf(1)
	paths    = make(map[Path]int)
	cities   = []string{}
)

func main() {
	file, err := os.Open("input.txt")
	defer file.Close()

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		pathInput := scanner.Text()

		var (
			fromCity, toCity string
			distance         int
		)
		if n, _ := fmt.Sscanf(pathInput, "%s to %s = %d", &fromCity, &toCity, &distance); n == 3 {
			path := Path{fromCity: fromCity, toCity: toCity}
			addCityToList(fromCity)
			addCityToList(toCity)
			paths[path] = distance
		}
	}

	graph := make([][]Path, len(cities), len(cities))

	for i, fromCity := range cities {
		for _, toCity := range cities {
			path := Path{fromCity: fromCity, toCity: toCity}
			if _, ok := paths[path]; !ok {
				if fromCity == toCity {
					paths[path] = 0
				} else {
					paths[path] = int(math.Inf((1)))
				}

			}
			graph[i] = append(graph[i], path)

		}
	}

	fmt.Print(len(graph))
}

func floydWarshall(graph [][]Path) {

}
func addCityToList(city string) {

	for _, c := range cities {
		if c == city {
			return
		}
	}

	cities = append(cities, city)
}
