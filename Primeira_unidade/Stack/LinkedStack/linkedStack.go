package main

import (
	"errors"
	"fmt"
)

type Stack interface {
	Size() int
	Push(e int) error
	Pop() (int, error)
	Top() (int, error)
}

type LinkedStack struct {
	top      *Node
	inserted int
}

type Node struct {
	value int
	next  *Node
}

func (s *LinkedStack) Size() int {
	return s.inserted
}

func (s *LinkedStack) Push(e int) {
	newNode := &Node{value: e}

	if s.top == nil {
		s.top = newNode
	}

	newNode.next = s.top
	s.top = newNode
	s.inserted++
}

func (s *LinkedStack) Pop() (int, error) {
	if s.inserted == 0 {
		return -1, errors.New("A pilha está vazia")
	}

	e := s.top.value
	s.top = s.top.next
	s.inserted--
	return e, nil
}

func (s *LinkedStack) Top() (int, error) {
	if s.inserted == 0 {
		return -1, errors.New("A pilha está vazia")
	}
	return s.top.value, nil
}

func main() {
	var s LinkedStack

	val, err := s.Top()
	fmt.Println(val, err)

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)

	e, _ := s.Top()
	fmt.Println(e)

	for i := 0; i < 6; i++ {
		e, _ = s.Pop()
		fmt.Println(e)
	}
}
