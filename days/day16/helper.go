package day16

import "sort"

type pair struct {
	idx, value int
}

type pairs []pair

func (p pairs) Len() int {
	return len(p)
}

func (p pairs) Less(i, j int) bool {
	return p[i].value > p[j].value
}

func (p pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func sortPairs(flowRates []int) []pair {
	pairs := make(pairs, len(flowRates))
	for i, flowRate := range flowRates {
		pairs[i] = pair{i, flowRate}
	}
	sort.Sort(pairs)
	return pairs
}
