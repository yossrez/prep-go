package diagonaldiff

import "testing"

func TestDiagonalDiff(t *testing.T) {
	cases := []struct {
		name     string
		matrix   [][]int32
		expected int32
	}{
		{
			"Case_1",
			[][]int32{
				{1, 2, 3},
				{4, 5, 6},
				{9, 8, 9}},
			2,
		},
		{
			"Case_2",
			[][]int32{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16}},
			0,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := diagonalDiff(tt.matrix)
			if tt.expected != got {
				t.Fatalf("expected %d, got %d", tt.expected, got)
			}
		})
	}
}
