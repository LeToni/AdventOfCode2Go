package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	off int = iota
	on
)

var (
	lights     [1000][1000]int
	brightness [1000][1000]int
)

func main() {
	file, readErr := os.Open("input.txt")
	defer file.Close()

	if readErr != nil {
		panic(readErr)
	}

	lightsReg := regexp.MustCompile(`\d+`)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines := scanner.Text()
		lightsInfo := lightsReg.FindAllString(lines, -1)

		fromX, _ := strconv.Atoi(lightsInfo[0])
		fromY, _ := strconv.Atoi(lightsInfo[1])
		toX, _ := strconv.Atoi(lightsInfo[2])
		toY, _ := strconv.Atoi(lightsInfo[3])

		if strings.Contains(lines, "on") {
			turnLightsOn(fromX, fromY, toX, toY)
			turnBrightnessOn(fromX, fromY, toX, toY)
		} else if strings.Contains(lines, "off") {
			turnLightsOff(fromX, fromY, toX, toY)
			turnBrightnessOff(fromX, fromY, toX, toY)
		} else if strings.Contains(lines, "toggle") {
			toggleLights(fromX, fromY, toX, toY)
			toggleBrightness(fromX, fromY, toX, toY)

		}
	}

	totalLightsOn := 0
	for _, row := range lights {
		for _, col := range row {
			if col == on {
				totalLightsOn++
			}
		}
	}
	fmt.Println("Total lights on:", totalLightsOn)

	totalBrightness := 0
	for _, row := range brightness {
		for _, col := range row {
			totalBrightness = totalBrightness + col
		}
	}
	fmt.Println("Total brightness:", totalBrightness)
}

func turnLightsOn(fromX, fromY, toX, toY int) {

	for i := fromX; i <= toX; i++ {
		for j := fromY; j <= toY; j++ {
			lights[i][j] = on
		}
	}
}

func turnBrightnessOn(fromX, fromY, toX, toY int) {
	for i := fromX; i <= toX; i++ {
		for j := fromY; j <= toY; j++ {
			brightness[i][j] = brightness[i][j] + 1
		}
	}
}

func turnLightsOff(fromX, fromY, toX, toY int) {
	for i := fromX; i <= toX; i++ {
		for j := fromY; j <= toY; j++ {
			lights[i][j] = off
		}
	}
}

func turnBrightnessOff(fromX, fromY, toX, toY int) {
	for i := fromX; i <= toX; i++ {
		for j := fromY; j <= toY; j++ {
			brightness[i][j] = brightness[i][j] - 1
			if brightness[i][j] < 0 {
				brightness[i][j] = 0
			}
		}
	}
}

func toggleLights(fromX, fromY, toX, toY int) {
	for i := fromX; i <= toX; i++ {
		for j := fromY; j <= toY; j++ {
			if lights[i][j] == off {
				lights[i][j] = on
			} else {
				lights[i][j] = off
			}
		}
	}
}

func toggleBrightness(fromX, fromY, toX, toY int) {
	for i := fromX; i <= toX; i++ {
		for j := fromY; j <= toY; j++ {
			brightness[i][j] = brightness[i][j] + 2
		}
	}
}
