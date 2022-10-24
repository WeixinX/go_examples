package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Edge struct {
	from string
	to   string
	cost float32
}

var (
	src   = "1"
	nodes = map[string]string{
		"1": "1",
		"2": "2",
		"3": "3",
		"4": "4",
		"5": "5",
		"6": "6",
	}
	edges = map[string]map[string]*Edge{
		"1": {
			"2": &Edge{cost: 7},
			"3": &Edge{cost: 9},
			"6": &Edge{cost: 14},
		},
		"2": {
			"1": &Edge{cost: 7},
			"3": &Edge{cost: 10},
			"4": &Edge{cost: 15},
		},
		"3": {
			"1": &Edge{cost: 9},
			"2": &Edge{cost: 10},
			"4": &Edge{cost: 11},
			"6": &Edge{cost: 2},
		},
		"4": {
			"2": &Edge{cost: 15},
			"3": &Edge{cost: 11},
			"5": &Edge{cost: 6},
		},
		"5": {
			"4": &Edge{cost: 6},
			"6": &Edge{cost: 9},
		},
		"6": {
			"1": &Edge{cost: 14},
			"3": &Edge{cost: 2},
			"5": &Edge{cost: 9},
		},
	}
)

const PrevNULL = "null"

// Dijkstra
// TODO 可以使用堆(优先级队列)进行优化
func Dijkstra(src string, nodes map[string]string, edges map[string]map[string]*Edge) (float32, map[string][]string) {
	var maxCost float32 = math.MaxFloat32 / 2
	visit := make(map[string]bool)
	cost := make(map[string]float32)
	path := make(map[string]string)

	for _, nid := range nodes {
		cost[nid] = maxCost
		visit[nid] = false
	}
	cost[src] = 0
	path[src] = PrevNULL

	for i := 0; i < len(nodes); i++ {
		minCost := maxCost
		var mid string
		for _, next := range nodes {
			if !visit[next] && cost[next] < minCost {
				minCost = cost[next]
				mid = next
			}
		}
		visit[mid] = true
		for _, next := range nodes {
			if visit[next] {
				continue
			}
			if _, ok := edges[mid][next]; !ok {
				continue
			}
			if cost[next] > cost[mid]+edges[mid][next].cost {
				cost[next] = cost[mid] + edges[mid][next].cost
				path[next] = mid
			}
		}
	}

	var retCost float32 = 0
	for _, c := range cost {
		if c != maxCost {
			retCost += c
		}
		//fmt.Printf("%s -> %s: %.2f\n", src, dest, c)
	}
	//for next, prev := range path {
	//	fmt.Println(next, prev)
	//}
	return retCost, getPath(src, nodes, path)
}

func DijkstraByPriority(src string, nodes map[string]string, edges map[string]map[string]*Edge) (float32, map[string][]string) {
	var maxCost float32 = math.MaxFloat32 / 2
	visit := make(map[string]bool)
	cost := make(map[string]float32)
	path := make(map[string]string)

	for _, nid := range nodes {
		cost[nid] = maxCost
		visit[nid] = false
	}
	cost[src] = 0
	path[src] = PrevNULL

	q := newPriorityQueue(len(nodes))
	q.push(&Dist{src, 0})
	for !q.empty() {
		mid := q.pop().nodeId
		visit[mid] = true
		for next, edge := range edges[mid] {
			if !visit[next] && cost[next] > cost[mid]+edge.cost {
				cost[next] = cost[mid] + edge.cost
				path[next] = mid
				q.push(&Dist{next, cost[next]})
			}
		}
	}
	var retCost float32 = 0
	for _, c := range cost {
		if c != maxCost {
			retCost += c
		}
		//fmt.Printf("%s -> %s: %.2f\n", src, dest, c)
	}
	//for next, prev := range path {
	//	fmt.Println(next, prev)
	//}
	return retCost, getPath(src, nodes, path)
}

type Dist struct {
	nodeId string
	cost   float32
}

type priorityQueue struct {
	q *PriorityQueue
}

func newPriorityQueue(length int) *priorityQueue {
	q := NewPriorityQueue(length)
	heap.Init(q)
	return &priorityQueue{q: q}
}

func (q *priorityQueue) push(d *Dist) {
	heap.Push(q.q, NewEntry(d, d.cost))
}

func (q *priorityQueue) pop() *Dist {
	return ExtractItermByEntry(heap.Pop(q.q)).(*Dist)
}

func (q *priorityQueue) empty() bool {
	return q.q.Empty()
}

func getPath(src string, nodes map[string]string, row map[string]string) map[string][]string {
	s := NewStack()
	ret := make(map[string][]string)
	for _, nid := range nodes {
		s.Push(nid)
		i := nid
		for row[i] != PrevNULL {
			s.Push(row[i])
			i = row[i]
		}
		if _, ok := ret[nid]; !ok {
			ret[nid] = make([]string, 0, s.Size())
		}
		for !s.Empty() {
			id := s.Pop().(string)
			ret[nid] = append(ret[nid], id)
		}
	}
	//for dest, path := range ret {
	//	fmt.Printf("%s -> %s: ", src, dest)
	//	for _, p := range path {
	//		fmt.Printf("%s ", p)
	//	}
	//	fmt.Println()
	//}
	return ret
}

func main() {
	cost, paths := Dijkstra(src, nodes, edges)
	fmt.Println("total cost: ", cost)
	for dest, path := range paths {
		fmt.Printf("%s -> %s: ", src, dest)
		for _, p := range path {
			fmt.Printf("%s ", p)
		}
		fmt.Println()
	}

	fmt.Println()

	cost, paths = DijkstraByPriority(src, nodes, edges)
	fmt.Println("total cost: ", cost)
	for dest, path := range paths {
		fmt.Printf("%s -> %s: ", src, dest)
		for _, p := range path {
			fmt.Printf("%s ", p)
		}
		fmt.Println()
	}
}
