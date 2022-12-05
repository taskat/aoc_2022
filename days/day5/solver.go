package day5

import (
	"aoc_2022/days/day5/modern_stack"
	"aoc_2022/days/day5/stack"
	"fmt"
	"strings"
)

type Solver struct {}

func (*Solver) SolvePart1(input string, extraParams ...any) string {
	stacks, commands := parseInput(input, initStacks)
	ship := newShip(stacks)
	ship.executeCommands(commands)
	return ship.getTopOfStack()
}

type stacker interface {
	Push(r rune)
	PushAll(r []rune)
	Pop(n int) []rune
	Top() rune
}

func parseInput(input string, initStacks func(n int) []stacker) ([]stacker, []command) {
	parts := strings.Split(input, "\n\n")
	stacks := parseStacks(parts[0], initStacks)
	commands := parseCommands(parts[1])
	return stacks, commands
}

func initStacks(n int) []stacker {
	stacks := make([]stacker, n)
	for i := 0; i < n; i++ {
		stacks[i] = stack.NewStack()
	}
	return stacks
}

func parseStacks(input string, initStacks func(n int) []stacker) []stacker {
	lines := strings.Split(input, "\n")
	lastLine := lines[len(lines)-1]
	stacks := initStacks((len(lastLine) + 1) / 4)
	for i := len(lines) - 2; i >= 0; i-- {
		line := lines[i]
		for j := 0; j < len(line); j += 4 {
			if line[j + 1] == ' ' {
				continue
			}
			stacks[j / 4].Push(rune(line[j + 1]))
		}
	}
	return stacks
}

type command struct {
	quantity int
	source int
	destination int
}

func parseCommand(line string) command {
	parts := strings.Split(line, " ")
	quantity := 0
	source := 0
	destination := 0
	fmt.Sscanf(parts[1], "%d", &quantity)
	fmt.Sscanf(parts[3], "%d", &source)
	fmt.Sscanf(parts[5], "%d", &destination)
	//so it will match the indeces
	source--
	destination--
	return command{quantity, source, destination}
}

func parseCommands(input string) []command {
	lines := strings.Split(input, "\n")
	commands := make([]command, len(lines))
	for i := 0; i < len(lines); i++ {
		commands[i] = parseCommand(lines[i])
	}
	return commands
}

type ship struct {
	stacks []stacker
}

func newShip(stacks []stacker) *ship {
	return &ship{stacks: stacks}
}

func (s *ship) executeCommand(command command) {
	s.stacks[command.destination].PushAll(s.stacks[command.source].Pop(command.quantity))
}

func (s *ship) executeCommands(commands []command) {
	for _, command := range commands {
		s.executeCommand(command)
	}
}

func (s *ship) getTopOfStack() string {
	result := ""
	for _, stack := range s.stacks {
		result += string(stack.Top())
	}
	return result
}

func (*Solver) SolvePart2(input string, extraParams ...any) string {
	stacks, commands := parseInput(input, initModernStacks)
	ship := newShip(stacks)
	ship.executeCommands(commands)
	return ship.getTopOfStack()
}

func initModernStacks(n int) []stacker {
	stacks := make([]stacker, n)
	for i := 0; i < n; i++ {
		stacks[i] = modern_stack.NewStack()
	}
	return stacks
}
