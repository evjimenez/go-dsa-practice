package main

import (
	"testing"
)

func TestRotationSlice(t *testing.T) {
	type args struct {
		numbers  []int
		position int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{{
		name: "Test success rotation slice",
		args: args{
			numbers:  []int{1, 2, 3, 4, 5},
			position: 2,
		},
		want: []int{4, 5, 1, 2, 3},
	}, {
		name: "Test error empty slice",
		args: args{
			numbers:  []int{},
			position: 3,
		},
		want: []int{},
	}, {
		name: "Test error zero position",
		args: args{
			numbers:  []int{10, 20, 30},
			position: 0,
		},
		want: []int{10, 20, 30},
	}, {
		name: "Test error",
		args: args{
			numbers:  []int{10, 20, 30},
			position: 2,
		},
		want: []int{20, 30, 10},
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotationSlice(tt.args.numbers, tt.args.position)
		})
	}
}
