package main

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

func main() {

}
