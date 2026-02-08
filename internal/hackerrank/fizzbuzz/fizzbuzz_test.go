package fizzbuzz

import (
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	cases := []struct {
		name     string
		n        int32
		expected []string
	}{
		{"Case_1", 15, []string{
			"1", "2", "Fizz", "4", "Buzz", "Fizz", "7",
			"8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz",
		}},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			fizzBuzz(tt.n)
			got := testBucket
			if !reflect.DeepEqual(tt.expected, got) {
				t.Fatalf("expected %s, got %s", tt.expected, got)
			}
		})
	}
}
