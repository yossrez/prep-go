package savetheprisoner

import (
	"testing"
)

func TestSaveThePrisoner(t *testing.T) {
	cases := []struct {
		name     string
		n, m, s  int32
		expected int32
	}{
		{"Case_1", 4, 6, 2, 3},
		{"Case_2", 5, 2, 1, 2},
		{"Case_3", 5, 2, 2, 3},
		{"Case_4", 7, 19, 2, 6},
		{"Case_5", 3, 7, 3, 3},
		{"Case_6", 3, 3, 3, 2},
		{"Case_7", 3, 3, 1, 3},
		{"Case_8", 3, 4, 1, 1},
		{"Case_9", 5, 10, 1, 5},
		{"Case_10", 5, 2, 5, 1},
		{"Case_11", 5, 1, 5, 5},
		{"Case_12", 100, 110, 100, 9},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := saveThePrisoner(tt.n, tt.m, tt.s)
			if tt.expected != got {
				t.Fatalf("expected %d, got %d", tt.expected, got)
			}
		})
	}
}
