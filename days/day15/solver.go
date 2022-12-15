package day15

import (
	"strconv"
	"strings"
)

type Solver struct{}

func (*Solver) SolvePart1(input string, extraParams ...any) string {
	lineY := getLineNumber(extraParams)
	pairs := parsePairs(input)
	ranges := getRanges(pairs, lineY)
	ranges = mergeRanges(ranges)
	invalid := calculateInvalidPositions(ranges)
	occupied := calculateOccupiedPositions(pairs, ranges, lineY)
	return strconv.Itoa(invalid - occupied)
}

func getLineNumber(extraParams []any) int {
	if len(extraParams) > 0 {
		line, _ := strconv.Atoi(extraParams[0].(string))
		return line
	}
	return 2_000_000
}

type position struct {
	x, y int
}

func parsePosition(data string) position {
	data = strings.Trim(data, ":")
	parts := strings.Split(data, ",")
	x, _ := strconv.Atoi(parts[0][2:])
	y, _ := strconv.Atoi(parts[1][2:])
	return position{x, y}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func less(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func greater(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (p position) manhattanDistance(other position) int {
	return abs(p.x-other.x) + abs(p.y-other.y)
}

type pair struct {
	sensor, beacon position
}

func (p pair) cover(lineY int) intRange {
	sensorBeaconDistance := p.sensor.manhattanDistance(p.beacon)
	sensorGoalDistance := abs(p.sensor.y - lineY)
	if sensorGoalDistance > sensorBeaconDistance {
		return intRange{0, 0}
	}
	offset := sensorBeaconDistance - sensorGoalDistance
	return intRange{p.sensor.x - offset, p.sensor.x + offset + 1}
}

type intRange struct {
	from, to int
}

func (r intRange) contains(p position) bool {
	return r.from <= p.x && r.to > p.x
}

func (r intRange) mergable(other intRange) bool {
	return r.to == other.from || r.from == other.to ||
		(r.from <= other.from && r.to >= other.to) ||
		(other.from <= r.from && other.to >= r.to) ||
		(r.from <= other.from && r.to >= other.from) ||
		(other.from <= r.from && other.to >= r.from)
}

func (r intRange) merge(other intRange) intRange {
	from := less(r.from, other.from)
	to := greater(r.to, other.to)
	return intRange{from, to}
}

func (r intRange) size() int {
	return r.to - r.from
}

func parsePairs(input string) []pair {
	lines := strings.Split(input, "\n")
	pairs := make([]pair, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")
		sensor := parsePosition(parts[2] + parts[3])
		beacon := parsePosition(parts[8] + parts[9])
		pairs[i] = pair{sensor, beacon}
	}
	return pairs
}

func mergeRanges(ranges []intRange) []intRange {
	for i := 0; i < len(ranges); i++ {
		for j := i + 1; j < len(ranges); j++ {
			if ranges[i].mergable(ranges[j]) {
				ranges[i] = ranges[i].merge(ranges[j])
				ranges = append(ranges[:j], ranges[j+1:]...)
				j = i
			}
		}
	}
	return ranges
}

func findFullRange(pairs []pair) intRange {
	from := pairs[0].sensor.x
	to := from
	for _, p := range pairs {
		if p.sensor.x < from {
			from = p.sensor.x
		}
		if p.sensor.x > to {
			to = p.sensor.x
		}
		if p.beacon.x < from {
			from = p.beacon.x
		}
		if p.beacon.x > to {
			to = p.beacon.x
		}
	}
	return intRange{from, to + 1}
}

func getRanges(pairs []pair, lineY int) []intRange {
	ranges := make([]intRange, len(pairs))
	for i, p := range pairs {
		ranges[i] = p.cover(lineY)
	}
	return ranges
}

func cropRanges(ranges []intRange, fullRange intRange) []intRange {
	for i, r := range ranges {
		if r.from < fullRange.from {
			ranges[i].from = fullRange.from
		}
		if r.to > fullRange.to {
			ranges[i].to = fullRange.to
		}
	}
	return ranges
}

func calculateInvalidPositions(ranges []intRange) int {
	sum := 0
	for _, r := range ranges {
		sum += r.size()
	}
	return sum
}

func calculateOccupiedPositions(pairs []pair, ranges []intRange, lineY int) int {
	positions := make(map[position]struct{})
	for _, p := range pairs {
		if p.sensor.y == lineY {
			positions[p.sensor] = struct{}{}
		}
		if p.beacon.y == lineY {
			positions[p.beacon] = struct{}{}
		}
	}
	for p := range positions {
		contains := false
		for _, r := range ranges {
			if r.contains(p) {
				contains = true
				break
			}
		}
		if !contains {
			delete(positions, p)
		}
	}
	return len(positions)
}

func (*Solver) SolvePart2(input string, extraParams ...any) string {
	maxCoord := getMaxCoord(extraParams)
	fullRange := intRange{0, maxCoord + 1}
	pairs := parsePairs(input)
	for i := 0; i < maxCoord + 1; i++ {
		ranges := getRanges(pairs, i)
		ranges = cropRanges(ranges, fullRange)
		ranges = mergeRanges(ranges)
		occupied := calculateInvalidPositions(ranges)
		if occupied != maxCoord + 1 {
			x := findEmpty(ranges, fullRange, i)
			return strconv.Itoa(x * 4_000_000 + i)
		}
	}
	return "not found"
}

func getMaxCoord(extraParams []any) int {
	if len(extraParams) > 0 {
		line, _ := strconv.Atoi(extraParams[0].(string))
		return line
	}
	return 4_000_000
}

func findEmpty(ranges []intRange, fullRange intRange, line int) int {
	for i := fullRange.from; i < fullRange.to; i++ {
		empty := true
		p := position{i, line}
		for _, r := range ranges {
			if r.contains(p) {
				empty = false
				break
			}
		}
		if empty {
			return i
		}
	}
	return -1
}
