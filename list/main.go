package main

import (
	"fmt"

	"github.com/istsh/algorithm-go/util"
)

type cell struct {
	item int
	next *cell
}

type list struct {
	top *cell
}

func newCell(x int, cp *cell) *cell {
	return &cell{
		item: x,
		next: cp,
	}
}

func newList() *list {
	return &list{
		top: newCell(0, nil),
	}
}

func (l *list) output() {
	i := -1
	cp := l.top.next
	for {
		if i < 0 {
			i++
			continue
		}
		fmt.Printf("%d ", cp.item)
		if cp.next == nil {
			break
		}
		cp = cp.next
	}
	fmt.Println()
}

//func (cp *cell) nthCell(n int) *cell {
//	// topは先頭を意味するダミーなので、回避するために-1から始める
//	i := -1
//	result := cp
//	for result != nil {
//		if i == n {
//			return result
//		}
//		i++
//		result = result.next
//	}
//	return nil
//}
//
//func (l *list) nth(n int) (int, bool) {
//	cp := l.top.nthCell(n)
//	if cp == nil {
//		return 0, false
//	} else {
//		return cp.item, true
//	}
//}

func (l *list) isEmpty() bool {
	return l.top == nil || l.top.next == nil
}

func (l *list) add(x int) bool {
	if l.top == nil {
		l.top = newCell(0, nil)
	}

	// リストの最後を取得
	prev := l.top
	for prev.next != nil {
		prev = prev.next
	}
	prev.next = newCell(x, nil)
	return true
}

func (l *list) delete(n int) bool {
	if l.top == nil {
		l.top = &cell{}
		return false
	}

	i := -1
	var target *cell
	prev := l.top
	for {
		fmt.Printf("i:%d, n:%d\n", i, n)
		if i == n {
			target = prev.next
			break
		}
		if prev.next == nil {
			break
		}
		prev = prev.next
		i++
	}

	if target == nil {
		prev = nil
	} else {
		prev.next = target.next
	}

	return true
}

func main() {
	nums, err := util.ReadIntArray(" ")
	if err != nil {
		panic(err)
	}

	l := newList()
	for _, num := range nums {
		l.add(num)
		l.output()
	}

	for i := len(nums) - 1; i >= 0; i-- {
		isDeleted := l.delete(i)
		l.output()
		fmt.Println(isDeleted)
	}
}
