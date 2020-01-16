package main

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestQuicksort(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{input: []int{5, 6, 4, 7, 3, 8, 2, 9, 1, 0}, want: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}

	for _, test := range tests {
		Quicksort(0, 9, test.input)
		if !reflect.DeepEqual(test.input, test.want) {
			t.Errorf("\n want: %v, \n actual: %v", test.want, test.input)
		}
	}
}

func BenchmarkQuicksort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmp := make([]int, len(Data))
		copy(tmp, Data)

		Quicksort(0, 9, tmp)
	}
}

func BenchmarkQuickSorted(b *testing.B) {
	tmp := make([]int, len(SortedData))

	for i := 0; i < b.N; i++ {
		copy(tmp, SortedData)

		Quicksort(0, 9, tmp)
	}
}

func BenchmarkQuickSortRandomizedData(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Seed(time.Now().UnixNano())
		tmp := rand.Perm(len(Data))

		Quicksort(0, 9, tmp)
	}
}
