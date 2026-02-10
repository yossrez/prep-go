package superreducedstring

import "testing"

func TestSuperReducedString(t *testing.T) {
	cases := []struct {
		name     string
		s        string
		expected string
	}{
		{
			"Case_1",
			"aab",
			"b",
		},
		{
			"Case_2",
			"abba",
			"Empty String",
		},
		{
			"Case_3",
			"aaabccddd",
			"abd",
		},
		{
			"Case_4",
			"aba",
			"aba",
		},
		{
			"Case_5",
			"abcdefaghijklmmlnob",
			"abcdefaghijknob",
		},
		{
			"Case_6",
			"abcdeedcba",
			"Empty String",
		},
		{
			"Case_7",
			"a",
			"a",
		},
		{
			"Case_8",
			"aa",
			"Empty String",
		},
		{
			"Case_9",
			"baabccddd",
			"d",
		},
		{
			"Case_10",
			"ababbaaa",
			"ab",
		},
		{
			"Case_11",
			"ab",
			"ab",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := superReducedString(tt.s)
			if tt.expected != got {
				t.Fatalf("expected '%s', got '%s'", tt.expected, got)
			}
		})
	}
}
