package gameofthrones

import "testing"

func TestGameOfThrones(t *testing.T) {
	cases := []struct {
		name     string
		s        string
		expected string
	}{
		{
			"Case_1",
			"cdcdcdcdeeeef",
			"YES",
		},
		{
			"Case_2",
			"cdefghmnopqrstuvw",
			"NO",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := gameOfThrones(tt.s)
			if tt.expected != got {
				t.Fatalf("expected %s, got %s", tt.expected, got)
			}
		})
	}
}
