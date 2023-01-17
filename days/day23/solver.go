package day23

import (
	"fmt"
	"strings"
)

type Solver struct {}

func (*Solver) SolvePart1(input string, extraParams ...any) string {
	elves := parseElves(input)
	elves = doRounds(elves, 10)
	return fmt.Sprintf("%d", calculateEmptyGrounds(elves))
}

type position struct {
	row int
	col int
}

func (p position) adjacent(d direction) position {
	switch d {
	case N:
		return position{p.row - 1, p.col}
	case NE:
		return position{p.row - 1, p.col + 1}
	case E:
		return position{p.row, p.col + 1}
	case SE:
		return position{p.row + 1, p.col + 1}
	case S:
		return position{p.row + 1, p.col}
	case SW:
		return position{p.row + 1, p.col - 1}
	case W:
		return position{p.row, p.col - 1}
	case NW:
		return position{p.row - 1, p.col - 1}
	default:
		panic("invalid direction")
	}
} 

func parseElves(input string) map[position]int {
	elves := make(map[position]int)
	lines := strings.Split(input, "\n")
	elfCounter := 0
	for r, line := range lines {
		for c, char := range line {
			if char == '#' {
				elves[position{r, c}] = elfCounter
				elfCounter++
			}
		}
	}
	return elves
}

type direction int

const (
	N direction = iota
	NE
	E
	SE
	S
	SW
	W
	NW
)

func directionsToCheck(d direction) []direction {
	switch d {
	case N:
		return []direction{NW, N, NE}
	case E:
		return []direction{NE, E, SE}
	case S:
		return []direction{SE, S, SW}
	case W:
		return []direction{SW, W, NW}
	default:
		panic("invalid direction")
	}
}

func canMove(d direction, from position, elves map[position]int) bool {
	for _, dirToCheck := range directionsToCheck(d) {
		if _, ok := elves[from.adjacent(dirToCheck)]; ok {
			return false
		}
	}
	return true
}

func propose(elves map[position]int, directions []direction) (map[int]position, bool) {
	proposals := make(map[int]position, len(elves))
	everyoneStayed := true
	for elfPos, elf := range elves {
		if isAlone(elfPos, elves) {
			proposals[elf] = elfPos
			continue
		}
		everyoneStayed = false
		for _, dir := range directions {
			if canMove(dir, elfPos, elves) {
				newPos := elfPos.adjacent(dir)
				proposals[elf] = newPos
				break
			}
		}
		if _, ok := proposals[elf]; !ok {
			proposals[elf] = elfPos
		}
	}
	if len(proposals) != len(elves) {
		panic("not all elves have a proposal")
	}
	return proposals, everyoneStayed
}

func isAlone(p position, elves map[position]int) bool {
	for _, dir := range []direction{N, NE, E, SE, S, SW, W, NW} {
		if _, ok := elves[p.adjacent(dir)]; ok {
			return false
		}
	}
	return true
}

func move(elves map[position]int, proposals map[int]position) map[position]int {
	revProposals := make(map[position]int, len(proposals))
	remainingElves := make([]int, 0)
	for elf, pos := range proposals {
		otherElf, occupied := revProposals[pos]
		if !occupied {
			revProposals[pos] = elf
			continue
		}
		remainingElves = append(remainingElves, elf, otherElf)
	}
	for _, elf := range remainingElves {
		for pos, e := range elves {
			if e == elf {
				proposals[elf] = pos
				break
			}
		}
	}
	newPositions := make(map[position]int, len(proposals))
	for elf, pos := range proposals {
		newPositions[pos] = elf
	}
	return newPositions
}

func doRounds(elves map[position]int, rounds int) map[position]int {
	directions := []direction{N, S, W, E}
	for i := 0; i < rounds; i++ {
		proposals, _ := propose(elves, directions)
		elves = move(elves, proposals)
		directions = append(directions[1:], directions[0])
	}
	return elves
}

func calculateEmptyGrounds(elves map[position]int) int {
	rows := getRowCount(elves)
	cols := getColCount(elves)
	allTiles := rows * cols
	occupiedTiles := len(elves)
	return allTiles - occupiedTiles
}

func getRowCount(elves map[position]int) int {
	minRow := 0
	maxRow := 0
	for pos := range elves {
		minRow = pos.row
		maxRow = pos.row
		break
	}
	for pos := range elves {

		minRow = min(minRow, pos.row)
		maxRow = max(maxRow, pos.row)
	}
	return abs(maxRow - minRow) + 1
}

func getColCount(elves map[position]int) int {
	minCol := 0
	maxCol := 0
	for pos := range elves {
		minCol = pos.col
		maxCol = pos.col
		break
	}
	for pos := range elves {
		minCol = min(minCol, pos.col)
		maxCol = max(maxCol, pos.col)
	}
	return abs(maxCol - minCol) + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func (*Solver) SolvePart2(input string, extraParams ...any) string {
	elves := parseElves(input)
	rounds := findRound(elves)
	return fmt.Sprintf("%d", rounds)
}

func findRound(elves map[position]int) int {
	directions := []direction{N, S, W, E}
	for rounds := 1; ; rounds++ {
		proposals, everyoneStayed := propose(elves, directions)
		if everyoneStayed {
			return rounds
		}
		elves = move(elves, proposals)
		directions = append(directions[1:], directions[0])
	}
}
