package day19

import (
	"fmt"
	"strconv"
	"strings"
)

type Solver struct{}

func (*Solver) SolvePart1(input string, extraParams ...any) string {
	minutes := 24
	bluePrints := parseInput(input)
	qualityLevel := qualityLevel(bluePrints, minutes)
	return strconv.Itoa(qualityLevel)
}

type resource int

const (
	GEODE resource = iota
	OBSIDIAN
	CLAY
	ORE
	numberOfResources
)

func parseResource(s string) resource {
	switch s {
	case "ore":
		return ORE
	case "clay":
		return CLAY
	case "obsidian":
		return OBSIDIAN
	case "geode":
		return GEODE
	default:
		panic("invalid resource")
	}
}

func (r resource) String() string {
	switch r {
	case ORE:
		return "ore"
	case CLAY:
		return "clay"
	case OBSIDIAN:
		return "obsidian"
	case GEODE:
		return "geode"
	default:
		panic("invalid resource")
	}
}

type bluePrint struct {
	robotCosts [numberOfResources][numberOfResources]int
}

func parseBluePrint(line string) bluePrint {
	line = strings.Split(line, ":")[1]
	line = strings.TrimSuffix(line, ".")
	sentences := strings.Split(line, ".")
	var robotCosts [numberOfResources][numberOfResources]int
	for _, sentence := range sentences {
		sentence = strings.TrimSpace(sentence)
		sentence = strings.TrimSuffix(sentence, ".")
		words := strings.Split(sentence, " ")
		robotType := parseResource(words[1])
		robotCosts[robotType] = [numberOfResources]int{}
		resourceAmount := atoi(words[4])
		resourceType := parseResource(words[5])
		robotCosts[robotType][resourceType] = resourceAmount
		if len(words) > 6 {
			resourceAmount = atoi(words[7])
			resourceType = parseResource(words[8])
			robotCosts[robotType][resourceType] = resourceAmount
		}
	}
	return bluePrint{robotCosts}
}

func (b *bluePrint) canBuild(robot resource, res [numberOfResources]int) bool {
	for resourceType, resourceAmount := range b.robotCosts[robot] {
		if res[resourceType] < resourceAmount {
			return false
		}
	}
	return true
}

func (b *bluePrint) shouldWait(robot resource, robots [numberOfResources]int) bool {
	for resourceType, resourceAmount := range b.robotCosts[robot] {
		if resourceAmount > 0 && robots[resourceType] == 0 {
			return false
		}
	}
	return true
}

func (b bluePrint) String() string {
	robots := make([]string, 0, len(b.robotCosts))
	for robotType, robotCost := range b.robotCosts {
		s := resource(robotType).String() + ": "
		cost := make([]string, 0, len(robotCost))
		for resourceType, resourceAmount := range robotCost {
			cost = append(cost, strconv.Itoa(resourceAmount)+" "+resource(resourceType).String())
		}
		s += strings.Join(cost, ", ")
		robots = append(robots, s)
	}
	return strings.Join(robots, "\n")
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func parseInput(input string) []bluePrint {
	lines := strings.Split(input, "\n")
	bluePrints := make([]bluePrint, len(lines))
	for i, line := range lines {
		bluePrints[i] = parseBluePrint(line)
	}
	return bluePrints
}

type state struct {
	robots      [numberOfResources]int
	resources   [numberOfResources]int
	bp          bluePrint
	max         [numberOfResources]int
	wasPossible [numberOfResources]bool
}

func newState(bp bluePrint) state {
	state := state{bp: bp}
	state.robots[ORE] = 1
	for _, cost := range bp.robotCosts {
		for resourceType, resourceAmount := range cost {
			if resourceAmount > state.max[resourceType] {
				state.max[resourceType] = resourceAmount
			}
		}
	}
	return state
}

func (s *state) build(robot resource) {
	s.robots[robot]++
	for res, amount := range s.bp.robotCosts[robot] {
		s.resources[res] -= amount
	}
}

func (s *state) collectResources() {
	for robotType, robotAmount := range s.robots {
		s.resources[robotType] += robotAmount
	}
}

func (s *state) nextStates(onlyGeode bool) []state {
	if onlyGeode {
		newState := *s
		newState.wasPossible = [numberOfResources]bool{}
		if s.bp.canBuild(GEODE, s.resources) {
			newState.collectResources()
			newState.build(GEODE)
		} else {
			newState.collectResources()
		}
		return []state{newState}
	}
	states := make([]state, 0, numberOfResources+1)
	var wasPossible [numberOfResources]bool
	shouldWait := false
	for robotType := range s.robots {
		if (robotType != int(GEODE) && s.max[robotType] <= s.robots[robotType]) || s.wasPossible[robotType] {
			continue
		}
		if !s.bp.canBuild(resource(robotType), s.resources) {
			if !shouldWait {
				shouldWait = s.bp.shouldWait(resource(robotType), s.robots)
			}
			continue
		}
		newState := *s
		newState.wasPossible = [numberOfResources]bool{}
		newState.collectResources()
		newState.build(resource(robotType))
		states = append(states, newState)
		wasPossible[robotType] = true
	}
	if shouldWait {
		s.collectResources()
		s.wasPossible = wasPossible
		states = append(states, *s)
	}
	return states
}

func (s state) String() string {
	robots := make([]string, 0, len(s.robots))
	for robotType, robotAmount := range s.robots {
		robots = append(robots, strconv.Itoa(robotAmount)+" "+resource(robotType).String())
	}
	resources := make([]string, 0, len(s.resources))
	for resourceType, resourceAmount := range s.resources {
		resources = append(resources, strconv.Itoa(resourceAmount)+" "+resource(resourceType).String())
	}
	return "robots: " + strings.Join(robots, ", ") + " , resources: " + strings.Join(resources, ", ")
}

func (s *state) upperBound(minutes int) int {
	return minutes*(minutes+1)/2 + s.resources[GEODE] + s.robots[GEODE] * minutes
}

func dfs(from state, best, minutes int) int {
	if minutes == 1 {
		from.collectResources()
		if from.resources[GEODE] > best {
			return from.resources[GEODE]
		} else {
			return best
		}
	}
	onlyGeode := false
	if minutes == 2 {
		onlyGeode = true
	}
	for _, to := range from.nextStates(onlyGeode) {
		if to.upperBound(minutes-1) < best {
			continue
		}
		best = dfs(to, best, minutes-1)
	}
	return best
}

func qualityLevel(bluePrints []bluePrint, minutes int) int {
	sum := 0
	for i, bp := range bluePrints {
		start := newState(bp)
		best := dfs(start, 0, minutes)
		sum += best * (i + 1)
	}
	return sum
}

func (*Solver) SolvePart2(input string, extraParams ...any) string {
	minutes := 32
	bluePrints := parseInput(input)
	if len(bluePrints) > 3 {
		bluePrints = bluePrints[:3]
	}
	product := 1
	for i, bp := range bluePrints {
		best := dfs(newState(bp), 0, minutes)
		fmt.Println(i, best)
		product *= best
	}
	return strconv.Itoa(product)
}

