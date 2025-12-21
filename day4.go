package main

import ("fmt"; "os"; "bufio")

func readInput(path string) [][]string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var matrix [][]string
	scanner := bufio.NewScanner(file)
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

func copyMatrix(src [][]string) [][]string {
	dst := make([][]string, len(src))
	for i := range src {
		dst[i] = make([]string, len(src[i]))
		copy(dst[i], src[i])
	}
	return dst
}

func processRolls(matrix [][]string) [][]string {

	result := copyMatrix(matrix)
	for i := 0; i < len(matrix); i++ {
		
		for j := 0; j < len(matrix); j++ {
			totalRolls := 0
			result[i][j] = matrix[i][j]
			if matrix[i][j] == "." {			
				continue
			}

			// try up
			if i-1 >= 0 && matrix[i-1][j] == "@" {
				totalRolls++
			}

			// try up-left
			if i-1 >= 0 && j-1 >= 0 && matrix[i-1][j-1] == "@" {
				totalRolls++
			}

			// try up-right
			if i-1 >= 0 && j+1 < len(matrix[i]) && matrix[i-1][j+1] == "@" {
				totalRolls++
			}

			// try left
			if j-1 >= 0 && matrix[i][j-1] == "@" {
				totalRolls++
			}

			// try right
			if j+1 < len(matrix[i]) && matrix[i][j+1] == "@" {
				totalRolls++
			}

			// try down-left
			if i+1 < len(matrix[i]) && j-1 >= 0 && matrix[i+1][j-1] == "@" {
				totalRolls++
			}

			// try down-right
			if i+1 < len(matrix[i]) && j+1 < len(matrix[i]) && matrix[i+1][j+1] == "@" {
				totalRolls++
			}

			// try down
			if i+1 < len(matrix[i]) && matrix[i+1][j] == "@" {
				totalRolls++
			}

			//fmt.Println("Total rolls for", i, j, ":", totalRolls)

			if totalRolls < 4 {
				result[i][j] = "X"
			}

		}
	}
	return result
}

func countProcessed(matrix [][]string) int {
	count := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == "X" {
				count++
			}
		}
	}
	return count
}

func removeRollsProcessed(matrix [][]string) [][]string {
	result := copyMatrix(matrix)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == "X" {
				result[i][j] = "."
			}
		}
	}
	return result
}

func main() {
	fmt.Println("Day 4")

	matrix := readInput("inputs/day-4-1")
	//fmt.Println("Matrix:", matrix)
	processed := processRolls(matrix)
	//fmt.Println("Processed:", processed)
	count := countProcessed(processed)
	
	// second part
	totalCount := count

	for count > 0 {
		matrix = removeRollsProcessed(processed)
		processed = processRolls(matrix)
		count = countProcessed(processed)
		totalCount += count
	}
	
	fmt.Println("Count of processed:", totalCount)


}