package main

import "testing"

func Test_stack_validateChar(t *testing.T) {
	type args struct {
		text string
	}

	tests := []struct {
		name string
		s    stack
		args args
		want bool
	}{
		{
			name: "Test success parentesis",
			s:    stack{},
			args: args{text: "()"},
			want: true,
		},
		{
			name: "Test success parentesis y corchete",
			s:    stack{},
			args: args{text: "([])"},
			want: true,
		},
		{
			name: "Test error parentesis y corchete",
			s:    stack{},
			args: args{text: "([)]"},
			want: false,
		},
		{
			name: "Test error parentesis",
			s:    stack{},
			args: args{text: "((()"},
			want: false,
		},
		{
			name: "Test success parentesis, llave y corchete",
			s:    stack{},
			args: args{text: "{[()]}"},
			want: true,
		},
		{
			name: "Test error parentesis, llevae y corchete",
			s:    stack{},
			args: args{text: "{[(])}"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.validateChar(tt.args.text); got != tt.want {
				t.Errorf("validateChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
