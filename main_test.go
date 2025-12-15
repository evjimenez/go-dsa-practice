package main

import "testing"

func Test_headNode_Dequeue(t *testing.T) {
	type fields struct {
		head *node
	}
	tests := []struct {
		name      string
		fields    fields
		want      int
		isSuccess bool
	}{
		{
			name:      "Test success",
			fields:    fields{head: &node{[]int{10, 20, 30}, nil}},
			want:      10,
			isSuccess: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &headNode{
				head: tt.fields.head,
			}
			got, got1 := h.Dequeue()
			if got != tt.want {
				t.Errorf("Dequeue() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.isSuccess {
				t.Errorf("Dequeue() got1 = %v, want %v", got1, tt.isSuccess)
			}
		})
	}
}
