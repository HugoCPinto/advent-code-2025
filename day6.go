package main

import ("fmt";"os"; "bufio"; "strings"; "strconv")

func readInput(path string) ([][]int, []rune) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	numbers := [][]int{}
	var operators []rune

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		if strings.ContainsAny(line, "*+-/") {
			fields := strings.Fields(line)
			for _, f := range fields {
				operators = append(operators, rune(f[0]))
			}
			continue
		}

		fields := strings.Fields(line)
		row := make([]int, len(fields))
		for i, f := range fields {
			row[i], _ = strconv.Atoi(f)
		}
		numbers = append(numbers, row)

	}

	return numbers, operators
}

func evaluateProblem(numbers [][]int, operators []rune) int {
	grandTotal := 0

	for i, operator := range operators {
		columnTotal := numbers[0][i]

		for row := 1; row < len(numbers); row++ {
			switch operator {
			case '+':
				columnTotal += numbers[row][i]
			case '*':
				columnTotal *= numbers[row][i]
			}
		}

		grandTotal += columnTotal
	}

	return grandTotal
}

func main(){
	fmt.Println("Day 6")

	numbers, operators := readInput("inputs/day-6-1")
	fmt.Println("Numbers:", numbers)
	fmt.Println("Operators:", operators)
	result := evaluateProblem(numbers, operators)
	fmt.Println("Result:", result)
}