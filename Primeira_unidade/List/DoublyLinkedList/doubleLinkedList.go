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
	prev  *Node[T]
	next  *Node[T]
}

type DoublyLinkedList[T any] struct {
	head     *Node[T]
	tail     *Node[T]
	inserted int
}

func (d *DoublyLinkedList[T]) Show() {
	var sb strings.Builder

	sb.WriteString("(")
	current := d.head

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

func (d *DoublyLinkedList[T]) Size() int {
	return d.inserted
}

func (d *DoublyLinkedList[T]) Get(index int) T {
	if index < 0 || index >= d.inserted {
		fmt.Println("Index out of bounds")
	} else if d.head == nil {
		fmt.Println("Empty list")
	} else {
		current := d.head

		for i := 0; i < index; i++ {
			current = current.next
		}

		return current.value
	}

	var defaultValue T
	return defaultValue
}

func (d *DoublyLinkedList[T]) Add(value T) {
	newNode := &Node[T]{value: value}

	if d.head == nil {
		d.head = newNode
	} else {
		d.tail.next = newNode
		newNode.prev = d.tail
	}

	d.tail = newNode
	d.inserted++
}

func (d *DoublyLinkedList[T]) AddOnIndex(value T, index int) {
	newNode := &Node[T]{value: value}

	if index >= 0 && index < d.inserted {
		if index == 0 {
			if d.head == nil {
				d.head = newNode
				d.tail = newNode
			} else {
				newNode.next = d.head
				d.head.prev = newNode
				d.head = newNode
			}
		} else {
			current := d.head

			for i := 0; i < index-1; i++ {
				current = current.next
			}

			newNode.next = current.next
			newNode.prev = current
			current.next = newNode

			if newNode.next == nil {
				d.tail = newNode
			} else {
				newNode.next.prev = newNode
			}
		}
	}

	fmt.Println("Index out of bounds")
}

func (d *DoublyLinkedList[T]) Remove(index int) {
	if index < 0 || index >= d.inserted {
		fmt.Println("Index out of bounds")
	} else if d.head == nil {
		fmt.Println("Empty list")
	} else if index == 0 {
		d.head = d.head.next
		d.inserted--
	} else {
		current := d.head

		for i := 0; i < index-1; i++ {
			current = current.next
		}

		current.next = current.next.next
		d.inserted--
	}
}

func main() {

	doublyLinkedList := DoublyLinkedList[int]{}

	doublyLinkedList.Add(1)
	doublyLinkedList.Add(2)
	doublyLinkedList.Add(3)
	doublyLinkedList.Add(4)

	doublyLinkedList.Show()

	doublyLinkedList.AddOnIndex(5, 3)

	doublyLinkedList.Show()

}
