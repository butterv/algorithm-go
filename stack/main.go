package main

import (
	"fmt"

	"github.com/istsh/algorithm-go/util"
)

type node struct {
	value int
}

type stack struct {
	nodes []*node
}

// NewStack returns a new stack.
func NewStack() *stack {
	return &stack{}
}

func (s *stack) output() {
	for _, node := range s.nodes {
		fmt.Printf("%d ", node.value)
	}
	fmt.Println()
}

func (s *stack) isEmptyStack() bool {
	return len(s.nodes) == 0
}

// push adds a node to the stack.
func (s *stack) push(n *node) {
	s.nodes = append(s.nodes, n)
}

// pop removes and returns a node from the stack in last to first order.
func (s *stack) pop() *node {
	if s.isEmptyStack() {
		return nil
	}
	p := s.nodes[len(s.nodes)-1]
	tmp := make([]*node, len(s.nodes)-1)
	for i, node := range s.nodes[:len(s.nodes)-1] {
		tmp[i] = node
	}
	s.nodes = tmp
	return p
}

// echo 1 2 3 4 5 6 7 8 9 10 | go run ./main.go
func main() {
	nums, err := util.ReadIntArray(" ")
	if err != nil {
		panic(err)
	}

	s := NewStack()
	for _, num := range nums {
		s.push(&node{value: num})
		s.output()
	}

	for i := 0; i < len(nums); i++ {
		n := s.pop()
		if n == nil {
			panic(err)
		}
		s.output()
	}
}
