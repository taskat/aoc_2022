package day4

import (
	"strconv"
	"strings"
)

type Solver struct {}

func (*Solver) SolvePart1(input string, extraParams ...any) string {
	pairs := createPairs(input)
	return strconv.Itoa(count(pairs, Range.fullyContains))
}

type Range struct {
	from int
	to int
}

func parseRange(s string) Range {
	parts := strings.Split(s, "-")
	from, _ := strconv.Atoi(parts[0])
	to, _ := strconv.Atoi(parts[1])
	return Range{from: from, to: to}
}

func (r Range) fullyContains(other Range) bool {
	return r.from <= other.from && r.to >= other.to
}

func (r Range) overlaps(other Range) bool {
	return r.from >= other.from && r.from <= other.to ||
		r.to >= other.from && r.to <= other.to
}

type pair struct {
	elves [2]Range
}

func parsePair(line string) pair {
	parts := strings.Split(line, ",")
	elves := [2]Range{parseRange(parts[0]), parseRange(parts[1])}
	return pair{elves: elves}
}

func createPairs(input string) []pair {
	lines := strings.Split(input, "\n")
	pairs := make([]pair, len(lines))
	for i, line := range lines {
		pairs[i] = parsePair(line)
	}
	return pairs
}

func count(pairs []pair, f func(Range, Range) bool) int {
	count := 0
	for _, pair := range pairs {
		if f(pair.elves[0], pair.elves[1]) || f(pair.elves[1], pair.elves[0]) {
			count++
		}
	}
	return count
}

func (*Solver) SolvePart2(input string, extraParams ...any) string {
	pairs := createPairs(input)
	return strconv.Itoa(count(pairs, Range.overlaps))
}
