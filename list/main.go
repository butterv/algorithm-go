package main

type node struct {
	value int
	prev  *int
	next  *int
}

type list struct {
	nodes []*node
}

func NewList() *list {
	return &list{}
}

func (l *list) isEmptyList() bool {
	return len(l.nodes) == 0
}

func (l *list) search(value int) *node {
	for _, v := range l.nodes {
		if v.value == value {
			return v
		}
	}
	return nil
}

func (l *list) insert(n *node) {
	if l.isEmptyList() {
		next := 1
		n.next = &next
		l.nodes = append(l.nodes, n)
		return
	}

	last := l.nodes[len(l.nodes)-1]
	next := *last.next + 1
	n.prev = last.next
	n.next = &next
	l.nodes = append(l.nodes, n)
	return
}

func (l *list) delete(n *node) {
	if l.isEmptyList() {
		return
	}

	l.search()

	tmp := make([]*node, len(l.nodes)-1)
	for _, v := range l.nodes {

	}
}
