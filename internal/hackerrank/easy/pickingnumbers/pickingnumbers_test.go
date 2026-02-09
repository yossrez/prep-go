package pickingnumbers

import (
	"testing"
)

func TestPickingNumbers(t *testing.T) {
	cases := []struct {
		name     string
		arr      []int32
		expected int32
	}{
		{"Case_1", []int32{1, 1, 2, 2, 4, 4, 5, 5, 5}, 5},
		{"Case_2", []int32{4, 6, 5, 3, 3, 1}, 3},
		{"Case_3", []int32{1, 2, 2, 3, 1, 2}, 5},
		{"Case_4", []int32{1, 2}, 2},
		{"Case_5", []int32{1, 3}, 1},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := pickingNumbers(tt.arr)
			if tt.expected != got {
				t.Fatalf("expected %d, got %d", tt.expected, got)
			}
		})
	}
}
