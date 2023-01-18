package day25

import (
	"strings"
)

type Solver struct {}

func (*Solver) SolvePart1(input string, extraParams ...any) string {
	snafus := parseInput(input)
	decs := convertToDec(snafus)
	sum := sum(decs)
	result := fromDec(sum)
	return result.String()
}

var digits = map[rune]int{'=': -2, '-': -1, '0': 0, '1': 1, '2': 2}
var inverseDigits = map[int]rune{-2: '=', -1: '-', 0: '0', 1: '1', 2: '2'}

type snafu []rune

func parseSnafu(line string) snafu {
	return snafu(line)
}

func fromDec(dec int) snafu {
	s := make(snafu, 0)
	digit := 0
	carry := false
	for dec > 0 {
		digit = dec % 5
		if digit > 2 {
			digit -= 5
		}
		carry = digit < 0
		s = append([]rune{inverseDigits[digit]}, s...)
		dec /= 5
		if carry {
			dec++
		}
	}	
	return s
}

func (s snafu) String() string {
	return string(s)
}

func (s snafu) toDec() int {
	var dec int
	multiplier := 1
	for i := len(s) - 1; i >= 0; i-- {
		dec += digits[s[i]] * multiplier
		multiplier *= 5
	}
	return dec
}

func parseInput(input string) []snafu {
	lines := strings.Split(input, "\n")
	snafus := make([]snafu, len(lines))
	for i, line := range lines {
		snafus[i] = parseSnafu(line)
	}
	return snafus
}

func convertToDec(snafus []snafu) []int {
	dec := make([]int, len(snafus))
	for i, s := range snafus {
		dec[i] = s.toDec()
	}
	return dec
}

func sum(arr []int) int {
	var sum int
	for _, i := range arr {
		sum += i
	}
	return sum
}

func (*Solver) SolvePart2(input string, extraParams ...any) string {
	return ""
}
