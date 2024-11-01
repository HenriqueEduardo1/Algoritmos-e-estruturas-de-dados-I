package main

import (
	"fmt"
	"strings"
)

type List[T any] interface {
	Size() int
	Get(index int) T
	Add(value T)
	AddOnIndex(value T, index int)
	Remove(index int)
	Show()
}

type Node[T any] struct {
	value T
	next  *Node[T]
}

type LinkedList[T any] struct {
	head     *Node[T]
	inserted int
}

func (a *LinkedList[T]) Size() int {
	return a.inserted
}

func (a *LinkedList[T]) Show() {
	var sb strings.Builder

	sb.WriteString("(")
	current := a.head

	for current != nil {
		sb.WriteString(fmt.Sprintf("%v", current.value))

		if current.next != nil {
			sb.WriteString(", ")
		}

		current = current.next
	}

	sb.WriteString(")")

	fmt.Println(sb.String())
}

func (a *LinkedList[T]) Add(value T) {
	newNode := &Node[T]{value: value}

	if a.head == nil {
		a.head = newNode
	} else {
		current := a.head

		for current.next != nil {
			current = current.next
		}

		current.next = newNode
		a.inserted++
	}
}

func (a *LinkedList[T]) Get(index int) T {

	if index < 0 || index >= a.inserted {
		fmt.Println("Index out of bounds")
	} else if a.head == nil {
		fmt.Println("Empty list")
	} else {
		current := a.head

		for i := 0; i < index; i++ {
			current = current.next
		}

		return current.value
	}

	var defaultValue T
	return defaultValue
}

func (a *LinkedList[T]) AddOnIndex(value T, index int) {
	newNode := &Node[T]{value: value}

	if index < 0 || index > a.inserted {
		fmt.Println("Index out of bounds")
	} else if index == 0 {
		newNode.next = a.head
		a.head = newNode
		a.inserted++
	} else {
		current := a.head

		for i := 0; i < index-1; i++ {
			current = current.next
		}

		newNode.next = current.next
		current.next = newNode
		a.inserted++
	}
}

func (a *LinkedList[T]) Remove(index int) {
	if index < 0 || index > a.inserted {
		fmt.Println("Index out of bounds")
	} else if a.head == nil {
		fmt.Println("Empty list")
	} else if index == 0 {
		a.head = a.head.next
		a.inserted--
	} else {
		current := a.head

		for i := 0; i < index-1; i++ {
			current = current.next
		}

		current.next = current.next.next
		a.inserted--
	}
}

func (list *LinkedList[T]) Reverse() {
	if list.head == nil {
		fmt.Println("Empty list")
	} else {
		var previous *Node[T]
		current := list.head
		var next *Node[T]

		for current != nil {
			next = current.next
			current.next = previous
			previous = current
			current = next
		}

		list.head = previous
	}
}

func main() {
	list := LinkedList[int]{head: nil, inserted: 0}

	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)

	list.Show()
	list.Reverse()
	list.Show()
}
