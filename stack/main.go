package main

import "fmt"

type stack []int

func main() {
	s := stack{}
	s.push(10)
	s.push(20)
	s.push(30)
	fmt.Println(s)
	s.pop()
	s.peek()
	fmt.Println(s)
	s.pop()
	fmt.Println(s)
}

func (s *stack) push(val int) {
	*s = append(*s, val)
}

func (s *stack) pop() (int, bool) {
	if len(*s) == 0 {
		return 0, false
	}

	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return top, true
}

func (s stack) peek() (int, bool) {
	if len(s) == 0 {
		return 0, false
	}

	top := (s)[len(s)-1]
	return top, true
}
