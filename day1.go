package main

import ("fmt"; "os"; "bufio"; "strings"; "strconv")

func readInput(path string) []string {
    file, err := os.Open(path)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    var directions []string
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()
        directions = append(directions, line)
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

	return directions
}

func calculateZeros(directions []string) int {
	var startingPosition int = 50

	totalZeros := 0
	for _, direction := range directions {

		num,_ := strconv.Atoi(direction[1:]) 
		if strings.HasPrefix(direction, "L") {
			startingPosition -= num
		}else{
			startingPosition += num
		}
		
        startingPosition = ((startingPosition % 100) + 100) % 100

		if startingPosition == 0 {
			totalZeros++
		}
	}
	return totalZeros
}

func calculateAllZeros(directions []string) int {
	var startingPosition int = 50

	totalZeros := 0
	for _, direction := range directions {

		num,_ := strconv.Atoi(direction[1:])
		for i := 0; i < num; i++ {
            
			if direction[0] == 'L' {
                startingPosition--
            } else {
                startingPosition++
            }

            startingPosition = (startingPosition + 100) % 100

            if startingPosition == 0 {
                totalZeros++
            }
        }
	}
	return totalZeros
}


func main(){
	fmt.Println("Day 1")
	
	input := readInput("inputs/day1-1-example")	
	totalZeros := calculateZeros(input)
	fmt.Printf("1 - Example total Zeros: %d\n", totalZeros)

	input = readInput("inputs/day1-1")	
	totalZeros = calculateZeros(input)
	fmt.Printf("1 - Total Zeros: %d\n", totalZeros)

	input = readInput("inputs/day1-1-example")	
	totalZeros = calculateAllZeros(input)
	fmt.Printf("2 - Example total Zeros: %d\n", totalZeros)

	input = readInput("inputs/day1-1")	
	totalZeros = calculateAllZeros(input)
	fmt.Printf("2 - Total Zeros: %d\n", totalZeros)
}	