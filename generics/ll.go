package main

import "fmt"

type NodeLL[T comparable] struct {
	val  T
	next *NodeLL[T]
}

type LinkedList[T comparable] struct {
	root *NodeLL[T]
}

func (ll *LinkedList[T]) Add(val T) {
	if ll == nil {
		return
	}
	if ll.root == nil {
		ll.root = &NodeLL[T]{val: val}
		return
	}
	tmp := ll.root
	for tmp.next != nil {
		tmp = tmp.next
	}
	tmp.next = &NodeLL[T]{val: val}
}

func (ll *LinkedList[T]) Insert(val T, index int) {
	if ll == nil || ll.root == nil {
		return
	}

	if index == 0 {
		ll.root = &NodeLL[T]{val: val, next: ll.root}
		return
	}

	tmp := ll.root
	var tmp1 *NodeLL[T]
	for index > 0 && tmp != nil {
		tmp1 = tmp
		tmp = tmp.next
		index -= 1
	}

	if index != 0 {
		return
	}
	tmp1.next = &NodeLL[T]{val: val, next: tmp}
}

func (ll *LinkedList[T]) Index(val T) int {
	if ll == nil || ll.root == nil {
		return -1
	}
	tmp := ll.root
	index := 0
	for tmp != nil {
		if tmp.val == val {
			return index
		}
		tmp = tmp.next
		index += 1
	}

	return -1
}

func (ll *LinkedList[T]) Print() {
	if ll == nil {
		return
	}
	tmp := ll.root

	for tmp != nil {
		fmt.Println(tmp.val)
		tmp = tmp.next
	}
}

func LLmain() {
	ll := &LinkedList[int]{}
	ll.Add(5)
	ll.Add(6)
	ll.Add(7)
	ll.Add(9)
	ll.Insert(8, 3)
	ll.Print()
	fmt.Println(ll.Index(8), ll.Index(-8))
}
