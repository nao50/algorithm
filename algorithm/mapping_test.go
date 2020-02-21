package main

import (
	"reflect"
	"testing"
)

func TestMapping(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{input: []int{}, want: []int{}},
		{input: []int{1, 2}, want: []int{1, 4}},
		{input: []int{1, 2, 3}, want: []int{1, 4, 9}},
	}

	for _, test := range tests {
		got := Mapping(square, test.input)
		if !reflect.DeepEqual(got, test.want) {
			t.Fatalf("\n want: %v, \n actual: %v", test.want, got)
		}
	}
}
