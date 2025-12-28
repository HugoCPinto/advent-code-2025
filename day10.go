package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Machine struct {
	Diagram string
	Buttons []string
	Joltage []int
}

func readInput(path string) []Machine {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reDiagram := regexp.MustCompile(`\[(.*?)\]`)
	reButtons := regexp.MustCompile(`\(([^)]+)\)`)
	reJoltage := regexp.MustCompile(`\{([^}]+)\}`)

	machines := []Machine{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		diagramMatch := reDiagram.FindStringSubmatch(line)
		if len(diagramMatch) < 2 {
			continue
		}
		diagram := diagramMatch[1]

		buttonMatches := reButtons.FindAllStringSubmatch(line, -1)
		buttons := []string{}
		for _, m := range buttonMatches {
			buttons = append(buttons, m[1])
		}

		jMatch := reJoltage.FindStringSubmatch(line)
		joltage := []int{}
		if len(jMatch) > 1 {
			parts := strings.Split(jMatch[1], ",")
			for _, p := range parts {
				v, err := strconv.Atoi(strings.TrimSpace(p))
				if err != nil {
					panic(err)
				}
				joltage = append(joltage, v)
			}
		}

		machine := Machine{
			Diagram: diagram,
			Buttons: buttons,
			Joltage: joltage,
		}

		machines = append(machines, machine)
	}

	return machines
}

func getDiagramLightPostions(diagram string) []int{
	positions := []int{}
	for i,x := range strings.Split(diagram, "") {
		if x == "#" {
			positions = append(positions, i)
		}
	}
	return positions
}

func combinations(items []string) [][]string {
	n := len(items)
	result := [][]string{}

	for mask := 1; mask < (1 << n); mask++ {
		comb := []string{}
		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				comb = append(comb, items[i])
			}
		}
		result = append(result, comb)
	}

	return result
}

func matchesDiagram(state []bool, diagramLights []int) bool {
	for i := 0; i < len(state); i++ {
		shouldBeOn := false
		for _, d := range diagramLights {
			if d == i {
				shouldBeOn = true
				break
			}
		}
		if state[i] != shouldBeOn {
			return false
		}
	}
	return true
}

func processInput(machines []Machine) int {
	
	total := 0

	for _,m := range machines {		
		diagramLights := getDiagramLightPostions(m.Diagram)

		minPresses := math.MaxInt
		//var minCombo [][]int

		buttonsPositions := []([]int){}
		for _, button := range m.Buttons {
			pos := []int{}
			for _, p := range strings.Split(button, ",") {
				v, _ := strconv.Atoi(p)
				pos = append(pos, v)
			}
			buttonsPositions = append(buttonsPositions, pos)
		}

		n := len(buttonsPositions)
		for mask := 1; mask < (1 << n); mask++ {
			combo := [][]int{}
			for i := 0; i < n; i++ {
				if mask&(1<<i) != 0 {
					combo = append(combo, buttonsPositions[i])
				}
			}

			state := make([]bool, len(m.Diagram))
			for _, btn := range combo {
				for _, p := range btn {
					state[p] = !state[p]
				}
			}

			if matchesDiagram(state, diagramLights) {
				if len(combo) < minPresses {
					minPresses = len(combo)
					//minCombo = combo
				}
			}
		}

		//fmt.Println(minCombo, minPresses)
		total += minPresses
	}
	
	return total
}

func main(){
	fmt.Println("Day 10")

	machines := readInput("inputs/day-10-1")
	//fmt.Println(machines)
	total := processInput(machines)
	fmt.Println("total: ", total)
}