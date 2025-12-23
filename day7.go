package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInput(path string) [][]string {

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var matrix [][]string
    for scanner.Scan() {
        line := scanner.Text()

        row := make([]string, 0, len(line))
        for _, ch := range line {
            row = append(row, string(ch))
        }

        matrix = append(matrix, row)
    }

	return matrix
}

func processMatrix(matrix [][]string, startPosition [2]int) {
	currentRow := startPosition[0]
	currentCol := startPosition[1]


	for currentRow < len(matrix) && currentCol < len(matrix[0]) {

        if matrix[currentRow][currentCol] == "|" {
            return
        }

		if matrix[currentRow][currentCol] == "."{
			matrix[currentRow][currentCol] = "|"
		}

		if matrix[currentRow][currentCol] == "^"{
			
            if currentCol-1 >= 0 && matrix[currentRow][currentCol-1] == "." {
                processMatrix(matrix, [2]int{currentRow, currentCol - 1})
            }

            if currentCol+1 < len(matrix[0]) && matrix[currentRow][currentCol+1] == "." {
                processMatrix(matrix, [2]int{currentRow, currentCol + 1})
            }

			return
						
		}

		currentRow++
	}

/*
	fmt.Println("Matrix State: ")
	for _, row := range matrix {
		fmt.Println(row)
	}
*/
	return
}

func countValidSplits(matrix [][]string) int {
    total := 0
    rows := len(matrix)
    if rows == 0 {
        return 0
    }
    cols := len(matrix[0])

    for i := 1; i < rows; i++ {
        for j := 1; j < cols-1; j++ {
            if matrix[i][j] == "^" {
                if matrix[i-1][j] == "|" {
                    if matrix[i][j-1] == "|" && matrix[i][j+1] == "|" {
                        total++
                    }
                }
            }
        }
    }

    return total
}

func main() {
	fmt.Println("Day 7")

	matrix := readInput("inputs/day-7-1")

	startPosition := [2]int{}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == "S" {
				startPosition = [2]int{i, j}
				break
			}
		}
	}

	startPosition[0]++ // Move down one row to start processing
	processMatrix(matrix, startPosition)
	fmt.Println("Total energized cells:", countValidSplits(matrix))
	/*
	fmt.Println("Matrix State: ")
	for _, row := range matrix {
		fmt.Println(row)
	}
	*/


}