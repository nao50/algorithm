package fizzbuzz

import "testing"

func Testfizzbuzz(t *testing.T) {
	tests := []struct {
		input int
		want  string
		err   bool
	}{
		{input: -1, want: "", err: true},
		{input: 0, want: "", err: true},
		{input: 1, want: "1", err: false},
		{input: 2, want: "2", err: false},
		{input: 3, want: "Fizz", err: false},
		{input: 5, want: "Buzz", err: false},
		{input: 15, want: "FizzBuzz", err: false},
		{input: 101, want: "", err: true},
	}

	for _, test := range tests {
		got, err := fizzbuzz(test.input)
		if !test.err && err != nil {
			t.Fatalf("should not be error for %v but: %v", test.input, err)
		}
		if test.err && err == nil {
			t.Fatalf("should be error for %v but not:", test.input)
		}
		if got != test.want {
			t.Fatalf("want %q, but %q:", test.want, got)
		}
	}
}
