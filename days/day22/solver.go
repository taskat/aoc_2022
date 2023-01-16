package day22

import (
	"aoc_2022/days/day22/board"
	"aoc_2022/days/day22/direction"
	"fmt"
	"strconv"
	"strings"
)

type Solver struct{}

func (*Solver) SolvePart1(input string, extraParams ...any) string {
	sideLength := getSideLength(extraParams...)
	board, commands := parseInput(input, sideLength, false)
	state := newState(board)
	state.executeAll(commands)
	pos := state.b.GetAbsolutePosition(state.pos)
	password := 1000*(pos.Row+1) + 4*(pos.Col+1) + int(state.facing)
	return strconv.Itoa(password)
}

func parseInput(input string, sideLength int, cube bool) (board.Board, []command) {
	parts := strings.Split(input, "\n\n")
	b := board.ParseBoard(parts[0], sideLength, cube)
	commands := parseCommands(parts[1])
	return b, commands
}

type command interface {
	execute(s *state)
	fmt.Stringer
}

type state struct {
	pos    board.Position
	facing direction.Direction
	b      board.Board
}

func newState(b board.Board) *state {
	var s state
	s.b = b
	s.pos = b.GetFirst()
	s.facing = direction.Right
	return &s
}

func (s *state) executeAll(commands []command) {
	for _, c := range commands {
		c.execute(s)
	}
}

type moveCommand int

func (m moveCommand) execute(s *state) {
	for i := 0; i < int(m); i++ {
		newPos, newDir := s.b.Move(s.pos, s.facing)
		if newPos == s.pos {
			break
		}
		s.pos = newPos
		s.facing = newDir
	}
}

func (m moveCommand) String() string {
	return fmt.Sprintf("move %d", m)
}

type turnCommand bool

func newTurnRight() turnCommand {
	return turnCommand(true)
}

func newTurnLeft() turnCommand {
	return turnCommand(false)
}

func (t turnCommand) execute(s *state) {
	if t {
		s.facing = s.facing.TurnRight()
	} else {
		s.facing = s.facing.TurnLeft()
	}
}

func (t turnCommand) String() string {
	if t {
		return "turn right"
	}
	return "turn left"
}

func parseCommands(line string) []command {
	commands := make([]command, 0)
	numberStart := 0
	for i, c := range line {
		if c == 'R' || c == 'L' {
			numberString := line[numberStart:i]
			number, err := strconv.Atoi(numberString)
			if err != nil {
				panic(err)
			}
			commands = append(commands, moveCommand(number))
			if c == 'R' {
				commands = append(commands, newTurnRight())
			} else {
				commands = append(commands, newTurnLeft())
			}
			numberStart = i + 1
		}
	}
	numberString := line[numberStart:]
	number, err := strconv.Atoi(numberString)
	if err != nil {
		panic(err)
	}
	commands = append(commands, moveCommand(number))
	return commands
}

func (*Solver) SolvePart2(input string, extraParams ...any) string {
	sideLength := getSideLength(extraParams...)
	board, commands := parseInput(input, sideLength, true)
	state := newState(board)
	state.executeAll(commands)
	pos := state.b.GetAbsolutePosition(state.pos)
	fmt.Println(pos.Row, pos.Col, state.facing)
	password := 1000*(pos.Row+1) + 4*(pos.Col+1) + int(state.facing)
	return strconv.Itoa(password)
}

func getSideLength(params ...any) int {
	if len(params) == 0 {
		return 50
	}
	sideLenght, err := strconv.Atoi(params[0].(string))
	if err != nil {
		panic(err)
	}
	return sideLenght
}
