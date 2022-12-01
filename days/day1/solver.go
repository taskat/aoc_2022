package day1

import (
	"sort"
	"strconv"
	"strings"
)

type Solver struct{}

func (*Solver) SolvePart1(input string, extraParams ...any) string {
	elves := getElves(input)
	return strconv.Itoa(findMaxElf(elves))
}

func getElves(input string) []int {
	chunks := strings.Split(input, "\n\n")
	elves := make([]int, len(chunks))
	for i, chunk := range chunks {
		lines := strings.Split(chunk, "\n")
		sum := 0
		for _, line := range lines {
			food, _ := strconv.Atoi(line)
			sum += food
		}
		elves[i] = sum
	}
	return elves
}

func findMaxElf(elves []int) int {
	max := elves[0]
	for _, elf := range elves {
		if elf > max {
			max = elf
		}
	}
	return max
}

func (*Solver) SolvePart2(input string, extraParams ...any) string {
	elves := getElves(input)
	return strconv.Itoa(findTop3Elves(elves))
}

func findTop3Elves(elves []int) int {
	sort.Ints(elves)
	return elves[len(elves)-3] + elves[len(elves)-2] + elves[len(elves)-1]
}
