package main

import (
	"fmt"
)

type MTypes interface {
	int | int8 | int16 | float32 | float64 | ~string
}

type Node[T MTypes] struct {
	next *Node[T]
	data T
}

type LinkedList[T MTypes] struct {
	head *Node[T]
	tail *Node[T]
}

func (list *LinkedList[T]) AddNode(v T) {
	node := &Node[T]{
		next: nil,
		data: v,
	}

	if list.head == nil {
		list.head = node
		list.tail = node
	} else {
		list.tail.next = node
		list.tail = node
	}
}

func (list *LinkedList[T]) SearchNode(v T) []*Node[T] {
	var result []*Node[T]
	current := list.head

	for current != nil {
		if current.data == v {
			result = append(result, current)
		}
		current = current.next
	}

	return result
}

func (list *LinkedList[T]) AddLast(v T) {
	newNode := &Node[T]{next: nil, data: v}
	list.tail.next = newNode
}

func (list *LinkedList[T]) AddFirst(v T) {
	newNode := &Node[T]{next: list.head, data: v}
	list.head = newNode
}

func (list *LinkedList[T]) RemoveNode(v T) {
	current := list.head
	var prev *Node[T]

	for current != nil && current.data != v {
		prev = current
		current = current.next
	}

	if current != nil {
		if prev != nil {
			prev.next = current.next
		} else {
			list.head = current.next
		}
		current = nil
	}

}

func (list *LinkedList[T]) RemoveAllNode() {
	list.head = nil
	list.tail = nil
}

func (list *LinkedList[T]) ReverseList() {

	if list.IsEmpty() {
		fmt.Println("The list is empty")
	}

	var current = list.head
	var next *Node[T] = nil
	var prev *Node[T] = nil

	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}

	list.tail = list.head
	list.head = prev
}

func (list *LinkedList[T]) GetLastNode() *Node[T] {

	if list.IsEmpty() {
		fmt.Println("The list is empty")
		return nil
	}

	return list.tail.next
}

func (list *LinkedList[T]) GetFirstNode() *Node[T] {

	if list.IsEmpty() {
		fmt.Println("The list is empty")
		return nil
	}

	return list.head
}

func (list *LinkedList[T]) Size() int {
	current := list.head
	count := 0
	for current != nil {
		current = current.next
		count++
	}

	return count
}

func (list *LinkedList[T]) IsEmpty() bool {
	return list.head == nil
}

func (list *LinkedList[T]) InsertAtPosition(position int, v T) {
	current := list.head
	var prev *Node[T]
	count := 1
	newNode := &Node[T]{data: v}

	if list.IsEmpty() {
		newNode.next = nil
		list.head = newNode
		return
	}

	for current != nil && count < position {
		prev = current
		current = current.next
		count++
	}

	if position == 1 {
		newNode.next = list.head
		list.head = newNode
	} else if count == position {
		newNode.next = current
		prev.next = newNode
		if current == nil {
			list.tail = newNode
		}
	}
}
