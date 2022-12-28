package day17

import (
	"fmt"
	"strconv"
	"strings"
)

const WIDTH = 7

type Solver struct{}

func (*Solver) SolvePart1(input string, extraParams ...any) string {
	directions := parseDirections(input)
	return strconv.Itoa(simulate(directions, 2022))
}

func parseDirections(input string) []direction {
	directions := make([]direction, 0, len(input))
	for _, c := range input {
		if c == '>' {
			directions = append(directions, right)
		} else if c == '<' {
			directions = append(directions, left)
		} else {
			panic(fmt.Sprintf("Unknown direction %c", c))
		}
	}
	return directions
}

type direction bool

var left direction = false
var right direction = true

func (d direction) String() string {
	if d == left {
		return "left"
	}
	return "right"
}

type row struct {
	data   [WIDTH]bool
	ranges []Range
}

func newRow(height int) row {
	return row{data: [WIDTH]bool{}, ranges: []Range{*newRange(0, WIDTH, height)}}
}

func (r *row) decreaseStartingHeight(n int) {
	for i := 0; i < len(r.ranges); i++ {
		r.ranges[i].decreaseStartingHeight(n)
	}
}

func (r *row) isEmpty() bool {
	for i := 0; i < WIDTH; i++ {
		if r.data[i] {
			return false
		}
	}
	return true
}

func (r *row) minHeight() int {
	if len(r.ranges) == 0 {
		return -1
	}
	min := r.ranges[0].startingHeight
	for _, interval := range r.ranges {
		if interval.startingHeight < min {
			min = interval.startingHeight
		}
	}
	return min
}

func (r *row) update(currHeight int) {
	newRanges := make([]Range, 0, 4)
	start := -1
	for i := 0; i < WIDTH; i++ {
		if !r.data[i] && start == -1 {
			start = i
		}
		if r.data[i] && start != -1 {
			newRanges = append(newRanges, *newRange(start, i, currHeight))
			start = -1
		}
	}
	if start != -1 {
		newRanges = append(newRanges, *newRange(start, WIDTH, currHeight))
	}
	r.ranges = newRanges
}

func (r *row) updateStartingHeight(below *row, current int) bool {
	updated := false
	for i, interval := range r.ranges {
		heights := make([]int, 0, 4)
		for _, belowInterval := range below.ranges {
			if interval.hasIntersection(&belowInterval) {
				heights = append(heights, belowInterval.startingHeight)
			}
		}
		if len(heights) == 0 {
			r.ranges[i].startingHeight = current
			continue
		}
		min := heights[0]
		for _, h := range heights {
			if h < min {
				min = h
			}
		}
		r.ranges[i].startingHeight = min
	}
	return updated
}

func (r *row) String() string {
	cave := make([]byte, 0, WIDTH)
	for i := 0; i < WIDTH; i++ {
		if r.data[i] {
			cave = append(cave, '#')
		} else {
			cave = append(cave, '.')
		}
	}
	ranges := make([]string, 0, len(r.ranges))
	for i := 0; i < len(r.ranges); i++ {
		ranges = append(ranges, r.ranges[i].String())
	}
	return string(cave) + strings.Join(ranges, ", ")
}

type chamber struct {
	rows   []row
	height int
}

func newChamber() *chamber {
	return &chamber{rows: []row{}, height: 0}
}

func (c *chamber) addShape(s shape) {
	baseHeight := len(c.rows)
	for i := 0; i < 3+len(s.rocks); i++ {
		c.rows = append(c.rows, newRow(baseHeight))
	}
}

func (c *chamber) decreaseStartingHeight(n int) {
	for _, r := range c.rows {
		r.decreaseStartingHeight(n)
	}
	c.height += n
}

func (c *chamber) fixShape(s shape) {
	for i := 0; i < len(s.rocks); i++ {
		for j := 0; j < WIDTH; j++ {
			if s.rocks[i][j] {
				c.rows[i+s.height].data[j] = true
			}
		}
		c.rows[i+s.height].update(i + s.height)
		if i+s.height > 0 {
			c.rows[i+s.height].updateStartingHeight(&c.rows[i+s.height-1], i+s.height)
		}
	}
	c.removeEmptyLines()
	for i := s.height + 1; i < len(c.rows); i++ {
		c.rows[i].updateStartingHeight(&c.rows[i-1], i)
	}
	c.removeUnnecessaryLines()
}

func (c *chamber) isValid(s shape) bool {
	if s.height < 0 {
		return false
	}
	for i := 0; i < len(s.rocks); i++ {
		for j := 0; j < WIDTH; j++ {
			if s.rocks[i][j] && c.rows[i+s.height].data[j] {
				return false
			}
		}
	}
	return true
}

func (c *chamber) currentMaxHeight() int {
	for i := len(c.rows) - 1; i >= 0; i-- {
		if !c.rows[i].isEmpty() {
			return i + 1
		}
	}
	return 0
}

func (c *chamber) maxHeight() int {
	return c.height + c.currentMaxHeight()
}

func (c *chamber) minNecessaryHeight() int {
	return c.rows[len(c.rows)-1].minHeight()
}

func (c *chamber) removeEmptyLines() {
	firstEmpty := -1
	for i := len(c.rows) - 1; i >= 0; i-- {
		if c.rows[i].isEmpty() {
			firstEmpty = i
		} else {
			break
		}
	}
	if firstEmpty != -1 {
		c.rows = c.rows[:firstEmpty]
	}
}

func (c *chamber) removeUnnecessaryLines() {
	minHeight := c.minNecessaryHeight()
	if minHeight > 0 {
		c.rows = c.rows[minHeight:]
		c.decreaseStartingHeight(minHeight)
	}
}

func (c *chamber) String() string {
	lines := make([]string, len(c.rows))
	for i := 0; i < len(c.rows); i++ {
		lines[len(c.rows)-1-i] = c.rows[i].String()
		lines[len(c.rows)-1-i] += fmt.Sprintf(" %d", c.maxHeight())
	}
	return strings.Join(lines, "\n")
}

type shape struct {
	rocks  [][WIDTH]bool
	height int
}

func newHorizontalLine(height int) shape {
	return shape{
		rocks: [][WIDTH]bool{
			{false, false, true, true, true, true, false},
		},
		height: height + 3,
	}
}

func newCross(height int) shape {
	return shape{
		rocks: [][WIDTH]bool{
			{false, false, false, true, false, false, false},
			{false, false, true, true, true, false, false},
			{false, false, false, true, false, false, false},
		},
		height: height + 3,
	}
}

func newAngle(height int) shape {
	return shape{
		rocks: [][WIDTH]bool{
			{false, false, true, true, true, false, false},
			{false, false, false, false, true, false, false},
			{false, false, false, false, true, false, false},
		},
		height: height + 3,
	}
}

func newVerticalLine(height int) shape {
	return shape{
		rocks: [][WIDTH]bool{
			{false, false, true, false, false, false, false},
			{false, false, true, false, false, false, false},
			{false, false, true, false, false, false, false},
			{false, false, true, false, false, false, false},
		},
		height: height + 3,
	}
}

func newSquare(height int) shape {
	return shape{
		rocks: [][WIDTH]bool{
			{false, false, true, true, false, false, false},
			{false, false, true, true, false, false, false},
		},
		height: height + 3,
	}
}

func (s *shape) fall() {
	s.height--
}

func (s *shape) revertfall() {
	s.height++
}

func (s *shape) move(d direction, c *chamber) {
	if d == left {
		for i := 0; i < len(s.rocks); i++ {
			if s.rocks[i][0] {
				return
			}
			for j := 1; j < WIDTH; j++ {
				if s.rocks[i][j] && c.rows[s.height+i].data[j-1] {
					return
				}
			}
		}

		for i := 0; i < len(s.rocks); i++ {
			for j := 0; j < 6; j++ {
				s.rocks[i][j] = s.rocks[i][j+1]
			}
			s.rocks[i][6] = false
		}
	} else {
		for i := 0; i < len(s.rocks); i++ {
			if s.rocks[i][6] {
				return
			}
			for j := 5; j >= 0; j-- {
				if s.rocks[i][j] && c.rows[s.height+i].data[j+1] {
					return
				}
			}
		}
		for i := 0; i < len(s.rocks); i++ {
			for j := 6; j > 0; j-- {
				s.rocks[i][j] = s.rocks[i][j-1]
			}
			s.rocks[i][0] = false
		}
	}
}

func (s *shape) String() string {
	lines := make([]string, 0, len(s.rocks))
	for i := 0; i < len(s.rocks); i++ {
		line := make([]rune, 0, WIDTH)
		for j := 0; j < WIDTH; j++ {
			if s.rocks[i][j] {
				line = append(line, '#')
			} else {
				line = append(line, '.')
			}
		}
		lines = append(lines, string(line))
	}
	return strings.Join(lines, "\n")
}

type rockGenerator struct {
	counter      int
	constructors []func(int) shape
}

func newRockGenerator() *rockGenerator {
	constructors := []func(int) shape{
		newHorizontalLine,
		newCross,
		newAngle,
		newVerticalLine,
		newSquare,
	}
	return &rockGenerator{0, constructors}
}

func (r *rockGenerator) next(height int) shape {
	constructor := r.constructors[r.counter%len(r.constructors)]
	r.counter++
	return constructor(height)
}

func simulate(directions []direction, maxRocks int) int {
	chamber := newChamber()
	rockGenerator := newRockGenerator()
	directionCounter := 0
	for i := 0; i < maxRocks; i++ {
		s := rockGenerator.next(chamber.currentMaxHeight())
		chamber.addShape(s)
		for chamber.isValid(s) {
			d := directions[directionCounter]
			s.move(d, chamber)
			directionCounter = (directionCounter + 1) % len(directions)
			s.fall()
		}
		s.revertfall()
		chamber.fixShape(s)
	}
	return chamber.maxHeight()
}

type Range struct {
	from           int
	to             int
	startingHeight int
}

func newRange(from, to, startingHeight int) *Range {
	return &Range{from, to, startingHeight}
}

func (r *Range) decreaseStartingHeight(n int) {
	r.startingHeight -= n
	if r.startingHeight < 0 {
		r.startingHeight = 0
	}
}

func (r *Range) hasIntersection(other *Range) bool {
	return r.from < other.to && other.from < r.to
}

func (r *Range) String() string {
	return fmt.Sprintf("from %d to %d (start: %d)", r.from, r.to, r.startingHeight)
}

func (*Solver) SolvePart2(input string, extraParams ...any) string {
	directions := parseDirections(input)
	return strconv.Itoa(simulateAndCollect(directions, 1_000_000_000_000))
}

type state struct {
	directionIdx int
	rockIdx      int
	rocks        [20]uint8
}

func newState(c *chamber, directionIdx, rockIdx int) state {
	rocks := [20]uint8{}
	for i := 0; i < len(c.rows) && i < 20; i++ {
		var value uint8 = 0
		for j := 0; j < WIDTH; j++ {
			value <<= 1
			if c.rows[i].data[j] {
				value++
			}
		}
		rocks[i] = value
	}
	return state{directionIdx, rockIdx, rocks}
}

func simulateAndCollect(directions []direction, maxRocks int) int {
	chamber := newChamber()
	rockGenerator := newRockGenerator()
	directionCounter := 0
	seen := make(map[state]struct{height, i int})
	for i := 0; i < maxRocks; i++ {
		s := rockGenerator.next(chamber.currentMaxHeight())
		chamber.addShape(s)
		for chamber.isValid(s) {
			d := directions[directionCounter]
			s.move(d, chamber)
			directionCounter = (directionCounter + 1) % len(directions)
			s.fall()
		}
		s.revertfall()
		chamber.fixShape(s)
		newState := newState(chamber, directionCounter, i % 5)
		if data, ok := seen[newState]; ok {
			period := i - data.i
			heightPeriod := chamber.maxHeight() - data.height
			numberOfPeriods := (maxRocks - i) / period
			chamber.height += heightPeriod * numberOfPeriods
			i += period * numberOfPeriods
		} else {
			seen[newState] = struct{height, i int}{chamber.maxHeight(), i}
		}
		if i % 100_000 == 0 {
			fmt.Println(i, chamber.maxHeight(), len(seen))
		}
	}
	return chamber.maxHeight()
}
