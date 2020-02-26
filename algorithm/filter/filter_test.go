package filter

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{input: []int{1, 2}, want: []int{2}},
		{input: []int{1, 2, 3, 4, 5}, want: []int{2, 4}},
		{input: []int{}, want: []int{}},
	}

	for _, test := range tests {
		got := Filter(isEven, test.input)
		if !reflect.DeepEqual(got, test.want) {
			t.Fatalf("\n want: %v, \n actual: %v", test.want, got)
		}
	}
}
