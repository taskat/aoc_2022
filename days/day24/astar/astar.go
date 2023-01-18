package astar

import "container/heap"

type Pather interface {
	PathNeighbors() []Pather
	PathNeighborCost(to Pather) float64
	PathEstimatedCost(to Pather) float64
	GetHash() uint64
}

func Path(start Pather, isGoal func(p Pather) bool) (int, bool) {
	neighbors := &PriorityQueue{}
	found := make(map[uint64]int)
	heap.Init(neighbors)
	heap.Push(neighbors, &Node{pather: start, Priority: 0})
	for {
		if neighbors.Len() == 0 {
			return 0, false
		}
		current := heap.Pop(neighbors).(*Node)
		if isGoal(current.pather) {
			return current.Cost, true
		}
		for _, possibility := range current.pather.PathNeighbors() {
			newHash := possibility.GetHash()
			oldCost, ok := found[newHash]
			cost := current.Cost + int(current.pather.PathNeighborCost(possibility))
			if !ok || oldCost > cost {
				found[newHash] = cost
				prio := cost + int(current.pather.PathEstimatedCost(possibility))
				heap.Push(neighbors, &Node{pather: possibility, Priority: int(prio), Cost: cost})
			}
		}
	}
}