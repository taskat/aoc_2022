package day16

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Solver struct{}

func (*Solver) SolvePart1(input string, extraParams ...any) string {
	flowRates, p, sortedFlowRates, startNode := parseAndPrepare(input)
	bestForVisited := make(map[int]int)
	best := branchAndBound(flowRates, sortedFlowRates, p,
		newState(startNode, 30), &bestForVisited, 0, func(bound, best int) bool {
		return bound > best
	})
	return strconv.Itoa(best)
}

type valve struct {
	name      string
	flow      int
	neighbors []string
}

func (v valve) interesting() bool {
	return v.flow != 0 || v.name == "AA"
}

func parseInput(input string) []valve {
	lines := strings.Split(input, "\n")
	valves := make([]valve, len(lines))
	for i, line := range lines {
		valves[i] = parseValve(line)
	}
	return valves
}

func parseValve(line string) valve {
	words := strings.Split(line, " ")
	name := words[1]
	flowString := strings.TrimPrefix(words[4], "rate=")
	flowString = strings.TrimSuffix(flowString, ";")
	flow, _ := strconv.Atoi(flowString)
	neighbors := strings.Split(strings.Join(words[9:], ""), ",")
	return valve{name, flow, neighbors}
}

func (v valve) String() string {
	return fmt.Sprintf("%s: %d, [%s]", v.name, v.flow, strings.Join(v.neighbors, ", "))
}

func floyd_warshall(nodes []valve) [][]int {
	nameToIdx := make(map[string]int, len(nodes))
	for i, node := range nodes {
		nameToIdx[node.name] = i
	}
	distances := make([][]int, len(nodes))
	for i := range nodes {
		distances[i] = make([]int, len(nodes))
		for j := range nodes {
			distances[i][j] = 100
		}
	}
	for i, node := range nodes {
		for _, neighbor := range node.neighbors {
			distances[i][nameToIdx[neighbor]] = 1
		}
	}
	for i := range nodes {
		distances[i][i] = 0
	}
	for k := range nodes {
		for i := range nodes {
			for j := range nodes {
				dist := distances[i][k] + distances[k][j]
				if dist < distances[i][j] {
					distances[i][j] = dist
				}
			}
		}
	}
	return distances
}

func parseAndPrepare(input string) (flowRates []int, paths [][]int, sortedFlowRateIndices []int, startNode int) {
	valves := parseInput(input)
	p := floyd_warshall(valves)
	interestingIndices := make([]int, 0)
	for i, valve := range valves {
		if valve.interesting() {
			interestingIndices = append(interestingIndices, i)
		}
	}
	flowRates = make([]int, len(interestingIndices))
	for i, idx := range interestingIndices {
		flowRates[i] = valves[idx].flow
	}
	paths = make([][]int, len(interestingIndices))
	for i, idx := range interestingIndices {
		paths[i] = make([]int, len(interestingIndices))
		for j, idx2 := range interestingIndices {
			paths[i][j] = p[idx][idx2]
		}
	}
	sortedPairs := sortPairs(flowRates)
	sortedFlowRateIndices = make([]int, len(sortedPairs))
	for i, pair := range sortedPairs {
		sortedFlowRateIndices[i] = pair.idx
	}
	for i, idx := range interestingIndices {
		if valves[idx].name == "AA" {
			startNode = i
			break
		}
	}
	return flowRates, paths, sortedFlowRateIndices, startNode
}

type state struct {
	visited          int
	avoid            int
	pressureReleased int
	minutesRemaining int
	position         int
}

func newState(position, minutesRemaining int) state {
	return state{0, 1 << position, 0, minutesRemaining, position}
}

func (s *state) canVisit(i int) bool {
	return (s.visited|s.avoid)&(1<<i) == 0
}

func (s *state) bound(flowRates, sortedFlowRateIndices []int) int {
	released := s.pressureReleased
	c := (s.minutesRemaining - 1) / 2
	if c < 0 {
		c = 0
	}
	minutes := make([]int, 0, c)
	for i := s.minutesRemaining - 1; i >= 0; i -= 2 {
		minutes = append(minutes, i)
	}
	validSortedFlowRateIndices := make([]int, 0, len(sortedFlowRateIndices))
	for _, idx := range sortedFlowRateIndices {
		if s.canVisit(idx) {
			validSortedFlowRateIndices = append(validSortedFlowRateIndices, flowRates[idx])
		}
	}
	for i := 0; i < len(minutes) && i < len(validSortedFlowRateIndices); i++ {
		released += minutes[i] * validSortedFlowRateIndices[i]
	}
	return released
}

func (s *state) branch(flowRates []int, paths [][]int) []state {
	branches := make([]state, 0)
	for dest, dist := range paths[s.position] {
		if !s.canVisit(dest) {
			continue
		}
		remainingMinutes := s.minutesRemaining - dist - 1
		var newState state
		newState.visited = s.visited | (1 << dest)
		newState.avoid = s.avoid
		newState.pressureReleased = s.pressureReleased + remainingMinutes*flowRates[dest]
		newState.minutesRemaining = remainingMinutes
		newState.position = dest
		branches = append(branches, newState)
	}
	return branches
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type bb struct {
	bound int
	branch state
}

func branchAndBound(flowRates, sortedFlowRateIndices []int, paths [][]int, s state,
	bestForVisited *map[int]int, best int, filterBound func(int, int) bool) int {
	if currentBest, ok := (*bestForVisited)[s.visited]; ok {
		(*bestForVisited)[s.visited] = max(currentBest, s.pressureReleased)
	}
	best = max(best, s.pressureReleased)
	boundBranchPairs := make([]bb, 0)
	for _, branch := range s.branch(flowRates, paths) {
		bound := branch.bound(flowRates, sortedFlowRateIndices)
		if !filterBound(bound, best) {
			continue
		}
		boundBranchPairs = append(boundBranchPairs, bb{bound, branch})
	}
	sort.Slice(boundBranchPairs, func(i, j int) bool {
		return boundBranchPairs[i].bound > boundBranchPairs[j].bound
	})
	for _, bb := range boundBranchPairs {
		if !filterBound(bb.bound, best) {
			continue
		}
		best = branchAndBound(flowRates, sortedFlowRateIndices, paths, bb.branch, bestForVisited, best, filterBound)
	}
	return best
}

type ib struct {
	idx, best int
}

func (*Solver) SolvePart2(input string, extraParams ...any) string {
	flowRates, p, sortedFlowRates, startNode := parseAndPrepare(input)
	bpvLen := 1 << len(flowRates)
	bestPerVisited := make(map[int]int, bpvLen)
	for i := 0; i < bpvLen; i++ {
		bestPerVisited[i] = -1
	}
	branchAndBound(flowRates, sortedFlowRates, p, newState(startNode, 26), &bestPerVisited, 0, func(bound, best int) bool {
		return bound > best * 3 / 4
	})
	bestPerVisitedSorted := make([]ib, 0)
	for i, best := range bestPerVisited {
		if best <= 0 {
			continue
		}
		bestPerVisitedSorted = append(bestPerVisitedSorted, ib{i, best})
	}
	sort.Slice(bestPerVisitedSorted, func(i, j int) bool {
		return bestPerVisitedSorted[i].best > bestPerVisitedSorted[j].best
	})
	best := 0
	for i, ib := range bestPerVisitedSorted {
		myVisited := ib.idx
		myBest := ib.best
		for j := i + 1; j < len(bestPerVisitedSorted); j++ {
			elephantVisited := bestPerVisitedSorted[j].idx
			elephantBest := bestPerVisitedSorted[j].best
			score := myBest + elephantBest
			if score <= best {
				break
			}
			if myVisited & elephantVisited == 0 {
				best = score
			}
		}
	}
	return strconv.Itoa(best)
}
