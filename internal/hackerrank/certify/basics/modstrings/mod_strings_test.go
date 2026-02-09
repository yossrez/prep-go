package modstrings

import (
	"testing"
)

func TestModifyString_Table(t *testing.T) {
	cases := []struct {
		name     string
		args     string
		expected string
	}{
		{"Case_1", "   oll123eh  ", "hello"},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := ModifyString(tt.args)
			if tt.expected != got {
				t.Fatalf("expected %s, got %s", tt.expected, got)
			}
		})
	}
}
