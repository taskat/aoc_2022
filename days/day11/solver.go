package day11

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Solver struct{}

func (*Solver) SolvePart1(input string, extraParams ...any) string {
	monkeys := parseMonkeys(input)
	for i := 0; i < 20; i++ {
		round(monkeys, true)
	}
	return strconv.Itoa(getMaxMonkeyBusiness(monkeys))
}

func parseMonkeys(input string) []monkey {
	lines := strings.Split(input, "\n\n")
	monkeys := make([]monkey, len(lines))
	for i, line := range lines {
		monkeys[i] = *parseMonkey(line)
	}
	return monkeys
}

type monkey struct {
	items []int
	op    func(int) int
	test  func(int) int

	numberOfInspections int
	optimizationNumber  int
}

func newMonkey(items []int, op, test func(int) int) *monkey {
	return &monkey{items, op, test, 0, -1}
}

func parseMonkey(data string) *monkey {
	lines := strings.Split(data, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}
	itemsData := strings.Split(lines[1], ": ")[1]
	itemStrings := strings.Split(itemsData, ", ")
	items := make([]int, len(itemStrings))
	for i, item := range itemStrings {
		items[i], _ = strconv.Atoi(item)
	}
	opData := strings.Split(lines[2], " = ")
	opParts := strings.Split(opData[1], " ")
	var op func(item int) int
	if opParts[1] == "+" {
		if opParts[2] == "old" {
			op = func(item int) int { return item + item }
		} else {
			operand, _ := strconv.Atoi(opParts[2])
			op = func(item int) int { return item + operand }
		}
	} else {
		if opParts[2] == "old" {
			op = func(item int) int { return item * item }
		} else {
			operand, _ := strconv.Atoi(opParts[2])
			op = func(item int) int { return item * operand }
		}
	}
	testOperandString := strings.Split(lines[3], " ")[3]
	testOperand, _ := strconv.Atoi(testOperandString)
	targetIfTrue, _ := strconv.Atoi(strings.Split(lines[4], " ")[5])
	targetIfFalse, _ := strconv.Atoi(strings.Split(lines[5], " ")[5])
	test := func(item int) int {
		if item%testOperand == 0 {
			return targetIfTrue
		} else {
			return targetIfFalse
		}
	}
	return newMonkey(items, op, test)
}

func (m *monkey) addItem(item int) {
	m.items = append(m.items, item)
}

func (m *monkey) turn(monkeys *[]monkey, relief bool) {
	m.numberOfInspections += len(m.items)
	for _, item := range m.items {
		newItem := m.op(item)
		if relief {
			newItem /= 3
		}
		if m.optimizationNumber != -1 {
			newItem %= m.optimizationNumber
		}
		(*monkeys)[m.test(newItem)].addItem(newItem)
	}
	m.items = []int{}
}

func (m *monkey) setOptimizationNumber(number int) {
	m.optimizationNumber = number
}

func (m *monkey) String() string {
	return fmt.Sprintf("%v, (%d)", m.items, m.numberOfInspections)
}

func round(monkeys []monkey, relief bool) {
	for i := range monkeys {
		monkeys[i].turn(&monkeys, relief)
	}
}

func getMaxMonkeyBusiness(monkeys []monkey) int {
	inspections := make([]int, len(monkeys))
	for i := range monkeys {
		inspections[i] = monkeys[i].numberOfInspections
	}
	sort.Ints(inspections)
	return inspections[len(inspections)-1] * inspections[len(inspections)-2]
}

func (*Solver) SolvePart2(input string, extraParams ...any) string {
	monkeys := parseMonkeys(input)
	optimizationNumber := findWorryOptimizationNumber(input)
	for i := range monkeys {
		monkeys[i].setOptimizationNumber(optimizationNumber)
	}
	for i := 0; i < 10_000; i++ {
		round(monkeys, false)
	}
	return strconv.Itoa(getMaxMonkeyBusiness(monkeys))
}

func findWorryOptimizationNumber(input string) int {
	blocks := strings.Split(input, "\n\n")
	numbers := make([]int, len(blocks))
	for i, block := range blocks {
		lines := strings.Split(block, "\n")
		testLine := strings.TrimSpace(lines[3])
		parts := strings.Split(testLine, " ")
		numbers[i], _ = strconv.Atoi(parts[3])
	}
	product := 1
	for _, number := range numbers {
		product *= number
	}
	return product
}
