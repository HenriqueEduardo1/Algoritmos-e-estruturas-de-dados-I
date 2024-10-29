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

type ArrayStack struct {
	values   []int
	inserted int
}

func (s *ArrayStack) InitArrayStack(size int) error {
	if size > 0 {
		s.values = make([]int, size)
		return nil
	} else {
		return errors.New("O tamanho da pilha deve ser maior que 0")
	}
}

func (s *ArrayStack) Size() int {
	return s.inserted
}

func (s *ArrayStack) Push(e int) error {
	if s.inserted == len(s.values) {
		return errors.New("A pilha está cheia")
	}
	s.values[s.inserted] = e
	s.inserted++
	return nil
}

func (s *ArrayStack) Pop() (int, error) {
	if s.inserted == 0 {
		return -1, errors.New("A pilha está vazia")
	}
	e := s.values[s.inserted-1]
	s.inserted--
	return e, nil
}

func (s *ArrayStack) Top() (int, error) {
	if s.inserted == 0 {
		return -1, errors.New("A pilha está vazia")
	}
	return s.values[s.inserted-1], nil
}

func balparenteses(par string) bool {
	var s ArrayStack
	s.InitArrayStack(len(par))

	for i := 0; i < len(par); i++ {
		if par[i] == '(' {
			s.Push(1)
		} else if par[i] == ')' {
			if s.Size() == 0 {
				return false
			}
			s.Pop()
		}
	}

	if s.Size() == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	var s ArrayStack
	s.InitArrayStack(5)

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

	par := "(())"

	fmt.Println(balparenteses(par))

}
