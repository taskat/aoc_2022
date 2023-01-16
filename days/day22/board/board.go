package board

import (
	"aoc_2022/days/day22/direction"
	"fmt"
	"strings"
)

type field bool

func (f field) String() string {
	if f {
		return "."
	}
	return "#"
}

type block struct {
	fields [][]field
	id     int
}

func newBlock(fields [][]field, id int) block {
	return block{fields, id}
}

type adjacentSide struct {
	block    int
	side     direction.Direction
	reversed bool
}

type Board struct {
	blocks     [][]block
	neighbors  [6]map[direction.Direction]adjacentSide
	sideLength int
}

type Position struct {
	Block, Row, Col int
}

func (p Position) String() string {
	return fmt.Sprintf("(%d, %d, %d)", p.Block, p.Row, p.Col)
}

func maxLineLength(lines []string, sideLength int) int {
	max := 0
	for i := 0; i < len(lines); i += sideLength {
		if len(lines[i]) > max {
			max = len(lines[i])
		}
	}
	return max
}

func createBlocks(input string, sideLength int) [][]block {
	lines := strings.Split(input, "\n")
	blocks := make([][]block, len(lines)/sideLength)
	blockCounter := 0
	maxLength := maxLineLength(lines, sideLength)
	for i := range blocks {
		blocks[i] = make([]block, maxLength/sideLength)
		for j := range blocks[i] {
			blocks[i][j] = newBlock(nil, -1)
		}
	}
	for blockI := 0; blockI < len(lines); blockI += sideLength {
		for blockJ := 0; blockJ < len(lines[blockI]); blockJ += sideLength {
			if lines[blockI][blockJ] == ' ' {
				continue
			}
			b := createBlock(lines, blockI, blockJ, sideLength)
			blocks[blockI/sideLength][blockJ/sideLength] = newBlock(b, blockCounter)
			blockCounter++
		}
	}
	return blocks
}

func createBlock(lines []string, row, col, sideLength int) [][]field {
	b := make([][]field, sideLength)
	for i := range b {
		b[i] = make([]field, sideLength)
		for j := range b[i] {
			b[i][j] = field(lines[row+i][col+j] == '.')
		}
	}
	return b
}

func createNeighbors(blocks [][]block, cube bool) [6]map[direction.Direction]adjacentSide {
	neighbors := [6]map[direction.Direction]adjacentSide{}
	for i := range neighbors {
		neighbors[i] = make(map[direction.Direction]adjacentSide)
	}
	addPhysicalNeighbors(neighbors, blocks)
	if cube {
		addCubeNeighbors(neighbors, blocks)
	} else {
		addWrapNeighbors(neighbors, blocks)
	}
	return neighbors
}

func addPhysicalNeighbors(neighbors [6]map[direction.Direction]adjacentSide, blocks [][]block) {
	for i, line := range blocks {
		for j, b := range line {
			if b.id == -1 {
				continue
			}
			if j < len(blocks[i])-1 {
				if blocks[i][j+1].id != -1 {
					neighbors[blocks[i][j].id][direction.Right] = adjacentSide{blocks[i][j+1].id, direction.Left, false}
					neighbors[blocks[i][j+1].id][direction.Left] = adjacentSide{blocks[i][j].id, direction.Right, false}
				}
			}
			if i < len(blocks)-1 {
				if blocks[i+1][j].id != -1 {
					neighbors[blocks[i][j].id][direction.Down] = adjacentSide{blocks[i+1][j].id, direction.Up, false}
					neighbors[blocks[i+1][j].id][direction.Up] = adjacentSide{blocks[i][j].id, direction.Down, false}
				}
			}
		}
	}
}

func addWrapNeighbors(neighbors [6]map[direction.Direction]adjacentSide, blocks [][]block) {
	for i, row := range blocks {
		for j, block := range row {
			if block.id == -1 {
				continue
			}
			if _, contains := neighbors[block.id][direction.Right]; !contains {
				for k := 0; k < len(blocks[i]); k++ {
					if blocks[i][k].id != -1 {
						neighbors[block.id][direction.Right] = adjacentSide{blocks[i][k].id, direction.Left, false}
						neighbors[blocks[i][k].id][direction.Left] = adjacentSide{block.id, direction.Right, false}
						break
					}
				}
			}
			if _, contains := neighbors[block.id][direction.Down]; !contains {
				for k := 0; k < len(blocks); k++ {
					if blocks[k][j].id != -1 {
						neighbors[block.id][direction.Down] = adjacentSide{blocks[k][j].id, direction.Up, false}
						neighbors[blocks[k][j].id][direction.Up] = adjacentSide{block.id, direction.Down, false}
						break
					}
				}
			}
		}
	}
}

func getReversed() map[direction.Direction]map[direction.Direction]bool {
	directions := []direction.Direction{direction.Up, direction.Right, direction.Down, direction.Left}
	reversed := make(map[direction.Direction]map[direction.Direction]bool)
	for _, dir := range directions {
		reversed[dir] = make(map[direction.Direction]bool)
	}
	reversed[direction.Up][direction.Up] = true
	reversed[direction.Up][direction.Right] = true
	reversed[direction.Up][direction.Down] = false
	reversed[direction.Up][direction.Left] = false
	reversed[direction.Right][direction.Up] = true
	reversed[direction.Right][direction.Right] = true
	reversed[direction.Right][direction.Down] = false
	reversed[direction.Right][direction.Left] = false
	reversed[direction.Down][direction.Up] = false
	reversed[direction.Down][direction.Right] = false
	reversed[direction.Down][direction.Down] = true
	reversed[direction.Down][direction.Left] = true
	reversed[direction.Left][direction.Up] = false
	reversed[direction.Left][direction.Right] = false
	reversed[direction.Left][direction.Down] = true
	reversed[direction.Left][direction.Left] = true
	return reversed
}

func addCubeNeighbors(neighbors [6]map[direction.Direction]adjacentSide, blocks [][]block) {
	directions := []direction.Direction{direction.Left, direction.Up, direction.Right, direction.Down}
	isReversed := getReversed()
	for _, row := range blocks {
		for _, block := range row {
			if block.id == -1 {
				continue
			}
			for _, dir := range directions {
				if _, contains := neighbors[block.id][dir]; !contains {
					if firstNeighbor, ok := neighbors[block.id][dir.TurnLeft()]; ok {
						side := firstNeighbor.side.TurnLeft()
						if secondNeighbor, ok := neighbors[firstNeighbor.block][side]; ok {
							side := secondNeighbor.side.TurnLeft()
							reversed := isReversed[dir][side]
							neighbors[block.id][dir] = adjacentSide{secondNeighbor.block, side, reversed}
							neighbors[secondNeighbor.block][side] = adjacentSide{block.id, dir, reversed}
						}
					}
					if firstNeighbor, ok := neighbors[block.id][dir.TurnRight()]; ok {
						side := firstNeighbor.side.TurnRight()
						if secondNeighbor, ok := neighbors[firstNeighbor.block][side]; ok {
							side := secondNeighbor.side.TurnRight()
							reversed := isReversed[dir][side]
							neighbors[block.id][dir] = adjacentSide{secondNeighbor.block, side, reversed}
							neighbors[secondNeighbor.block][side] = adjacentSide{block.id, dir, reversed}
						}
					}
				}
			}
		}
	}
}

func ParseBoard(input string, sideLength int, cube bool) Board {
	blocks := createBlocks(input, sideLength)
	neighbors := createNeighbors(blocks, cube)
	return Board{blocks, neighbors, sideLength}
}

func (b Board) String() string {
	lines := make([]string, 0)
	for i, m := range b.neighbors {
		for dir, adj := range m {
			lines = append(lines, fmt.Sprintf("block %d %s: %d %s", i, direction.Direction(dir), adj.block, adj.side))
		}
	}
	return strings.Join(lines, "\n")
}

func (b *Board) get(p Position) field {
	i, j := b.getBlock(p.Block)
	return b.blocks[i][j].fields[p.Row][p.Col]
}

func (b *Board) GetAbsolutePosition(p Position) Position {
	i, j := b.getBlock(p.Block)
	return Position{Row: i*b.sideLength + p.Row, Col: j*b.sideLength + p.Col, Block: p.Block}
}

func (b *Board) getBlock(blockId int) (int, int) {
	for i, row := range b.blocks {
		for j, block := range row {
			if block.id == blockId {
				return i, j
			}
		}
	}
	panic("block not found")
}

func (b *Board) GetFirst() Position {
	p := Position{0, 0, 0}
	for !b.get(p) {
		p.Col++
	}
	return p
}

func (b *Board) GetSideLength() int {
	return b.sideLength
}

func (b *Board) jump(p Position, d direction.Direction) (Position, direction.Direction) {
	adj := b.neighbors[p.Block][d]
	var newRow, newCol, from, goal int
	sum := b.sideLength - 1
	switch d {
	case direction.Right, direction.Left:
		from = p.Row
	case direction.Down, direction.Up:
		from = p.Col
	}
	if adj.reversed {
		goal = sum - from
	} else {
		goal = from
	}
	switch adj.side {
	case direction.Right:
		newCol = sum
		newRow = goal
	case direction.Down:
		newCol = goal
		newRow = sum
	case direction.Left:
		newCol = 0
		newRow = goal
	case direction.Up:
		newCol = goal
		newRow = 0
	}
	newPos := Position{adj.block, newRow, newCol}
	if !b.get(newPos) {
		return p, d
	}
	newDir := adj.side.Reverse()
	return newPos, newDir
}

func (b *Board) Move(from Position, d direction.Direction) (Position, direction.Direction) {
	switch d {
	case direction.Right:
		if from.Col == b.sideLength-1 {
			return b.jump(from, d)
		}
		newPos := Position{from.Block, from.Row, from.Col + 1}
		if !b.get(newPos) {
			return from, d
		}
		return newPos, d
	case direction.Down:
		if from.Row == b.sideLength-1 {
			return b.jump(from, d)
		}
		newPos := Position{from.Block, from.Row + 1, from.Col}
		if !b.get(newPos) {
			return from, d
		}
		return newPos, d
	case direction.Left:
		if from.Col == 0 {
			return b.jump(from, d)
		}
		newPos := Position{from.Block, from.Row, from.Col - 1}
		if !b.get(newPos) {
			return from, d
		}
		return newPos, d
	case direction.Up:
		if from.Row == 0 {
			return b.jump(from, d)
		}
		newPos := Position{from.Block, from.Row - 1, from.Col}
		if !b.get(newPos) {
			return from, d
		}
		return newPos, d
	default:
		panic("invalid direction.Direction")
	}
}
