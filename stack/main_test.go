package main

import (
	"reflect"
	"testing"
)

func TestPush(t *testing.T) {
	s := &stack{}
	s.push(5)
	if !reflect.DeepEqual([]int(*s), []int{5}) {
		t.Fatalf("push() stack=%v, want %v", *s, []int{5})
	}
}

func TestPop(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		s := stack{1, 2, 3}
		got, ok := s.pop()
		if !ok {
			t.Fatalf("pop() ok=false, want true")
		}
		if got != 3 {
			t.Fatalf("pop() got=%d, want %d", got, 3)
		}
		if !reflect.DeepEqual([]int(s), []int{1, 2}) {
			t.Fatalf("after pop stack=%v, want %v", s, []int{1, 2})
		}
	})

	t.Run("empty", func(t *testing.T) {
		s := stack{}
		got, ok := s.pop()
		if ok {
			t.Fatalf("pop() ok=true, want false")
		}
		if got != 0 {
			t.Fatalf("pop() got=%d, want %d", got, 0)
		}
	})
}

func TestPeek(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		s := stack{10, 20, 30}
		got, ok := s.peek()
		if !ok {
			t.Fatalf("peek() ok=false, want true")
		}
		if got != 30 {
			t.Fatalf("peek() got=%d, want %d", got, 30)
		}
		// peek no debe modificar
		if !reflect.DeepEqual([]int(s), []int{10, 20, 30}) {
			t.Fatalf("after peek stack=%v, want %v", s, []int{10, 20, 30})
		}
	})

	t.Run("empty", func(t *testing.T) {
		s := stack{}
		got, ok := s.peek()
		if ok {
			t.Fatalf("peek() ok=true, want false")
		}
		if got != 0 {
			t.Fatalf("peek() got=%d, want %d", got, 0)
		}
	})
}
