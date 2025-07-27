package main

import (
	"errors"
	"fmt"
)

type Node[T any] struct {
	val         T
	left, right *Node[T]
}

type Tree[T any] struct {
	OrderedFunc func(T, T) int
	root        *Node[T]
}

func (t *Tree[T]) Add(v T) error {
	if t == nil {
		return errors.New("tree not initialised")
	}

	if t.root == nil {
		t.root = &Node[T]{val: v}
		return nil
	}

	tmp := t.root
	var tmp1 *Node[T]
	for tmp != nil {
		tmp1 = tmp
		if t.OrderedFunc(tmp.val, v) <= 0 {
			tmp = tmp.right
			continue
		}
		tmp = tmp.left
	}

	if t.OrderedFunc(tmp1.val, v) <= 0 {
		tmp1.right = &Node[T]{val: v}
		return nil
	}
	tmp1.left = &Node[T]{val: v}
	return nil
}

func (t *Tree[T]) Contains(val T) bool {
	if t == nil || t.root == nil {
		return false
	}
	tmp := t.root
	var val1 int
	for tmp != nil {
		val1 = t.OrderedFunc(tmp.val, val)
		if val1 == 0 {
			return true
		} else if val1 < 0 {
			tmp = tmp.right
		} else {
			tmp = tmp.left
		}
	}
	return false
}

func (t *Tree[T]) Print() {
	if t == nil || t.root == nil {
		return
	}
	stack := []*Node[T]{t.root}
	for len(stack) > 0 {
		top := stack[0]
		stack = stack[1:]
		if top == nil {
			continue
		}
		stack = append(stack, top.left, top.right)
		fmt.Println(top.val)
	}
}

func main() {
	t := &Tree[int]{OrderedFunc: func(v1, v2 int) int { return v1 - v2 }}
	t.Add(4)
	t.Add(5)
	t.Add(2)
	t.Add(6)
	t.Add(1)
	t.Print()
	fmt.Println(t.Contains(7), t.Contains(-1), t.Contains(5))
}
