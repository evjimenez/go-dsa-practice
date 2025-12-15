package main

import "fmt"

type node struct {
	value []int
	next  *node
}

type headNode struct {
	head *node
}

func (h *headNode) Enqueue(value int) {
	h.head.value = append(h.head.value, value)
}

func (h *headNode) Dequeue() (int, bool) {
	if h.head == nil {
		return 0, false
	}

	value := h.head.value[0]
	h.head = h.head.next
	return value, true
}

func main() {
	newNode := headNode{head: &node{}}
	newNode.Enqueue(10)
	newNode.Enqueue(20)
	newNode.Enqueue(30)
	fmt.Println(newNode.head.value)
	value, isSuccess := newNode.Dequeue()
	fmt.Println(value, isSuccess)
}
