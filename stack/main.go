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
	count int
}

// NewStack returns a new stack.
func NewStack() *stack {
	return &stack{}
}

func (s *stack) isEmptyStack() bool {
	return s.count == 0
}

// push adds a node to the stack.
func (s *stack) push(n *node) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

// pop removes and returns a node from the stack in last to first order.
func (s *stack) pop() *node {
	if s.isEmptyStack() {
		return nil
	}
	s.count--
	return s.nodes[s.count]
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
		fmt.Println(s.nodes)
	}

	for i := 0; i < len(nums); i++ {
		n := s.pop()
		if n == nil {
			panic(err)
		}
		fmt.Println(s.nodes)
	}
}
