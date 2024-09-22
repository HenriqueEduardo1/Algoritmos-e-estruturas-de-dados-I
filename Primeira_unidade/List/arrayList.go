package main

import (
	"fmt"
	"strings"
)

type List[T any] interface {
	show() []T
	append(value T)
	get(index int) T
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
	newSize := len(a.items) * 2
	newItems := make([]T, newSize)
	for i := 0; i < len(a.items); i++ {
		newItems[i] = a.items[i]
	}
	a.items = newItems
}

func (a *arrayList[T]) show() string {
	var sb strings.Builder
	sb.WriteString("(")
	for i := 0; i < a.countItems; i++ {
		sb.WriteString(fmt.Sprintf("%v", a.items[i]))
		if i < a.countItems-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString(")")
	return sb.String()
}

func (a *arrayList[T]) append(value T) {
	if a.countItems == len(a.items) {
		a.doubleArray()
	}
	a.items[a.countItems] = value
	a.countItems++
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

	fmt.Println(list.show())

}
