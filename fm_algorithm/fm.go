package main

import (
	"fmt"
	"math"
)

type Node struct {
	id        string
	gain      int
	isSwapped bool
	partition int // 0: left, 1: right
}

type Graph struct {
	nodes map[string]*Node   // node id -> node
	to    map[string][]*Node // node id -> node list
}

func newGraph(input map[string][]string) *Graph {
	graph := &Graph{
		nodes: make(map[string]*Node),
		to:    make(map[string][]*Node),
	}
	for fromId, toList := range input {
		if _, ok := graph.nodes[fromId]; !ok {
			graph.nodes[fromId] = &Node{id: fromId}
		}
		graph.to[fromId] = make([]*Node, 0)
		for _, toId := range toList {
			if _, ok := graph.nodes[toId]; !ok {
				graph.nodes[toId] = &Node{id: toId}
			}
			graph.to[fromId] = append(graph.to[fromId], graph.nodes[toId])
		}
	}
	return graph
}

func (g *Graph) Print() {
	for fromId, toList := range g.to {
		fmt.Printf("%s: ", fromId)
		fmt.Printf("[")
		for i, node := range toList {
			fmt.Printf("%s", node.id)
			if i != len(toList)-1 {
				fmt.Printf(", ")
			}
		}
		fmt.Printf("]\n")
	}
}

type HyperGraph struct {
	nodes     map[string]*Node // node id -> node
	num       int              // index num
	indexs    map[string][]int // node id -> index
	hyperEdge map[int][]*Node  // index -> node list
}

func newHyperGraph(g *Graph) *HyperGraph {
	hpg := &HyperGraph{
		nodes:     g.nodes,
		num:       0,
		indexs:    make(map[string][]int),
		hyperEdge: make(map[int][]*Node),
	}

	for fromId, toList := range g.to {
		if len(toList) == 0 { // not only itself
			continue

		}
		hpg.indexs[fromId] = append(hpg.indexs[fromId], hpg.num)
		hpg.hyperEdge[hpg.num] = append(hpg.hyperEdge[hpg.num], hpg.nodes[fromId])

		for _, to := range toList {
			hpg.indexs[to.id] = append(hpg.indexs[to.id], hpg.num)
			hpg.hyperEdge[hpg.num] = append(hpg.hyperEdge[hpg.num], to)
		}
		hpg.num++
	}

	return hpg
}

func (h *HyperGraph) randomPartition() (*Partition, *Partition) {
	left := &Partition{
		maxGain: math.MinInt,
		nodes:   make(map[string]*Node),
	}
	right := &Partition{
		maxGain: math.MinInt,
		nodes:   make(map[string]*Node),
	}
	//cnt := 0
	//for _, node := range hpg.nodes {
	//	if cnt%2 == 0 {
	//		hpg.nodes[node.id].partition = 0
	//		left.nodes[node.id] = node
	//	} else {
	//		hpg.nodes[node.id].partition = 1
	//		right.nodes[node.id] = node
	//	}
	//	cnt++
	//}
	l := []string{"a", "c", "d", "g"}
	r := []string{"b", "e", "f", "h"}
	for _, id := range l {
		h.nodes[id].partition = 0
		left.nodes[id] = h.nodes[id]
		left.remain++
	}
	for _, id := range r {
		h.nodes[id].partition = 1
		right.nodes[id] = h.nodes[id]
		right.remain++
	}
	return left, right
}

func (h *HyperGraph) GetHyperEdgeById(nodeId string) map[int][]*Node {
	ret := make(map[int][]*Node)
	for _, idx := range h.indexs[nodeId] {
		ret[idx] = h.hyperEdge[idx]
	}
	return ret
}

func (h *HyperGraph) Print() {
	for _, nodes := range h.hyperEdge {
		fmt.Printf("[")
		for i, node := range nodes {
			fmt.Printf("%s", node.id)
			if i != len(nodes)-1 {
				fmt.Printf(", ")
			}
		}
		fmt.Printf("]\n")
	}
}

type Partition struct {
	maxGainId string
	maxGain   int
	nodes     map[string]*Node // node id -> node
	remain    int              // the number of nodes remaining
}

func (p *Partition) toString() []string {
	ret := make([]string, 0, len(p.nodes))
	for id, _ := range p.nodes {
		ret = append(ret, id)
	}
	return ret
}

type Record struct {
	cell    string
	gain    int
	sumGain int
	cutsize int
	left    []string
	right   []string
}

func newRecord(cell string, gain, sumGain, cutsize int, left, right *Partition) *Record {
	return &Record{
		cell:    cell,
		gain:    gain,
		sumGain: sumGain,
		cutsize: cutsize,
		left:    left.toString(),
		right:   right.toString(),
	}
}

func (r *Record) Print(num int) {
	fmt.Printf("#: %-3d cell: %-3s gain: %-3d sum: %-3d cutsize: %-3d  ", num, r.cell, r.gain, r.sumGain, r.cutsize)
	fmt.Printf("left: [")
	for i, id := range r.left {
		fmt.Printf("%s", id)
		if i != len(r.left)-1 {
			fmt.Printf(", ")
		}
		i++
	}
	fmt.Printf("]  ")

	fmt.Printf("right: [")
	for i, id := range r.right {
		fmt.Printf("%s", id)
		if i != len(r.right)-1 {
			fmt.Printf(", ")
		}
		i++
	}
	fmt.Printf("]\n")
}

type FM struct {
	g                   *Graph
	HPG                 *HyperGraph
	left                *Partition
	right               *Partition
	records             []*Record
	cutsize0            int // init cutsize
	minCutsize          int
	minCutsizeRecordIdx []int
}

func NewFM(input map[string][]string) *FM {
	graph := newGraph(input)
	hpGraph := newHyperGraph(graph)
	left, right := hpGraph.randomPartition()
	fm := &FM{
		g:                   graph,
		HPG:                 hpGraph,
		left:                left,
		right:               right,
		records:             make([]*Record, 0),
		minCutsizeRecordIdx: make([]int, 0),
	}
	fm.cutsize0 = fm.computeCutsize()
	fm.minCutsize = fm.cutsize0
	fm.initGain()
	fm.records = append(fm.records, newRecord("-", 0, 0, fm.cutsize0, left, right))
	fm.minCutsizeRecordIdx = append(fm.minCutsizeRecordIdx, len(fm.records)-1)
	return fm
}

func (f *FM) initGain() {
	for _, curNode := range f.HPG.nodes {
		fs, te := 0, 0
		for _, nodes := range f.HPG.GetHyperEdgeById(curNode.id) {
			samePar := 1
			for _, node := range nodes {
				if curNode.id == node.id {
					continue
				}
				if curNode.partition == node.partition {
					samePar++
				}
			}
			if samePar == 1 {
				fs++
			}
			if samePar == len(nodes) {
				te++
			}
		}
		curNode.gain = fs + te
		if curNode.partition == 0 { // left partition
			if curNode.gain > f.left.maxGain {
				f.left.maxGain = curNode.gain
				f.left.maxGainId = curNode.id
			}
		}
		if curNode.partition == 1 { // right partition
			if curNode.gain > f.right.maxGain {
				f.right.maxGain = curNode.gain
				f.right.maxGainId = curNode.id
			}
		}
	}
}

func (f *FM) computeGain(affected map[string]*Node) {
	for _, curNode := range affected { // only gains are calculated for nodes affected by the swap operation
		fs, te := 0, 0
		for _, nodes := range f.HPG.GetHyperEdgeById(curNode.id) {
			samePar := 1
			for _, node := range nodes {
				if curNode.id == node.id {
					continue
				}
				if curNode.partition == node.partition {
					samePar++
				}
			}
			if samePar == 1 {
				fs++
			}
			if samePar == len(nodes) {
				te++
			}
		}
		curNode.gain = fs - te // Not the same as `initGain`
		if curNode.partition == 0 {
			if curNode.gain > f.left.maxGain {
				f.left.maxGain = curNode.gain
				f.left.maxGainId = curNode.id
			}
		}
		if curNode.partition == 1 {
			if curNode.gain > f.right.maxGain {
				f.right.maxGain = curNode.gain
				f.right.maxGainId = curNode.id
			}
		}
	}
}

func (f *FM) computeCutsize() int {
	cutsize := 0
	var partition int
	for _, link := range f.HPG.hyperEdge {
		if len(link) == 1 {
			continue
		}

		for i, node := range link {
			if i == 0 {
				partition = node.partition
			} else {
				if partition != node.partition {
					cutsize++
					break
				}
			}
		}
	}
	return cutsize
}

func (f *FM) Run() {
	for f.left.remain != 0 || f.right.remain != 0 {
		// select and swap node of max gain
		// compute gains of affected nodes
		swapped, affecteds := f.selectAndSwap()

		// recompute gains of the affected nodes
		f.computeGain(affecteds)

		// structure and record
		record := f.structureRecord(swapped)
		f.records = append(f.records, record)
		if record.cutsize < f.minCutsize {
			f.minCutsize = record.cutsize
			f.minCutsizeRecordIdx = make([]int, 0)
			f.minCutsizeRecordIdx = append(f.minCutsizeRecordIdx, len(f.records)-1)
		} else if record.cutsize == f.minCutsize {
			f.minCutsizeRecordIdx = append(f.minCutsizeRecordIdx, len(f.records)-1)
		}
	}
}

func (f *FM) selectAndSwap() (*Node, map[string]*Node) {
	// select and swap
	var swapped *Node
	if f.right.remain >= f.left.remain { // right -> left
		swapped = f.right.nodes[f.right.maxGainId]
		f.left.nodes[swapped.id] = swapped
		delete(f.right.nodes, swapped.id)
		f.right.remain--
	} else { // left -> right
		swapped = f.left.nodes[f.left.maxGainId]
		f.right.nodes[swapped.id] = swapped
		delete(f.left.nodes, swapped.id)
		f.left.remain--
	}
	swapped.partition ^= 1
	swapped.isSwapped = true

	// search affected node by the swap operation
	affecteds := make(map[string]*Node)
	for _, link := range f.HPG.GetHyperEdgeById(swapped.id) {
		for _, afn := range link {
			if afn.id != swapped.id && !afn.isSwapped {
				affecteds[afn.id] = afn
			}
		}
	}

	// recompute max gain of each partition
	f.left.maxGain = math.MinInt
	for _, n := range f.left.nodes {
		if _, ok := affecteds[n.id]; ok {
			continue
		}
		if !n.isSwapped && n.gain > f.left.maxGain {
			f.left.maxGain = n.gain
			f.left.maxGainId = n.id
		}
	}
	f.right.maxGain = math.MinInt
	for _, n := range f.right.nodes {
		if _, ok := affecteds[n.id]; ok {
			continue
		}
		if !n.isSwapped && n.gain > f.right.maxGain {
			f.right.maxGain = n.gain
			f.right.maxGainId = n.id
		}
	}

	return swapped, affecteds
}

func (f *FM) structureRecord(swapped *Node) *Record {
	return &Record{
		cell:    swapped.id,
		gain:    swapped.gain,
		sumGain: swapped.gain + f.records[len(f.records)-1].sumGain,
		cutsize: f.computeCutsize(),
		left:    f.left.toString(),
		right:   f.right.toString(),
	}
}

func (f *FM) PrintRecord() {
	fmt.Println("Partition Records:")
	for i, record := range f.records {
		record.Print(i)
	}
	fmt.Println()
	fmt.Println("min cutsize: ", f.minCutsize)
	for _, idx := range f.minCutsizeRecordIdx {
		f.records[idx].Print(idx)
	}
}
