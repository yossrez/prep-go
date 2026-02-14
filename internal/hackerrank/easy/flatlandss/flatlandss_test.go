package flatlandss

import "testing"

func TestFlatlandSpaceStation(t *testing.T) {
	cases := []struct {
		name     string
		n        int32
		c        []int32
		expected int32
	}{
		{
			"Case_1",
			6,
			[]int32{0, 1, 2, 4, 3, 5},
			0,
		},
		{
			"Case_2",
			5,
			[]int32{4, 0, 1, 2},
			1,
		},
		{
			"Case_3",
			5,
			[]int32{0, 3, 4},
			1,
		},
		{
			"Case_4",
			5,
			[]int32{2, 3, 4},
			2,
		},
		{
			"Case_5",
			5,
			[]int32{0, 2, 4},
			1,
		},
		{
			"Case_6",
			5,
			[]int32{4, 0},
			2,
		},
		{
			"Case_7",
			5,
			[]int32{4, 1},
			1,
		},
		{
			"Case_8",
			5,
			[]int32{4, 3},
			3,
		},
		{
			"Case_9",
			5,
			[]int32{4},
			4,
		},
		{
			"Case_10",
			3,
			[]int32{1},
			1,
		},
		{
			"Case_11",
			3,
			[]int32{0},
			2,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := flatlandSpaceStation(tt.n, tt.c)
			if tt.expected != got {
				t.Fatalf("expected %d, got %d", tt.expected, got)
			}
		})
	}
}
