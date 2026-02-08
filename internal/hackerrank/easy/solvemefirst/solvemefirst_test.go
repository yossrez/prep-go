package solvemefirst

import (
	"testing"
)

func TestSolveMeFirst(t *testing.T) {
	cases := []struct {
		name     string
		a, b     uint32
		expected uint32
	}{
		{"Case_1", 2, 3, 5},
		{"Case_2", 100, 1000, 1100},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := solveMeFirst(tt.a, tt.b)
			if !(tt.expected == got) {
				t.Fatalf("expected %d, got %d", tt.expected, got)
			}
		})
	}
}
