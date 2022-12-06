package day6

import "strconv"

type Solver struct {}

func (*Solver) SolvePart1(input string, extraParams ...any) string {
	return strconv.Itoa(findPacketStart(input))
}

func findPacketStart(input string) int {
	for i := 3; i < len(input); i++ {
		if allDiff(input[i-3:i+1]) {
			return i + 1
		}
	}
	panic("No start found")		
}

func allDiff(input string) bool {
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			if input[i] == input[j] {
				return false
			}
		}
	}
	return true
}

func (*Solver) SolvePart2(input string, extraParams ...any) string {
	return strconv.Itoa(findMessageStart(input))
}

func findMessageStart(input string) int {
	for i := 13; i < len(input); i++ {
		if allDiff(input[i-13:i+1]) {
			return i + 1
		}
	}
	panic("No start found")		
}
