package twocharacters

import "testing"

func TestAlternate(t *testing.T) {
	cases := []struct {
		name     string
		s        string
		expected int32
	}{
		{"Case_1", "beabeefeab", 5},
		{"Case_2", "a", 0},
		{"Case_3", "aa", 0},
		{"Case_4", "ab", 2},
		{"Case_5", "asdcbsdcagfsdbgdfanfghbsfdab", 8},
		{"Case_6", "asvkugfiugsalddlasguifgukvsa", 0},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := alternate(tt.s)
			if tt.expected != got {
				t.Fatalf("expected %d, got %d", tt.expected, got)
			}
		})
	}
}
