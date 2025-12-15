package main

import "fmt"

type Queue struct {
	value []int
	next  *Queue
}

type Head struct {
	value int
	next  *Head
}

func main() {
	q := &Queue{}
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	q.Dequeue()
	q.Dequeue()
	fmt.Println(q)
	q.Front()
	q.IsEmpty()
	fmt.Println(q)
}

func (q *Queue) Enqueue(value int) {
	q.value = append(q.value, value)
}

func (q *Queue) Dequeue() (int, bool) {
	if q.value == nil {
		return 0, false
	}

	valueToRemove := q.value[0]
	q.value = q.value[1:]
	return valueToRemove, true
}

func (q *Queue) Front() (int, bool) {
	if q.value == nil {
		return 0, false
	}
	fmt.Println(q.value)
	return q.value[0], true
}

func (q *Queue) IsEmpty() bool {
	if q.value == nil {
		fmt.Println("Queue is empty")
		return true
	}
	fmt.Println("Queue is not empty")
	return false
}
