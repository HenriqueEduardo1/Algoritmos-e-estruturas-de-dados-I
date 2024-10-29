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

type ArrayQueue struct {
	v        []int
	front    int
	rear     int
	inserted int
}

func (q *ArrayQueue) InitArrayQueue(size int) error {
	if size > 0 {
		q.v = make([]int, size)
		return nil
	} else {
		return errors.New("O tamanho da fila deve ser maior que 0")
	}
}

func (q *ArrayQueue) Size() int {
	return q.inserted
}

func (q *ArrayQueue) Enqueue(e int) error {
	if q.rear == len(q.v) {
		return errors.New("A fila está cheia")
	}
	q.v[q.rear] = e
	q.rear++
	q.inserted++
	return nil
}

func (q *ArrayQueue) Dequeue() (int, error) {
	if q.inserted == 0 {
		return -1, errors.New("A fila está vazia")
	}
	e := q.v[q.front]
	q.front++
	q.inserted--
	return e, nil
}

func (q *ArrayQueue) Front() (int, error) {
	if q.inserted == 0 {
		return -1, errors.New("A fila está vazia")
	}
	return q.v[q.front], nil
}

func main() {
	q := &ArrayQueue{}
	q.InitArrayQueue(5)

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
