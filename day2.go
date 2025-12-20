package main

import ("fmt"; "os"; "bufio"; "strings"; "strconv")

func readInput(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var idRanges []string
	for scanner.Scan() {
		line := scanner.Text()
		idRanges = append(idRanges, strings.Split(line, ",")...)
	}
	
	return idRanges
}

func isRepeatedNumberAtLeastTwice(n int) bool {
    s := strconv.Itoa(n)
    length := len(s)

    for l := 1; l <= length/2; l++ {
        if length%l != 0 {
            continue
        }

        seq := s[:l]
        times := length / l
        repeated := strings.Repeat(seq, times)

        if repeated == s {
            return true
        }
    }

    return false
}

func isRepeatedNumber(n int) bool {
    s := strconv.Itoa(n)
    length := len(s)

    if length%2 != 0 {
        return false
    }

    half := length / 2
    firstHalf := s[:half]
    secondHalf := s[half:]

    return firstHalf == secondHalf
}


func findInvalidIDs(ranges []string) []int {

	invalidIDs := []int{}
	for _, r := range ranges {
		bounds := strings.Split(r, "-")
		if len(bounds) != 2 {
			continue
		}

		startBound, _ := strconv.Atoi(bounds[0])
		endBound, _   := strconv.Atoi(bounds[1])

		for id := startBound; id <= endBound; id++ {
			//fmt.Println("Checking ID:", id)
			isRepeatedNumberDigit := isRepeatedNumberAtLeastTwice(id)
			if isRepeatedNumberDigit {
				fmt.Println("Invalid ID found:", id)
				invalidIDs = append(invalidIDs, id)
			}
		}

	}

	return invalidIDs
}

func calculateSum(ids []int) int {
	sum := 0
	for _, id := range ids {
		sum += id
	}
	return sum
}

func main(){
	fmt.Println("Day 2")

	ranges := readInput("inputs/day2-1")
	invalids := findInvalidIDs(ranges)
	sum := calculateSum(invalids)
	fmt.Println("Sum of invalid IDs:", sum)
}