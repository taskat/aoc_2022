package day9

import (
	"strconv"
	"strings"
)

type Solver struct{}

func (*Solver) SolvePart1(input string, extraParams ...any) string {
	cmds := parseCommands(input)
	rope := newRope(2)
	return strconv.Itoa(simulate(cmds, rope))
}

type position struct {
	x, y int
}

type direction int

const (
	up direction = iota
	right
	down
	left
)

func (p *position) distance(other position) position {
	return position{
		x: other.x - p.x,
		y: other.y - p.y,
	}
}

func (p *position) move(d direction) {
	switch d {
	case up:
		p.y++
	case right:
		p.x++
	case down:
		p.y--
	case left:
		p.x--
	}
}

type rope struct {
	knots []position
}

func newRope(knots int) rope {
	return rope{
		knots: make([]position, knots),
	}
}

func (r *rope) move(d direction) {
	r.knots[0].move(d)
	for i := 1; i < len(r.knots); i++ {
		distance := r.knots[i].distance(r.knots[i - 1])
		xFar := false
		switch distance.x {
		case -2:
			xFar = true
			r.knots[i].x--
		case 2:
			xFar = true
			r.knots[i].x++
		}
		yFar := false
		switch distance.y {
		case -2:
			yFar = true
			r.knots[i].y--
		case 2:
			yFar = true
			r.knots[i].y++
		}
		if xFar && !yFar {
			if distance.y == -1 {
				r.knots[i].y--
			}
			if distance.y == 1 {
				r.knots[i].y++
			}
		}
		if yFar && !xFar {
			if distance.x == -1 {
				r.knots[i].x--
			}
			if distance.x == 1 {
				r.knots[i].x++
			}
		}
	}
}

func (r *rope) tail() position {
	return r.knots[len(r.knots) - 1]
}

type command struct {
	d direction
	n int
}

func newCommand(line string) command {
	parts := strings.Split(line, " ")
	var d direction
	switch parts[0] {
	case "U":
		d = up
	case "R":
		d = right
	case "D":
		d = down
	case "L":
		d = left
	}
	n, _ := strconv.Atoi(parts[1])
	return command{d, n}
}

func parseCommands(input string) []command {
	lines := strings.Split(input, "\n")
	commands := make([]command, len(lines))
	for i, line := range lines {
		commands[i] = newCommand(line)
	}
	return commands
}

func simulate(commands []command, r rope) int {
	visited := make(map[position]struct{})
	visited[r.tail()] = struct{}{}
	for _, command := range commands {
		for i := 0; i < command.n; i++ {
			r.move(command.d)
			visited[r.tail()] = struct{}{}
		}
	}
	return len(visited)
}

func (*Solver) SolvePart2(input string, extraParams ...any) string {
	cmds := parseCommands(input)
	rope := newRope(10)
	return strconv.Itoa(simulate(cmds, rope))
}
