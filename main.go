package main

import (
	"container/heap"
	"math"
)

type Node struct {
	ID     int
	Weight int
	Next   []Node
}

type Graph struct {
	Nodes []Node
}

// BuildGraph node->edges [][2]int
func BuildGraph(nodes int, edgeWeights [][][2]int) *Graph {
	g := &Graph{
		Nodes: make([]Node, nodes),
	}
	for node, edgeWeight := range edgeWeights {
		for _, edge := range edgeWeight {
			nextNode, weight := edge[0], edge[1]
			g.Nodes[node].Next = append(g.Nodes[node].Next, Node{
				ID:     nextNode,
				Weight: weight,
			})
		}
	}
	return g
}

type State struct {
	ID          int
	DistToStart int
}

type StateHeap []State

func (s *StateHeap) Len() int {
	return len(*s)
}

// Less 最小堆实现
func (s *StateHeap) Less(i, j int) bool {
	return (*s)[i].DistToStart < (*s)[j].DistToStart
}

func (s *StateHeap) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

func (s *StateHeap) Push(x any) {
	*s = append(*s, x.(State))
}

func (s *StateHeap) Pop() (x any) {
	x = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return
}

func dijkstra(graph *Graph, start int) []int {

	n := len(graph.Nodes)
	dist2Start := make([]int, n)
	for ix := range dist2Start {
		dist2Start[ix] = math.MaxInt64
	}
	dist2Start[start] = 0

	h := &StateHeap{}
	heap.Init(h)
	heap.Push(h, State{
		ID:          start,
		DistToStart: 0,
	})

	for h.Len() > 0 {

		state := h.Pop().(State)
		// 如果距离表里边存在更小的, 那么当前节点废弃
		if dist2Start[state.ID] < state.DistToStart {
			continue
		}

		dist2Start[state.ID] = state.DistToStart

		for _, node := range graph.Nodes[state.ID].Next {
			if dist2Start[state.ID]+node.Weight < dist2Start[node.ID] {
				heap.Push(h, State{
					ID:          node.ID,
					DistToStart: node.Weight,
				})
			}
		}

	}

	return dist2Start
}

func main() {

}
