package main

import "fmt"

type queue struct {
	data []int
	head int
}

func (q *queue) Enqueue(value int) {
	q.data = append(q.data, value)
}

func (q *queue) Dequeue() (int, bool) {
	if q.head >= len(q.data) {
		return 0, false
	}

	value := q.data[q.head]

	q.head++

	if q.head == len(q.data) {
		q.data = nil
		q.head = 0
	}
	return value, true
}

func main() {
	newQueue := queue{data: make([]int, 0), head: 0}
	newQueue.Enqueue(10)
	newQueue.Enqueue(20)
	newQueue.Enqueue(30)
	fmt.Println("queue inicial: ", newQueue.data)
	value, _ := newQueue.Dequeue()
	fmt.Println("primer valor a eliminar: ", value)
	value, _ = newQueue.Dequeue()
	fmt.Println("segundo valor a eliminar: ", value)
	value, _ = newQueue.Dequeue()
	fmt.Println("tercer valor a elinminar: ", value)
	newQueue.Enqueue(45)
	fmt.Println("queue final: ", newQueue.data)
}
