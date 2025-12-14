package main

import (
	"reflect"
	"testing"
)

func TestWordFrequency(t *testing.T) {
	test := []struct {
		name string
		text string
		want map[string]int
	}{
		{
			name: "Test success word frequency",
			text: "Hola mundo, hola Go! Mundo de Go",
			want: map[string]int{
				"hola":  2,
				"mundo": 2,
				"go":    2,
				"de":    1,
			},
		}, {
			name: "Test empty string",
			text: "",
			want: map[string]int{},
		}, {
			name: "Test only symbols",
			text: "%&!",
			want: map[string]int{},
		}}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			got := wordFrequency(tt.text)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("wordFrequency() = %v, want %v", got, tt.want)
			}
		})
	}
}
