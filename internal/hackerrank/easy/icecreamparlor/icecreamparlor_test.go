package icecreamparlor

import (
	"reflect"
	"testing"
)

func TestIceCreamParlor(t *testing.T) {
	cases := []struct {
		name     string
		money    int32
		costs    []int32
		expected [2]int32
	}{
		{
			"Case_1",
			6,
			[]int32{1, 3, 4, 5, 6},
			[2]int32{1, 4},
		},
		{
			"Case_2",
			4,
			[]int32{1, 4, 5, 3, 2},
			[2]int32{1, 4},
		},
		{
			"Case_3",
			4,
			[]int32{2, 2, 4, 3},
			[2]int32{1, 2},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := iceCreamParlor(tt.money, tt.costs)
			if !reflect.DeepEqual(tt.expected[:], got) {
				t.Fatalf("expected %d, got %d", tt.expected, got)
			}
		})
	}
}
