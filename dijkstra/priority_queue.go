package main

type entry struct {
	iterm    interface{}
	priority float32
}

type PriorityQueue []*entry

func (q PriorityQueue) Len() int {
	return len(q)
}

func (q PriorityQueue) Less(i, j int) bool {
	return q[i].priority < q[j].priority
}

func (q PriorityQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *PriorityQueue) Push(x interface{}) {
	*q = append(*q, x.(*entry))
}

func (q *PriorityQueue) Pop() interface{} {
	n := q.Len()
	e := (*q)[n-1]
	(*q)[n-1] = nil
	*q = (*q)[:n-1]
	return e
}

func (q *PriorityQueue) Empty() bool {
	return q.Len() == 0
}

func NewPriorityQueue(length int) *PriorityQueue {
	q := make(PriorityQueue, 0, length)
	return &q
}

func NewEntry(iterm interface{}, priority float32) *entry {
	return &entry{
		iterm:    iterm,
		priority: priority,
	}
}

func ExtractItermByEntry(e interface{}) interface{} {
	return e.(*entry).iterm
}
