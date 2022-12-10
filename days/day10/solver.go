package day10

import (
	"strconv"
	"strings"
)

type Solver struct {}

func (*Solver) SolvePart1(input string, extraParams ...any) string {
	commands := parseCommands(input)
	cpu := newCpu([]int{20, 60, 100, 140, 180, 220})
	for _, command := range commands {
		command.Execute(cpu)
	}
	signals := cpu.specificSignals
	return strconv.Itoa(sumSignalStrength(signals))
}

type CPU struct {
	x int
	cycles int

	specificSignals map[int]int
	screen [6][40]bool
}

func newCpu(specificCycles []int) *CPU {
	m := make(map[int]int, len(specificCycles))
	for _, cycle := range specificCycles {
		m[cycle] = -1
	}
	return &CPU{x: 1, cycles: 0, specificSignals: m}
}

func (cpu *CPU) draw() {
	value := false
	if cpu.x - (cpu.cycles % 40) >= -1 && cpu.x - (cpu.cycles % 40) <= 1 {
		value = true
	}
	cpu.screen[cpu.cycles / 40][cpu.cycles % 40] = value
}

func (cpu *CPU) tick() {
	for key := range cpu.specificSignals {
		if cpu.cycles + 1 == key {
			cpu.specificSignals[key] = cpu.x
		}
	}
	cpu.draw()
	cpu.cycles++
}

func (cpu *CPU) print() {
	for _, row := range cpu.screen {
		for _, value := range row {
			if value {
				print("#")
			} else {
				print(".")
			}
		}
		println()
	}
}

type command interface {
	Execute(cpu *CPU)
}

type noop struct {}

func (n *noop) Execute(cpu *CPU) {
	cpu.tick()
}

type addx struct {
	v int
}

func (a *addx) Execute(cpu *CPU) {
	cpu.tick()
	cpu.tick()
	cpu.x += a.v
}

func parseCommands(input string) []command {
	lines := strings.Split(input, "\n")
	commands := make([]command, len(lines))
	for i, line := range lines {
		if line == "noop" {
			commands[i] = &noop{}
		} else {
			parts := strings.Split(line, " ")
			v, _ := strconv.Atoi(parts[1])
			commands[i] = &addx{v}
		}
	}
	return commands
}

func sumSignalStrength(m map[int]int) int {
	sum := 0
	for k, v := range m {
		sum += k * v
	}
	return sum
}

func (*Solver) SolvePart2(input string, extraParams ...any) string {
	commands := parseCommands(input)
	cpu := newCpu(nil)
	for _, command := range commands {
		command.Execute(cpu)
	}
	cpu.print()
	return "See above"
}
