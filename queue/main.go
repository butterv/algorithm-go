package main

import (
	"fmt"

	"github.com/istsh/algorithm-go/util"
)

type node struct {
	value int
}

type queue struct {
	nodes []*node
}

// NewQueue returns a new queue.
func NewQueue() *queue {
	return &queue{}
}

func (q *queue) output() {
	for _, node := range q.nodes {
		fmt.Printf("%d ", node.value)
	}
	fmt.Printf("len: %d\n", len(q.nodes))
}

func (q *queue) isEmptyQueue() bool {
	return len(q.nodes) == 0
}

// enqueue adds a node to the queue.
func (q *queue) enqueue(n *node) {
	q.nodes = append(q.nodes, n)
}

// dequeue removes and returns a node from the queue.
func (q *queue) dequeue() *node {
	if q.isEmptyQueue() {
		return nil
	}
	p := q.nodes[0]
	tmp := make([]*node, len(q.nodes)-1)
	for i, node := range q.nodes[1:] {
		tmp[i] = node
	}
	q.nodes = tmp
	return p
}

// echo 1 2 3 4 5 6 7 8 9 10 | go run ./main.go
func main() {
	nums, err := util.ReadIntArray(" ")
	if err != nil {
		panic(err)
	}

	q := NewQueue()
	for _, num := range nums {
		q.enqueue(&node{value: num})
		q.output()
	}

	for i := 0; i < len(nums); i++ {
		n := q.dequeue()
		if n == nil {
			panic(err)
		}
		q.output()
	}
}
