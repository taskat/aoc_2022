package day21

import (
	"strconv"
	"strings"
)

type Solver struct {}

func (*Solver) SolvePart1(input string, extraParams ...any) string {
	monkeys := parseInput(input)
	return strconv.Itoa(monkeys["root"].result(monkeys))
}

type monkey interface {
	result(monkeys map[string]monkey) int
	solve(mokeys map[string]monkey) (int, bool)
	calculateHuman(monkeys map[string]monkey, result int) int
}

type simpleMonkey int

func parseSimpleMonkey(s string) simpleMonkey {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return simpleMonkey(i)
}

func (m simpleMonkey) result(_ map[string]monkey) int {
	return int(m)
}

func (m simpleMonkey) solve(_ map[string]monkey) (int, bool) {
	return int(m), true
}

func (m simpleMonkey) calculateHuman(_ map[string]monkey, result int) int {
	return result
}

type operator string

const (
	add operator = "+"
	mul operator = "*"
	subtract operator = "-"
	divide operator = "/"
)

func parseOperator(s string) operator {
	switch s {
	case "+":
		return add
	case "*":
		return mul
	case "-":
		return subtract
	case "/":
		return divide
	default:
		panic("unknown operator")
	}
}

type complexMonkey struct {
	op operator
	left, right string
}

func parseComplexMonkey(parts []string) complexMonkey {
	return complexMonkey{
		op: parseOperator(parts[1]),
		left: parts[0],
		right: parts[2],
	}
}

func (m complexMonkey) result(monkeys map[string]monkey) int {
	leftMonkey := monkeys[m.left]
	rightMonkey := monkeys[m.right]
	left := leftMonkey.result(monkeys)
	right := rightMonkey.result(monkeys)
	monkeys[m.left] = simpleMonkey(left)
	monkeys[m.right] = simpleMonkey(right)
	switch m.op {
	case add:
		return left + right
	case mul:
		return left * right
	case subtract:
		return left - right
	case divide:
		return left / right
	default:
		panic("unknown operator")
	}
}

func (m complexMonkey) solve(monkeys map[string]monkey) (int, bool) {
	if m.left == "humn" || m.right == "humn" {
		return 0, false
	}
	left, leftSolved := monkeys[m.left].solve(monkeys)
	right, rightSolved := monkeys[m.right].solve(monkeys)
	if leftSolved {
		monkeys[m.left] = simpleMonkey(left)
	}
	if rightSolved {
		monkeys[m.right] = simpleMonkey(right)
	}
	if leftSolved && rightSolved {
		return m.result(monkeys), true
	}
	return 0, false
}

func (m complexMonkey) calculateHuman(monkeys map[string]monkey, result int) int {
	left, leftSolved := monkeys[m.left].solve(monkeys)
	right, rightSolved := monkeys[m.right].solve(monkeys)
	leftMonkey := monkeys[m.left]
	rightMonkey := monkeys[m.right]
	if leftSolved {
		switch m.op {
		case add:
			return rightMonkey.calculateHuman(monkeys, result - left)
		case mul:
			return rightMonkey.calculateHuman(monkeys, result / left)
		case subtract:
			return rightMonkey.calculateHuman(monkeys, left - result)
		case divide:
			return rightMonkey.calculateHuman(monkeys, left / result)
		default:
			panic("unknown operator")
		}
	} else if rightSolved {
		switch m.op {
		case add:
			return leftMonkey.calculateHuman(monkeys, result - right)
		case mul:
			return leftMonkey.calculateHuman(monkeys, result / right)
		case subtract:
			return leftMonkey.calculateHuman(monkeys, result + right)
		case divide:
			return leftMonkey.calculateHuman(monkeys, result * right)
		default:
			panic("unknown operator")
		}
	} else {
		panic("no solved monkeys")
	}
}

func parseInput(input string) map[string]monkey {
	lines := strings.Split(input, "\n")
	monkeys := make(map[string]monkey, len(lines))
	for _, line := range lines {
		parts := strings.Split(line, ":")
		name := parts[0]
		parts = strings.Split(strings.TrimPrefix(parts[1], " "), " ")
		if len(parts) == 1 {
			monkeys[name] = parseSimpleMonkey(parts[0])
		} else {
			monkeys[name] = parseComplexMonkey(parts)
		}
	}
	return monkeys
}

func (*Solver) SolvePart2(input string, extraParams ...any) string {
	monkeys := parseInput(input)
	updateMonkeys(monkeys)
	human := monkeys["root"].calculateHuman(monkeys, 0)
	return strconv.Itoa(human)
}

type rootMonkey struct {
	complexMonkey
}

func (m rootMonkey) result(monkeys map[string]monkey) int {
	return 0
}

func (m rootMonkey) solve(monkeys map[string]monkey) (int, bool) {
	return m.complexMonkey.solve(monkeys)
}

func (m rootMonkey) calculateHuman(monkeys map[string]monkey, result int) int {
	left, leftSolved := monkeys[m.left].solve(monkeys)
	right, rightSolved := monkeys[m.right].solve(monkeys)
	leftMonkey := monkeys[m.left]
	rightMonkey := monkeys[m.right]
	if leftSolved {
		return rightMonkey.calculateHuman(monkeys, left)
	} else if rightSolved {
		return leftMonkey.calculateHuman(monkeys, right)
	} else {
		panic("no solved monkeys")
	}
}

func updateMonkeys(monkeys map[string]monkey) {
	oldRoot := monkeys["root"]
	newRoot := rootMonkey{complexMonkey: oldRoot.(complexMonkey)}
	monkeys["root"] = newRoot
	monkeys["humn"] = humanMonkey{}
}

type humanMonkey struct {}

func (m humanMonkey) result(_ map[string]monkey) int {
	return 0
}

func (m humanMonkey) solve(_ map[string]monkey) (int, bool) {
	return 0, false
}

func (m humanMonkey) calculateHuman(_ map[string]monkey, result int) int {
	return result
}