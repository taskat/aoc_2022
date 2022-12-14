package day14

import (
	"strconv"
	"strings"
)

type Solver struct {}

func (*Solver) SolvePart1(input string, extraParams ...any) string {
	c := parseCave(input, findSize)
	return strconv.Itoa(pourSand(c, position{500, 0}))
}

type cave struct {
	tiles [][]rune
	minX int
}

func newCave(i, j, minX int) cave {
	c := cave{minX: minX - 500}
	c.tiles = make([][]rune, j)
	for k := range c.tiles {
		c.tiles[k] = make([]rune, i)
		for l := range c.tiles[k] {
			c.tiles[k][l] = '.'
		}
	}
	return c
}

func (c *cave) drawLine(p1, p2 position) {
	if p1.x == p2.x {
		if p1.y > p2.y {
			p1, p2 = p2, p1
		}
		for y := p1.y; y <= p2.y; y++ {
			*c.get(p1.x, y) = '#'
		}
	} else {
		if p1.x > p2.x {
			p1, p2 = p2, p1
		}
		for x := p1.x; x <= p2.x; x++ {
			*c.get(x, p1.y) = '#'
		}
	}
}

func (c *cave) get(x, y int) *rune {
	newX := x - 500 - c.minX
	if newX < 0 || newX >= len(c.tiles[0]) {
		panic("index out of bounds")
	}
	return &c.tiles[y][newX]
}

func (c *cave) addSand(x, y int) (rest bool) {
	defer func() {
		if r := recover(); r != nil {
			rest = false
		}
	}()
	current := position{x, y}
	for {
		if *c.get(current.x, current.y) != '.' {
			return false
		}
		if *c.get(current.x, current.y + 1) == '.' {
			current.y++
			continue
		}
		if *c.get(current.x - 1, current.y + 1) == '.' {
			current.x--
			current.y++
			continue
		}
		if *c.get(current.x + 1, current.y + 1) == '.' {
			current.x++
			current.y++
			continue
		}
		*c.get(current.x, current.y) = 'O'
		return true
	}
}

func (c *cave) String() string {
	lines := make([]string, len(c.tiles))
	for i, line := range c.tiles {
		chars := make([]string, len(line))
		for j, char := range line {
			chars[j] = string(char)
		}
		lines[i] = strings.Join(chars, "")
	}
	return strings.Join(lines, "\n")
}

type position struct {
	x, y int
}

func parsePosition(input string) position {
	parts := strings.Split(input, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])	
	return position{x, y}
}

func parseCave(input string, size func(positions []position) (position, position)) cave {
	positions := parsePositions(input)
	min, max := size(positions)
	c := newCave(max.x - min.x + 1, max.y - min.y + 1, min.x)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		corners := make([]position, len(parts))
		for i, part := range parts {
			corners[i] = parsePosition(part)
		}
		for i := 0; i < len(corners) - 1; i++ {
			c.drawLine(corners[i], corners[i + 1])
		}
	}
	return c
}

func parsePositions(input string) []position {
	lines := strings.Split(input, "\n")
	positions := make([]position, 0, 2 * len(lines))
	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		for _, part := range parts {
			positions = append(positions, parsePosition(part))
		}	
	}
	return positions
}

func findSize(positions []position) (position, position) {
	maxPos := position{500, 0}
	minPos := position{500, 0}
	for _, pos := range positions {
		if pos.x > maxPos.x {
			maxPos.x = pos.x
		}
		if pos.y > maxPos.y {
			maxPos.y = pos.y
		}
		if pos.x < minPos.x {
			minPos.x = pos.x
		}
		if pos.y < minPos.y {
			minPos.y = pos.y
		}
	}
	return minPos, maxPos
}

func pourSand(c cave, pos position) int {
	count := 0
	for c.addSand(pos.x, pos.y) {
		count++
	}
	return count
}

func (*Solver) SolvePart2(input string, extraParams ...any) string {
	c := parseCave(input, findEndlessSize)
	positions := parsePositions(input)
	min, max := findEndlessSize(positions)
	start := position{min.x, max.y}
	end := position{max.x, max.y}
	c.drawLine(start, end)
	return strconv.Itoa(pourSand(c, position{500, 0}))
}

func findEndlessSize(positions []position) (position, position) {
	maxPos := position{500, 0}
	minPos := position{500, 0}
	for _, pos := range positions {
		if pos.x > maxPos.x {
			maxPos.x = pos.x
		}
		if pos.y > maxPos.y {
			maxPos.y = pos.y
		}
		if pos.x < minPos.x {
			minPos.x = pos.x
		}
		if pos.y < minPos.y {
			minPos.y = pos.y
		}
	}
	extra := maxPos.y
	minPos.x -= extra
	maxPos.x += extra
	maxPos.y += 2
	return minPos, maxPos
}
