package day12

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/beefsack/go-astar"
)

type Solver struct{}

func (*Solver) SolvePart1(input string, extraParams ...any) string {
	g := parseGrid(input)
	start := &g.squares[g.start.y][g.start.x]
	goal := &g.squares[g.goal.y][g.goal.x]
	_, cost, found := astar.Path(start, goal)
	fmt.Println(found)
	return strconv.Itoa(int(cost))
}

type position struct {
	x int
	y int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (p position) manhattanDistance(to position) int {
	return abs(p.x-to.x) + abs(p.y-to.y)
}

type grid struct {
	squares [][]square
	goal    position
	start   position
}

func parseGrid(input string) *grid {
	lines := strings.Split(input, "\n")
	squares := make([][]square, len(lines))
	var g grid
	for i, line := range lines {
		squares[i] = make([]square, len(line))
		for j, char := range line {
			squares[i][j] = *newSquare(char, position{x: j, y: i}, &g)
			if char == 'S' {
				g.start = position{x: j, y: i}
				squares[i][j].height = 'a'
			}
			if char == 'E' {
				g.goal = position{x: j, y: i}
				squares[i][j].height = 'z'
			}
		}
	}
	g.squares = squares
	return &g
}

type square struct {
	height rune
	pos    position
	g      *grid
}

func newSquare(height rune, pos position, g *grid) *square {
	return &square{height: height, pos: pos, g: g}
}

func (s *square) PathNeighbors() []astar.Pather {
	neighborCandidates := []square{}
	if s.pos.x > 0 {
		neighborCandidates = append(neighborCandidates, s.g.squares[s.pos.y][s.pos.x-1])
	}
	if s.pos.x < len(s.g.squares[s.pos.y])-1 {
		neighborCandidates = append(neighborCandidates, s.g.squares[s.pos.y][s.pos.x+1])
	}
	if s.pos.y > 0 {
		neighborCandidates = append(neighborCandidates, s.g.squares[s.pos.y-1][s.pos.x])
	}
	if s.pos.y < len(s.g.squares)-1 {
		neighborCandidates = append(neighborCandidates, s.g.squares[s.pos.y+1][s.pos.x])
	}
	neighbors := []astar.Pather{}
	for _, neighbor := range neighborCandidates {
		if neighbor.height-s.height <= 1 {
			neighbors = append(neighbors, &s.g.squares[neighbor.pos.y][neighbor.pos.x])
		}
	}
	return neighbors
}

func (s *square) PathNeighborCost(to astar.Pather) float64 {
	return 1
}

func (s *square) PathEstimatedCost(to astar.Pather) float64 {
	return float64(s.pos.manhattanDistance(to.(*square).pos))
}

func (*Solver) SolvePart2(input string, extraParams ...any) string {
	g := parseGrid(input)
	startingPositions := findStartingPositions(*g)
	steps := findSteps(*g, startingPositions)
	return strconv.Itoa(findMinSteps(steps))
}

func findStartingPositions(g grid) []*square {
	startingPositions := []*square{}
	for i, row := range g.squares {
		for j, square := range row {
			if square.height == 'a' {
				startingPositions = append(startingPositions, &g.squares[i][j])
			}
		}
	}
	return startingPositions
}

func findSteps(g grid, startingPositions []*square) []int {
	steps := []int{}
	goal := &g.squares[g.goal.y][g.goal.x]
	for _, startingPosition := range startingPositions {
		_, cost, found := astar.Path(startingPosition, goal)
		if found {
			steps = append(steps, int(cost))
		}
	}
	return steps
}

func findMinSteps(steps []int) int {
	min := steps[0]
	for _, step := range steps {
		if step < min {
			min = step
		}
	}
	return min
}
