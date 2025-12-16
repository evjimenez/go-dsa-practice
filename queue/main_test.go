package main

import "testing"

func Test_queue_Dequeue(t *testing.T) {
	type fields struct {
		data []int
		head int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
		want1  bool
	}{
		{
			name: "Test success dequeue",
			fields: fields{
				data: []int{10, 20, 30},
			},
			want:  10,
			want1: true,
		},
		{
			name: "Test error head greater than len data",
			fields: fields{
				data: []int{50, 23, 23},
				head: 5,
			},
			want:  0,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &queue{
				data: tt.fields.data,
				head: tt.fields.head,
			}
			got, got1 := q.Dequeue()
			if got != tt.want {
				t.Errorf("Dequeue() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Dequeue() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_queue_Enqueue(t *testing.T) {
	type fields struct {
		data []int
		head int
	}
	type args struct {
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test success enqueue",
			fields: fields{
				data: []int{},
				head: 0,
			},
			args: args{value: 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &queue{
				data: tt.fields.data,
				head: tt.fields.head,
			}
			q.Enqueue(tt.args.value)
		})
	}
}
