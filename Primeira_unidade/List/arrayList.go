package main

import (
	"fmt"
	"strings"
)

type List[T any] interface {
	show()
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

func (a *arrayList[T]) get(index int) T {
	if index < 0 || index >= a.countItems {
		fmt.Println("Index out of range")
	} else if a.countItems == 0 {
		fmt.Println("Empty list")
	} else {
		return a.items[index]
	}

	var defaultValue T
	return defaultValue
}

func (a *arrayList[T]) pop() T {
	if a.countItems == 0 {
		fmt.Println("Empty list")
		var defaultValue T
		return defaultValue
	}

	value := a.items[a.countItems-1]
	a.countItems--

	return value
}

func (a *arrayList[T]) update(value T, index int) {
	if index < 0 || index >= a.countItems {
		fmt.Println("Index out of range")
	} else {
		a.items[index] = value
	}
}

func (a *arrayList[T]) insert(value T, index int) {
	if index < 0 || index >= a.countItems {
		fmt.Println("Index out of range")
	} else {
		if a.countItems == len(a.items) {
			a.doubleArray()
		}

		for i := a.countItems; i > index; i-- {
			a.items[i] = a.items[i-1]
		}

		a.items[index] = value
		a.countItems++
	}
}

func (a *arrayList[T]) remove(index int) {
	if index < 0 || index >= a.countItems {
		fmt.Println("Index out of range")
	} else {
		for i := index; i < a.countItems-1; i++ {
			a.items[i] = a.items[i+1]
		}

		a.countItems--
	}
}

func main() {

	list := initArrayList[int](5)

	list.append(10)
	list.append(20)
	list.append(30)
	list.update(20, 0)
	list.append(40)
	list.append(50)

	list.show()

	list.insert(15, 2)
	list.insert(15, -1)
	list.insert(15, 10)

	list.show()

	list.insert(5, 5)

	list.show()

	list.remove(2)

	list.show()

	fmt.Println(list.get(2))
	fmt.Println(list.get(10))

	list.pop()

	list.show()

}
