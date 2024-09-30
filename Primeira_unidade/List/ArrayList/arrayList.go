package main

import (
	"fmt"
	"strings"
)

type List[T any] interface {
	Show()
	Append(value T)
	Size() int
	Get(index int) T
	Pop() T
	Update(value T, index int)
	Insert(value T, index int)
	Remove(index int)
}

type ArrayList[T any] struct {
	items    []T
	inserted int
}

func InitArrayList[T any](size int) *ArrayList[T] {
	return &ArrayList[T]{
		items:    make([]T, size),
		inserted: 0,
	}
}

func (a *ArrayList[T]) doubleArray() {
	newItems := make([]T, len(a.items)*2)

	for i := 0; i < len(a.items); i++ {
		newItems[i] = a.items[i]
	}

	a.items = newItems
}

func (a *ArrayList[T]) Show() {
	var sb strings.Builder

	sb.WriteString("(")

	for i := 0; i < a.inserted; i++ {
		sb.WriteString(fmt.Sprintf("%v", a.items[i]))

		if i < a.inserted-1 {
			sb.WriteString(", ")
		}
	}

	sb.WriteString(")")

	fmt.Println(sb.String())
}

func (a *ArrayList[T]) Append(value T) {
	if a.inserted == len(a.items) {
		a.doubleArray()
	}

	a.items[a.inserted] = value
	a.inserted++
}

func (a *ArrayList[T]) Size() int {
	return a.inserted
}

func (a *ArrayList[T]) Get(index int) T {
	if index < 0 || index >= a.inserted {
		fmt.Println("Index out of bounds")
	} else if a.inserted == 0 {
		fmt.Println("Empty list")
	} else {
		return a.items[index]
	}

	var defaultValue T
	return defaultValue
}

func (a *ArrayList[T]) Pop() T {
	if a.inserted == 0 {
		fmt.Println("Empty list")
		var defaultValue T
		return defaultValue
	}

	value := a.items[a.inserted-1]
	a.inserted--

	return value
}

func (a *ArrayList[T]) Update(value T, index int) {
	if index < 0 || index >= a.inserted {
		fmt.Println("Index out of bounds")
	} else {
		a.items[index] = value
	}
}

func (a *ArrayList[T]) Insert(value T, index int) {
	if index < 0 || index >= a.inserted {
		fmt.Println("Index out of bounds")
	} else {
		if a.inserted == len(a.items) {
			a.doubleArray()
		}

		for i := a.inserted; i > index; i-- {
			a.items[i] = a.items[i-1]
		}

		a.items[index] = value
		a.inserted++
	}
}

func (a *ArrayList[T]) Remove(index int) {
	if index < 0 || index >= a.inserted {
		fmt.Println("Index out of bounds")
	} else {
		for i := index; i < a.inserted-1; i++ {
			a.items[i] = a.items[i+1]
		}

		a.inserted--
	}
}

func main() {

	list := InitArrayList[int](5)

	list.Append(10)
	list.Append(20)
	list.Append(30)
	list.Update(20, 0)
	list.Append(40)
	list.Append(50)

	list.Show()

	list.Insert(15, 2)
	list.Insert(15, -1)
	list.Insert(15, 10)

	list.Show()

	list.Insert(5, 5)

	list.Show()

	list.Remove(2)

	list.Show()

	fmt.Println(list.Get(2))
	fmt.Println(list.Get(10))

	list.Pop()

	list.Show()
	fmt.Println(list.Size())

}
