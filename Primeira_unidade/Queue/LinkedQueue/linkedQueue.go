package main

import (
	"errors"
	"fmt"
)

type Queue interface {
	Size() int
	Enqueue(e int) error
	Dequeue() (int, error)
	Front() (int, error)
}

type Node struct {
	val  int
	next *Node
}

type LinkedQueue struct {
	front    *Node
	rear     *Node
	inserted int
}

func (q *LinkedQueue) Size() int {
	return q.inserted
}

func (q *LinkedQueue) Enqueue(e int) error {
	newNode := &Node{val: e}
	if q.rear == nil {
		q.front = newNode
		q.rear = newNode
	} else {
		q.rear.next = newNode
		q.rear = newNode
	}
	q.inserted++
	return nil
}

func (q *LinkedQueue) Dequeue() (int, error) {
	if q.inserted == 0 {
		return -1, errors.New("A fila está vazia")
	} else if q.inserted == 1 {
		e := q.front.val
		q.front = nil
		q.rear = nil
		q.inserted--
		return e, nil
	} else {
		e := q.front.val
		q.front = q.front.next
		q.inserted--
		return e, nil
	}
}

func (q *LinkedQueue) Front() (int, error) {
	if q.inserted == 0 {
		return -1, errors.New("A fila está vazia")
	}
	return q.front.val, nil
}

func main() {
	q := &LinkedQueue{}

	_, err := q.Dequeue()
	fmt.Println(err)

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)

	println(q.Size())

	e, _ := q.Front()
	fmt.Println(e)

	for i := 0; i < 6; i++ {
		e, _ := q.Dequeue()
		fmt.Println(e)
	}

}
