package main

import ("fmt"; "os"; "bufio"; "strings"; "strconv"; "sort")

type Range struct {
    Start int
    End   int
}

func readInput(path string) ([]Range, []int) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	ranges := make([]Range, 0)
	list := make([]int, 0)
	afterBlank := false

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			afterBlank = true
			continue
		}

		if !afterBlank {
			parts := strings.Split(line, "-")
			start, _ := strconv.Atoi(parts[0])
			end, _ := strconv.Atoi(parts[1])

			ranges = append(ranges, Range{Start: start, End: end})
		} else {
			val, _ := strconv.Atoi(line)
			list = append(list, val)
		}
	}

	return ranges, list
}

func processAvailableFreshIngredients(ranges []Range, ingredients []int) []int {
	seen := make(map[int]bool)
	result := make([]int, 0)

	for _, i := range ingredients {
		if seen[i] {
			continue
		}

		for _, r := range ranges {
			if i >= r.Start && i <= r.End {
				seen[i] = true
				result = append(result, i)
				break
			}
		}
	}

	return result
}

func countFreshIngredients(ranges []Range) int {
	if len(ranges) == 0 {
		return 0
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})

	total := 0
	currentStart := ranges[0].Start
	currentEnd := ranges[0].End

	for i := 1; i < len(ranges); i++ {
		r := ranges[i]

		if r.Start <= currentEnd+1 {
			if r.End > currentEnd {
				currentEnd = r.End
			}
		} else {
			total += currentEnd - currentStart + 1
			currentStart = r.Start
			currentEnd = r.End
		}
	}

	total += currentEnd - currentStart + 1

	return total
}


func main() {
	fmt.Println("Day 5")

	ranges, ingredients := readInput("inputs/day-5-1")
	availableFreshIngredients := processAvailableFreshIngredients(ranges, ingredients)
	count := countFreshIngredients(ranges)
	fmt.Println("Total available fresh ingredients:", len(availableFreshIngredients))
	fmt.Println("Total fresh ingredients in ranges:", count)
}