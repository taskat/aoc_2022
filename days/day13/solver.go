package day13

import (
	"aoc_2022/days/day13/integer"
	"aoc_2022/days/day13/list"
	"aoc_2022/days/day13/value"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Solver struct {}

func (*Solver) SolvePart1(input string, extraParams ...any) string {
	pairs := parsePairs(input)	
	return strconv.Itoa(sumCorrectPairs(pairs))
}

type pair struct {
	left, right value.Value
}

func parsePairs(input string) []pair {
	chunks := strings.Split(input, "\n\n")
	pairs := make([]pair, len(chunks))
	for i, chunk := range chunks {
		pairs[i] = parsePair(chunk)
	}
	return pairs
}

func parsePair(input string) pair {
	lines := strings.Split(input, "\n")
	first := parseValue(&lines[0])
	second := parseValue(&lines[1])
	return pair{first, second}
}

func parseValue(data *string) value.Value {
	if (*data)[0] != '[' {
		return parseInteger(data)
	}
	*data = (*data)[1:]
	var values []value.Value
	for (*data)[0] != ']' {
		values = append(values, parseValue(data))
		if (*data)[0] == ',' {
			*data = (*data)[1:]
		}
	}
	*data = (*data)[1:]
	return list.NewList(values...)
}

func parseInteger(data *string) integer.Integer {
	re := regexp.MustCompile(`^(\d+)`)
	number := re.FindString(*data)
	*data = (*data)[len(number):]
	i, _ := strconv.Atoi(number)
	return integer.NewInteger(i)
}

func sumCorrectPairs(pairs []pair) int {
	sum := 0
	for i, pair := range pairs {
		if pair.left.Compare(pair.right) == value.Right {
			sum += i + 1
		}
	}
	return sum
}

func (*Solver) SolvePart2(input string, extraParams ...any) string {
	input += "\n[[2]]"
	input += "\n[[6]]"
	values := parsePackets(input)
	sort.Sort(packets(values))
	first, second := values.findDividers()
	return strconv.Itoa(first * second)
}

func parsePackets(input string) packets {
	lines := strings.Split(removeEmptyLines(input), "\n")
	values := make(packets, len(lines))
	for i, line := range lines {
		values[i] = parseValue(&line)
	}
	return values
}

type packets []value.Value

func (p packets) Len() int {
	return len(p)
}

func (p packets) Less(i, j int) bool {
	return p[i].Compare(p[j]) == value.Right
}

func (p packets) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p packets) findDividers() (int, int) {
	var first, second int
	for i, packet := range p {
		if packet.String() == "[[2]]" {
			first = i + 1
		}
		if packet.String() == "[[6]]" {
			second = i + 1
		}
	}
	return first, second
}

func removeEmptyLines(input string) string {
	return strings.ReplaceAll(input, "\n\n", "\n")
}
