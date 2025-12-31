package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Coordinate struct {
	Width, Height int
}

type Present struct {
	Id    int
	Shape []Coordinate
}

type Region struct {
	Width, Height int
	PresentsID    []int
}

type Grid struct {
	Width, Height int
	Cells         [][]string
}

func readInput(path string) ([]Present, []Region) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	presents := []Present{}
	regions := []Region{}

	re := regexp.MustCompile(`^\d+:$`)
	regionRe := regexp.MustCompile(`^(\d+)x(\d+):\s*(.*)$`)

	var current *Present
	y := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		if re.MatchString(line) {
			id, _ := strconv.Atoi(strings.TrimSuffix(line, ":"))
			presents = append(presents, Present{Id: id})
			current = &presents[len(presents)-1]
			y = 0
			continue
		}

		if current != nil && !regionRe.MatchString(line) {
			for x, ch := range line {
				if ch == '#' {
					current.Shape = append(current.Shape, Coordinate{Width: y, Height: x})
				}
			}
			y++
		}

		if m := regionRe.FindStringSubmatch(line); m != nil {
			h, _ := strconv.Atoi(m[1])
			w, _ := strconv.Atoi(m[2])
			fields := strings.Fields(m[3])
			nums := []int{}
			for _, f := range fields {
				n, _ := strconv.Atoi(f)
				nums = append(nums, n)
			}
			regions = append(regions, Region{Width: w, Height: h, PresentsID: nums})
		}
	}
	return presents, regions
}

func createDefaultGrid(width int, height int) Grid {
	grid := Grid{
		Width:  width,
		Height: height,
		Cells:  make([][]string, width),
	}
	for y := 0; y < width; y++ {
		grid.Cells[y] = make([]string, height)
		for x := 0; x < height; x++ {
			grid.Cells[y][x] = "."
		}
	}
	return grid
}

func printGrid(grid Grid) {
	for y := 0; y < grid.Width; y++ {
		for x := 0; x < grid.Height; x++ {
			fmt.Print(grid.Cells[y][x])
		}
		fmt.Println()
	}
	fmt.Println()
}

func fillPresentGrid(grid Grid, present Present, symbol string) Grid {
	for _, c := range present.Shape {
		if c.Width >= 0 && c.Width < grid.Width && c.Height >= 0 && c.Height < grid.Height {
			grid.Cells[c.Width][c.Height] = symbol
		}
	}
	return grid
}

func gridFit(grid Grid, present Present) bool {
	for _, c := range present.Shape {
		if c.Width < 0 || c.Width >= grid.Width || c.Height < 0 || c.Height >= grid.Height {
			return false
		}
		if grid.Cells[c.Width][c.Height] != "." {
			return false
		}
	}
	return true
}

func boundingBox(coords []Coordinate) (minX, minY, maxX, maxY int) {
	minX, minY = coords[0].Height, coords[0].Width
	maxX, maxY = minX, minY
	for _, c := range coords {
		if c.Width < minX {
			minX = c.Width
		}
		if c.Width > maxX {
			maxX = c.Width
		}
		if c.Height < minY {
			minY = c.Height
		}
		if c.Height > maxY {
			maxY = c.Height
		}
	}
	return
}

func shiftPresent(p Present, dx, dy int) Present {
	newShape := make([]Coordinate, len(p.Shape))
	for i, c := range p.Shape {
		newShape[i] = Coordinate{Width: c.Width + dy, Height: c.Height + dx}
	}
	return Present{Id: p.Id, Shape: newShape}
}

func normalizePresent(p Present) Present {
	minX, minY, _, _ := boundingBox(p.Shape)
	return shiftPresent(p, -minX, -minY)
}

// Rotate 90° clockwise
func rotatePresent90(p Present) Present {
    //p = normalizePresent(p)
    newShape := make([]Coordinate, len(p.Shape))
    _, _, maxX, _ := boundingBox(p.Shape)
    for i, c := range p.Shape {
        newShape[i] = Coordinate{
            Width:  c.Height,
            Height: maxX - c.Width,
        }
    }
    return Present{Id: p.Id, Shape: newShape}
}

func fitPresent(grid Grid, present Present) (Present, bool) {
    //present = normalizePresent(present)

	ok := gridFit(grid, present)
	if ok {
		return present, true
	}

	// this logic is not working
	for r := 0; r < 4; r++ {
		_, _, maxX, maxY := boundingBox(present.Shape)
		boxWidth := maxX + 1
		boxHeight := maxY + 1

        for y := 0; y <= grid.Width-boxWidth; y++ {
            for x := 0; x <= grid.Height-boxHeight; x++ {
				candidate := shiftPresent(present, x, y)
				if gridFit(grid, candidate) {
					return candidate, true
				}
			}
		}

		present = rotatePresent90(present)		
    }

    return present, false
}

func getPresent(presents []Present, presentId int) Present {
	for _, p := range presents {
		if p.Id == presentId {
			return p
		}
	}
	panic("Present not found")
}

func processInput(presents []Present, regions []Region) int {
	total := 0
	for _, region := range regions {
		grid := createDefaultGrid(region.Width, region.Height)
		fmt.Println("Initial Grid:")
		printGrid(grid)

		symbol := 'A'

		for i := 0; i < len(region.PresentsID); i++ {
			if region.PresentsID[i] == 0 {
				continue
			}

			presentIndex := i
			numShapes := region.PresentsID[i]
			present := getPresent(presents, presentIndex)

			for j := 0; j < numShapes; j++ {
				candidate, ok := fitPresent(grid, present)
				if ok {
					grid = fillPresentGrid(grid, candidate, string(symbol))
					fmt.Printf("Placed Present %d (%c):\n", present, symbol)
					printGrid(grid)
				} else {
					fmt.Printf("Could not fit present %d (%c)\n", present.Id, symbol)
					break
				}
				symbol++
			}
			total++
		}
	}
	return total
}

func main() {
	fmt.Println("Day 12")
	presents, regions := readInput("inputs/day12-1-example.txt")
	fmt.Println(presents)
	total := processInput(presents, regions)
	fmt.Println("Total placed:", total)
}
