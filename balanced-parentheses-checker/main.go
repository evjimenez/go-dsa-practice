package main

import "fmt"

type stack []rune

func (s *stack) push(value rune) {
	*s = append(*s, value)
}

func (s *stack) pop() (rune, bool) {
	if len(*s) == 0 {
		return 0, false
	}

	value := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return value, true
}

func (s *stack) peek() (rune, bool) {
	if len(*s) == 0 {
		return 0, false
	}
	value := (*s)[len(*s)-1]
	return value, true
}

func (s *stack) validateChar(text string) bool {
	for _, v := range text {
		if (v == ')' || v == ']') || v == '}' {
			value, isValid := s.peek()
			if !isValid {
				return false
			}

			if value < v {
				s.pop()
				continue
			} else {
				return false
			}

		} else {
			s.push(v)
		}

	}

	if len(*s) == 0 {
		return true
	}

	return false
}

func main() {
	s := stack{}
	isValid := s.validateChar("{[()]}")
	if isValid {
		fmt.Println(isValid)
	} else {
		fmt.Println(isValid)
	}
}
