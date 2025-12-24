package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

type Circuit struct {
	From     []int
	To       []int
	JBoxes	 [][]int
	Distance float64
}

func readInput(path string) [][]int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var numbers [][]int
	for scanner.Scan() {
		line := scanner.Text()
		var row []int
		var num int
		for _, ch := range line {
			if ch >= '0' && ch <= '9' {
				num = num*10 + int(ch-'0')
			} else if ch == ',' {
				row = append(row, num)
				num = 0
			}
		}
		row = append(row, num)
		numbers = append(numbers, row)
	}

	return numbers
}

func calculateDistance(jBox1, jBox2 []int) float64 {
	dx := jBox1[0] - jBox2[0]
	dy := jBox1[1] - jBox2[1]
	dz := jBox1[2] - jBox2[2]
	return math.Sqrt(float64(dx*dx + dy*dy + dz*dz))
}

func compareJBoxs(jBox1, jBox2 []int) bool {
	return jBox1[0] == jBox2[0] && jBox1[1] == jBox2[1] && jBox1[2] == jBox2[2]
}

func jBoxExists(circuits []Circuit, jBox []int) bool {
	for _, c := range circuits {
		if compareJBoxs(c.From, jBox) || compareJBoxs(c.To, jBox) {
			return true
		}
	}
	return false
}

func containsJBox(list [][]int, jBox []int) bool {
	for _, jb := range list {
		if compareJBoxs(jb, jBox) {
			return true
		}
	}
	return false
}

func circuitContainsAny(c Circuit, jBoxes ...[]int) bool {
	for _, jb := range c.JBoxes {
		for _, target := range jBoxes {
			if compareJBoxs(jb, target) {
				return true
			}
		}
	}
	return false
}

func mergeCircuits(circuits []Circuit) []Circuit {
	changed := true

	for changed {
		changed = false
		result := make([]Circuit, 0)

		for _, c := range circuits {
			merged := false

			for i := range result {
				if circuitContainsAny(result[i], c.JBoxes...) {

					for _, jb := range c.JBoxes {
						if !containsJBox(result[i].JBoxes, jb) {
							result[i].JBoxes = append(result[i].JBoxes, jb)
						}
					}

					result[i].Distance += c.Distance
					merged = true
					changed = true
					break
				}
			}

			if !merged {
				result = append(result, c)
			}
		}

		circuits = result
	}

	return circuits
}

func collectUsedJBoxes(circuits []Circuit) map[string]bool {
	used := make(map[string]bool)
	for _, c := range circuits {
		for _, jb := range c.JBoxes {
			used[fmt.Sprint(jb)] = true
		}
	}
	return used
}

func processJBoxs(jBoxs [][]int) {

	circuits := make([]Circuit, 0)

	for i := 0; i < len(jBoxs); i++ {
		for j := i + 1; j < len(jBoxs); j++ {
			jBox1 := jBoxs[i]
			jBox2 := jBoxs[j]
			distance := calculateDistance(jBox1, jBox2)
			circuits = append(circuits, Circuit{
				From:     jBox1,
				To:       jBox2,
				JBoxes:   [][]int{jBox1, jBox2},
				Distance: distance,
			})
		}
	}

	sort.Slice(circuits, func(i, j int) bool {
		return circuits[i].Distance < circuits[j].Distance
	})

	if len(circuits) > 1000 {
		circuits = circuits[:1000]
	}

	/*
	for _, c := range circuits {
		fmt.Printf("From %v to %v → distance: %.2f\n", c.From, c.To, c.Distance)
	}
	*/

	mergedCircuits := mergeCircuits(circuits)

	used := collectUsedJBoxes(mergedCircuits)
	for _, jb := range jBoxs {
		if !used[fmt.Sprint(jb)] {
			mergedCircuits = append(mergedCircuits, Circuit{
				JBoxes:   [][]int{jb},
				Distance: 0,
			})
		}
	}

	/*
	fmt.Println("\nMerged circuits:")
	for _, c := range mergedCircuits {
		fmt.Println(c.JBoxes)
	}
	*/

	var sizes []int
	for _, c := range mergedCircuits {
		sizes = append(sizes, len(c.JBoxes))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

	result := sizes[0] * sizes[1] * sizes[2]
	fmt.Println("Result:", result)
}


func main() {
	fmt.Println("Day 8")

	jBoxes := readInput("inputs/day-8-1")
	processJBoxs(jBoxes)
}