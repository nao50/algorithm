package main

import (
	"reflect"
	"testing"
)

func TestReduce(t *testing.T) {
	tests := []struct {
		input []int
		init  int
		want  int
	}{
		{input: []int{}, init: 0, want: 0},
		{input: []int{1, 2, 3}, init: 0, want: 6},
		{input: []int{1, 2, 3, 4, 5}, init: 1, want: 161},
	}

	for _, test := range tests {
		got := Reduce(add, test.init, test.input)
		if !reflect.DeepEqual(got, test.want) {
			t.Fatalf("\n want: %v, \n actual: %v", test.want, got)
		}
	}
}
