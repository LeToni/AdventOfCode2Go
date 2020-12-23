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
	paths         = make(map[Path]int)
	cities        []string
	routes        [][]string
	shortestRoute []string
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
			path = Path{fromCity: toCity, toCity: fromCity}
			paths[path] = distance
		}
	}

	permutations(len(cities))

	shortestRoute := math.MaxInt32
	longestRoute := math.MinInt32
	for _, route := range routes {
		var total int = 0
		for i := 0; i < len(route)-1; i++ {
			path := Path{fromCity: route[i], toCity: route[i+1]}
			total = total + paths[path]
		}

		if shortestRoute > total {
			shortestRoute = total
		}
		if longestRoute < total {
			longestRoute = total
		}
	}

	fmt.Println("Shortest distance found: ", shortestRoute)
	fmt.Println("Longest distance found: ", longestRoute)
}

func permutations(n int) {
	if n == 1 {
		citiesCopy := make([]string, len(cities))
		copy(citiesCopy, cities)

		routes = append(routes, citiesCopy)
	} else {
		for i := 0; i < n-1; i++ {
			permutations(n - 1)
			if n%2 == 0 {
				swap(cities, i, n-1)
			} else {
				swap(cities, 0, n-1)
			}
		}
		permutations(n - 1)
	}
}

func swap(cities []string, i, j int) {
	cities[i], cities[j] = cities[j], cities[i]
}

func addCityToList(city string) {

	for _, c := range cities {
		if c == city {
			return
		}
	}

	cities = append(cities, city)
}
