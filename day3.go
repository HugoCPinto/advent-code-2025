package main

import ("fmt"; "os"; "bufio")

func readInput(path string) [][]int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var matrix [][]int
	scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()

        row := make([]int, 0, len(line))
        for _, ch := range line {
            row = append(row, int(ch-'0'))
        }

        matrix = append(matrix, row)
    }
	return matrix
}

func getLargestPossibleJoltage(matrix [][]int) []int {

	var banksLatestHigherJoltage []int
    for i := 0; i < len(matrix); i++ {
        row := matrix[i]
        maxJoltage := 0

        for j := 0; j < len(row)-1; j++ {
            for k := j + 1; k < len(row); k++ {
                value := row[j]*10 + row[k]
                if value > maxJoltage {
                    maxJoltage = value
                }
            }
        }

        banksLatestHigherJoltage = append(banksLatestHigherJoltage, maxJoltage)
    }
	return banksLatestHigherJoltage
}

func getLargestJoltage(matrix [][]int, length int) []int {
    var resultJoltages []int

    for _, row := range matrix {
        stack := []int{}
        for i, digit := range row {
            for len(stack) > 0 && stack[len(stack)-1] < digit && len(stack)-1+len(row)-i >= length {
                stack = stack[:len(stack)-1]
            }
            if len(stack) < length {
                stack = append(stack, digit)
            }
        }
        resultJoltages = append(resultJoltages, combineDigits(stack))
    }

    return resultJoltages
}

func combineDigits(digits []int) int {
    result := 0
    for _, d := range digits {
        result = result*10 + d
    }
    return result
}

func calculateSum(ids []int) int {
	sum := 0
	for _, id := range ids {
		sum += id
	}
	return sum
}

func main() {
	fmt.Println("Day 3")

	matrix := readInput("inputs/day3-1")
	fmt.Println(matrix)
	largestJoltage := getLargestJoltage(matrix, 12)
	fmt.Println("Largest possible joltage:", largestJoltage)
	sum := calculateSum(largestJoltage)
	fmt.Println("Sum of largest possible joltage:", sum)

}