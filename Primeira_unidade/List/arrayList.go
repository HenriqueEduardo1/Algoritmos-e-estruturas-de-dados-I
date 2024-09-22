package main

import (
	"errors"
	"fmt"
	"strings"
)

type List[T any] interface {
	show()
	append(value T)
	get(index int) (T, error)
	pop() T
	update(value T, index int)
	insert(value T, index int)
	remove(index int)
}

type arrayList[T any] struct {
	items      []T
	countItems int
}

func initArrayList[T any](size int) *arrayList[T] {
	return &arrayList[T]{
		items:      make([]T, size),
		countItems: 0,
	}
}

func (a *arrayList[T]) doubleArray() {
	newItems := make([]T, len(a.items)*2)

	for i := 0; i < len(a.items); i++ {
		newItems[i] = a.items[i]
	}

	a.items = newItems
}

func (a *arrayList[T]) show() {
	var sb strings.Builder

	sb.WriteString("(")

	for i := 0; i < a.countItems; i++ {
		sb.WriteString(fmt.Sprintf("%v", a.items[i]))

		if i < a.countItems-1 {
			sb.WriteString(", ")
		}
	}

	sb.WriteString(")")

	fmt.Println(sb.String())
}

func (a *arrayList[T]) append(value T) {
	if a.countItems == len(a.items) {
		a.doubleArray()
	}

	a.items[a.countItems] = value
	a.countItems++
}

func (a *arrayList[T]) get(index int) (T, error) {
	if index < 0 || index >= a.countItems {
		var defaultValue T
		return defaultValue, errors.New("Index out of range")
	} else if a.countItems == 0 {
		var defaultValue T
		return defaultValue, errors.New("Empty list")
	}

	return a.items[index], nil
}

func main() {

	list := initArrayList[int](5)

	list.append(1)
	list.append(2)
	list.append(3)
	list.append(4)
	list.append(5)
	list.append(6)
	list.append(7)
	list.append(8)

	list.show()

	fmt.Println(list.get(-1))
	fmt.Println(list.get(5))

}
