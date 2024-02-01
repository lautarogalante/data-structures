package main

import "fmt"

type Node struct {
	next *Node
	data string
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (list *LinkedList) AddNode(value string) {
	node := &Node{
		next: nil,
		data: value,
	}

	if list.head == nil {
		list.head = node
		list.tail = node
	} else {
		list.tail.next = node
		list.tail = node
	}
}

func (list *LinkedList) SearchNode(value string) []*Node {
	var result []*Node
	current := list.head

	for current != nil {
		if current.data == value {
			result = append(result, current)
		}
		current = current.next
	}

	return result
}

func (list *LinkedList) RemoveNode(value string) {
	current := list.head
	var prev *Node

	for current != nil && current.data != value {
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

func (list *LinkedList) AddLast(value string) {
	newNode := &Node{next: nil, data: value}
	list.tail.next = newNode
}

func (list *LinkedList) AddFirst(value string) {
	newNode := &Node{next: list.head, data: value}
	list.head = newNode
}

func (list *LinkedList) ReverseList() {

	if list.IsEmpty() {
		fmt.Println("The list is empty")
	}

	var current = list.head
	var next *Node = nil
	var prev *Node = nil

	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}

	list.head = prev

}

func (list *LinkedList) GetLastNode() *Node {

	if list.IsEmpty() {
		fmt.Println("The list is empty")
		return nil
	}

	return list.tail
}

func (list *LinkedList) Size() int {
	current := list.head
	count := 0
	for current != nil {
		current = current.next
		count++
	}

	return count
}

func (list *LinkedList) IsEmpty() bool {
	return list.head == nil
}

func (list *LinkedList) RemoveAllNode() {
	list.head = nil
	list.tail = nil
}

func (list *LinkedList) GetFirstNode() *Node {

	if list.IsEmpty() {
		fmt.Println("The list is empty")
		return nil
	}

	return list.head
}

func (list *LinkedList) InsertAtPosition(position int, value string) {
	current := list.head
	var prev *Node
	count := 1
	newNode := &Node{data: value}

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
