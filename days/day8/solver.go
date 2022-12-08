package day8

import (
	"strconv"
	"strings"
)

type Solver struct {}

func (*Solver) SolvePart1(input string, extraParams ...any) string {
	g := parseGrid(input)
	return strconv.Itoa(g.visibleFromOutside())
}

type tree int

type grid [][]tree

func parseGrid(input string) grid {
	lines := strings.Split(input, "\n")
	g := make(grid, len(lines))
	for i, line := range lines {
		g[i] = make([]tree, len(line))
		for j, c := range line {
			g[i][j] = tree(c)
		}
	}
	return g
}

func (g grid) visibleFromLeft(i, j int) bool {
	for c := 0; c < j; c++ {
		if g[i][c] >= g[i][j] {
			return false
		}
	}
	return true
}

func (g grid) visibleFromRight(i, j int) bool {
	for c := j + 1; c < len(g[i]); c++ {
		if g[i][c] >= g[i][j] {
			return false
		}
	}
	return true
}

func (g grid) visibleFromTop(i, j int) bool {
	for r := 0; r < i; r++ {
		if g[r][j] >= g[i][j] {
			return false
		}
	}
	return true
}

func (g grid) visibleFromBottom(i, j int) bool {
	for r := i + 1; r < len(g); r++ {
		if g[r][j] >= g[i][j] {
			return false
		}
	}
	return true
}

func (g grid) visible(i, j int) bool {
	visibleFromSides := []func(int, int) bool{
		g.visibleFromLeft,
		g.visibleFromRight,
		g.visibleFromTop,
		g.visibleFromBottom,
	}
	for _, visibleFromSide := range visibleFromSides {
		if visibleFromSide(i, j) {
			return true
		}
	}
	return false
}

func (g grid) visibleFromOutside() int {
	visible := 0
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[i]); j++ {
			if g.visible(i, j) {
				visible++
			}
		}
	}
	return visible
}

func (*Solver) SolvePart2(input string, extraParams ...any) string {
	g := parseGrid(input)
	return strconv.Itoa(g.maxScore())
}

func (g grid) scoreFromLeft(i, j int) int {
	score := 0
	for c := j - 1; c >= 0; c-- {
		score++
		if g[i][c] >= g[i][j] {
			break
		}
	}
	return score
}

func (g grid) scoreFromRight(i, j int) int {
	score := 0
	for c := j + 1; c < len(g[i]); c++ {
		score++
		if g[i][c] >= g[i][j] {
			break
		}
	}
	return score
}

func (g grid) scoreFromTop(i, j int) int {
	score := 0
	for r := i - 1; r >= 0; r-- {
		score++
		if g[r][j] >= g[i][j] {
			break
		}
	}
	return score
}

func (g grid) scoreFromBottom(i, j int) int {
	score := 0
	for r := i + 1; r < len(g); r++ {
		score++
		if g[r][j] >= g[i][j] {
			break
		}
	}
	return score
}

func (g grid) score(i, j int) int {
	scores := []func(int, int) int{
		g.scoreFromLeft,
		g.scoreFromRight,
		g.scoreFromTop,
		g.scoreFromBottom,
	}
	score := 1
	for _, s := range scores {
		score *= s(i, j)
	}
	return score
}

func (g grid) getScores() [][]int {
	scores := make([][]int, len(g))
	for i := 0; i < len(g); i++ {
		scores[i] = make([]int, len(g[i]))
		for j := 0; j < len(g[i]); j++ {
			scores[i][j] = g.score(i, j)
		}
	}
	return scores
}

func (g grid) maxScore() int {
	maxScore := 0
	for _, row := range g.getScores() {
		for _, score := range row {
			if score > maxScore {
				maxScore = score
			}
		}
	}
	return maxScore
}
