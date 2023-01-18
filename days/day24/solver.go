package day24

import (
	"aoc_2022/days/day24/astar"
	"fmt"
	"strings"
)

type Solver struct{}

func (*Solver) SolvePart1(input string, extraParams ...any) string {
	v := parseValley(input)
	m := newMaze(v)
	startPos := pos{row: -1, col: 0}
	goalPos := pos{row: len(v), col: len(v[0]) - 1}
	isGoal := func(p astar.Pather) bool {
		s, ok := p.(*state)
		if !ok {
			panic("not a state")
		}
		return s.pos == goalPos
	}
	best, found := astar.Path(&state{pos: startPos, m: m, time: 0}, isGoal)
	if !found {
		panic("No solution found")
	}
	return fmt.Sprintf("%d", best)
}

type blizzard int

const (
	empty blizzard = iota
	up
	right
	down
	left
)

func (t blizzard) String() string {
	switch t {
	case empty:
		return "."
	case up:
		return "^"
	case right:
		return ">"
	case down:
		return "v"
	case left:
		return "<"
	default:
		panic("Unknown tile")
	}
}

type tile []blizzard

func newEmptyTile() *tile {
	return &tile{empty}
}

func newTile(b blizzard) tile {
	return tile{b}
}

func parseTile(c rune) tile {
	switch c {
	case '.':
		return newTile(empty)
	case '^':
		return newTile(up)
	case '>':
		return newTile(right)
	case 'v':
		return newTile(down)
	case '<':
		return newTile(left)
	default:
		panic("Unknown tile")
	}
}

func (t *tile) add(b blizzard) {
	if len(*t) == 1 && (*t)[0] == empty {
		*t = tile{b}
	} else {
		*t = append(*t, b)
	}
}

func (t *tile) isEmpty() bool {
	return len(*t) == 1 && (*t)[0] == empty
}

func (t tile) String() string {
	if len(t) == 1 {
		return t[0].String()
	}
	return fmt.Sprintf("%d", len(t))
}

type valley [][]tile

func parseValley(input string) valley {
	lines := strings.Split(input, "\n")
	v := make([][]tile, len(lines)-2)
	for i := 1; i < len(lines)-1; i++ {
		v[i-1] = make([]tile, len(lines[i])-2)
		for j := 1; j < len(lines[i])-1; j++ {
			v[i-1][j-1] = parseTile(rune(lines[i][j]))
		}
	}
	return v
}

func (v valley) get(r, c int) *tile {
	if r == -1 && c == 0 {
		return newEmptyTile()
	}
	if r == len(v) && c == len(v[0])-1 {
		return newEmptyTile()
	}
	return &v[r][c]
}

func (v valley) next() valley {
	newValley := make([][]tile, len(v))
	for i := 0; i < len(v); i++ {
		newValley[i] = make([]tile, len(v[i]))
		for j := 0; j < len(v[i]); j++ {
			newValley[i][j] = *newEmptyTile()
		}
	}
	maxRow := len(v) - 1
	maxCol := len(v[0]) - 1
	for i := 0; i < len(v); i++ {
		for j := 0; j < len(v[i]); j++ {
			for k := 0; k < len(v[i][j]); k++ {
				switch v[i][j][k] {
				case empty:
					continue
				case up:
					if i == 0 {
						newValley[maxRow][j].add(up)
					} else {
						newValley[i-1][j].add(up)
					}
				case right:
					if j == maxCol {
						newValley[i][0].add(right)
					} else {
						newValley[i][j+1].add(right)
					}
				case down:
					if i == maxRow {
						newValley[0][j].add(down)
					} else {
						newValley[i+1][j].add(down)
					}
				case left:
					if j == 0 {
						newValley[i][maxCol].add(left)
					} else {
						newValley[i][j-1].add(left)
					}
				}
			}
		}
	}
	for i := 0; i < len(newValley); i++ {
		for j := 0; j < len(newValley[i]); j++ {
			if len(newValley[i][j]) == 0 {
				newValley[i][j] = append(newValley[i][j], empty)
			}
		}
	}
	return newValley
}

func (v valley) String() string {
	lines := make([]string, len(v))
	for i := 0; i < len(v); i++ {
		newLine := make([]string, len(v[i]))
		for j := 0; j < len(v[i]); j++ {
			newLine[j] = v[i][j].String()
		}
		lines[i] = strings.Join(newLine, "")
	}
	return strings.Join(lines, "\n")
}

type maze struct {
	valleys []valley
}

func newMaze(v valley) *maze {
	rows := len(v)
	cols := len(v[0])
	valleys := make([]valley, rows * cols)
	for i := 0; i < rows * cols; i++ {
		valleys[i] = v
		v = v.next()
	}
	return &maze{valleys}
}

func (m *maze) String() string {
	lines := make([]string, len(m.valleys))
	for i := 0; i < len(m.valleys); i++ {
		lines[i] = fmt.Sprintf("%d\n%s", i, m.valleys[i].String())
	}
	return strings.Join(lines, "\n\n")
}

type pos struct {
	row int
	col int
}

func (p pos) manhattanDistance(other pos) int {
	return abs(p.row-other.row) + abs(p.col-other.col)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

type state struct {
	pos pos
	m *maze
	time int
}

func (s *state) GetHash() uint64 {
	return ((uint64(s.time) * uint64(len(s.m.valleys[0]) + 1) + uint64(s.pos.row)) * uint64(len(s.m.valleys[0][0])) + 1) + uint64(s.pos.col)
}

func (s *state) PathEstimatedCost(to astar.Pather) float64 {
	return float64(s.pos.manhattanDistance(to.(*state).pos))
}

func (s *state) PathNeighborCost(to astar.Pather) float64 {
	return 1
}

func (s *state) PathNeighbors() []astar.Pather {
	neighbors := make([]astar.Pather, 0, 5)
	maxRow := len(s.m.valleys[0]) - 1
	maxCol := len(s.m.valleys[0][0]) - 1
	row := s.pos.row
	col := s.pos.col
	next := (s.time + 1) % len(s.m.valleys)
	if row == -1 && col == 0 {
		neighbors = append(neighbors, &state{pos{-1, 0}, s.m, s.time +1})
		if s.m.valleys[next].get(0, 0).isEmpty() {
			neighbors = append(neighbors, &state{pos{0, 0}, s.m, s.time +1})
		}
		return neighbors
	}
	if row == maxRow + 1 && col == maxCol {
		neighbors = append(neighbors, &state{pos{maxRow + 1, maxCol}, s.m, s.time +1})
		if s.m.valleys[next].get(maxRow, maxCol).isEmpty() {
			neighbors = append(neighbors, &state{pos{maxRow, maxCol}, s.m, s.time +1})
		}
		return neighbors
	}
	if row > 0 && s.m.valleys[next].get(row-1, col).isEmpty() {
		neighbors = append(neighbors, &state{pos{row-1, col}, s.m, s.time +1})
	}
	if row < maxRow && s.m.valleys[next].get(row+1, col).isEmpty() {
		neighbors = append(neighbors, &state{pos{row+1, col}, s.m, s.time +1})
	}
	if col > 0 && s.m.valleys[next].get(row, col-1).isEmpty() {
		neighbors = append(neighbors, &state{pos{row, col-1}, s.m, s.time +1})
	}
	if col < maxCol && s.m.valleys[next].get(row, col+1).isEmpty() {
		neighbors = append(neighbors, &state{pos{row, col+1}, s.m, s.time +1})
	}
	if row == 0 && col == 0 {
		neighbors = append(neighbors, &state{pos{row-1, col}, s.m, s.time +1})
	}
	if row == maxRow && col == maxCol {
		neighbors = append(neighbors, &state{pos{row+1, col}, s.m, s.time +1})
	}
	if s.m.valleys[next].get(row, col).isEmpty() {
		neighbors = append(neighbors, &state{s.pos, s.m, s.time +1})
	}
	return neighbors
}

func (s *state) String() string {
	return fmt.Sprintf("(row: %d, col: %d, time: %d)", s.pos.row, s.pos.col, s.time)
}

func (*Solver) SolvePart2(input string, extraParams ...any) string {
	v := parseValley(input)
	m := newMaze(v)
	startPos := pos{row: -1, col: 0}
	goalPos := pos{row: len(v), col: len(v[0]) - 1}
	isGoal := func(p astar.Pather) bool {
		s, ok := p.(*state)
		if !ok {
			panic("not a state")
		}
		return s.pos == goalPos
	}
	isStart := func(p astar.Pather) bool {
		s, ok := p.(*state)
		if !ok {
			panic("not a state")
		}
		return s.pos == startPos
	}
	var from, to *state
	time := 0
	isTo := isGoal
	isFrom := isStart
	from = &state{pos: startPos, m: m, time: 0}
	to = &state{pos: goalPos, m: m, time: 0}
	for i := 0; i < 3; i++ {
		best, found := astar.Path(from, isTo)
		if !found {
			panic("No solution found")
		}
		time += best
		to.time = time
		from.time = time
		from, to = to, from
		isFrom, isTo = isTo, isFrom
	}
	return fmt.Sprintf("%d", time)
}
