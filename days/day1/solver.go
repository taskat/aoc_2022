package day1

import (
	"strconv"
	"strings"
)

type Solver struct {}

func (*Solver) SolvePart1(input string, extraParams ...any) string {
	depths := getDepths(input)
	return strconv.Itoa(countIncreases(depths))
}

func getDepths(input string) []int {
	lines := strings.Split(input, "\n")
	depths := make([]int, len(lines))
	for i, line := range lines {
		depths[i], _ = strconv.Atoi(line)
	}
	return depths
}

func countIncreases(depths []int) int {
	count := 0
	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i-1] {
			count++
		}
	}
	return count
}

func (*Solver) SolvePart2(input string, extraParams ...any) string {
	depths := getDepths(input)
	return strconv.Itoa(count3increases(depths))
}

func count3increases(depths []int) int {
	count := 0
	for i := 3; i < len(depths); i++ {
		if depths[i] + depths[i - 1] + depths[i - 2] > depths[i - 1] + depths[i - 2] + depths[i - 3] {
			count++
		}
	}
	return count
}
