package beutifulpairs

import (
	"testing"
)

func TestBeutifulPairs(t *testing.T) {
	cases := []struct {
		name     string
		a        []int32
		b        []int32
		expected int32
	}{
		{
			"Case_1",
			[]int32{1, 2, 3, 4},
			[]int32{1, 2, 3, 3},
			4,
		},
		{
			"Case_2",
			[]int32{3, 5, 7, 11, 5, 8},
			[]int32{5, 7, 11, 10, 5, 8},
			6,
		},
		{
			"Case_3",
			[]int32{3, 5, 7, 12},
			[]int32{5, 7, 11, 10},
			3,
		},
		{
			"Case_4",
			[]int32{3, 5, 7, 12},
			[]int32{9, 1, 11, 10},
			1,
		},
		{
			"Case_5",
			[]int32{3, 5, 7, 12},
			[]int32{3, 5, 7, 12},
			3,
		},
		{
			"Case_6",
			[]int32{1, 2, 3, 4},
			[]int32{3, 3, 2, 1},
			4,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := beutifulPairs(tt.a, tt.b)
			if got != tt.expected {
				t.Fatalf("expected %d, got %d", tt.expected, got)
			}
		})
	}
}
