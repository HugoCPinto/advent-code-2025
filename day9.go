package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
    column   int
	line int    
}

func readInput(path string) []Point {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	points := make([]Point, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		parts := strings.Split(line, ",")
		c, _ := strconv.Atoi(parts[0])
		l, _ := strconv.Atoi(parts[1])
		points = append(points, Point{line: l, column: c})
	}

	return points
}

func validateRectangle(p1 Point, p2 Point) bool {
	return p1.line != p2.line && p1.column != p2.column
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func calculateArea(p1 Point, p2 Point) int {
	width := abs(p2.column - p1.column) + 1
	height := abs(p2.line - p1.line) + 1

	return width * height
}

func processInput(input []Point) int{

	largestArea := 0
	for _,point1 := range input {

		for _,point2 := range input {

			if point1.line == point2.line && point1.column == point2.column {
				continue
			}

			valid := validateRectangle(point1, point2)
			if valid {
				calculatedArea := calculateArea(point1, point2)
				//fmt.Println("Points: ", point1, point2, " calculated area: ", calculatedArea)

				if calculatedArea > largestArea {
					largestArea = calculatedArea
				}
			}
		}

	}

	return largestArea

}

func main(){
	fmt.Println("Day 9")

	input := readInput("inputs/day-9-1")
	//fmt.Println("Points: ", input)
	area := processInput(input)
	fmt.Println("largest area: ", area)
}