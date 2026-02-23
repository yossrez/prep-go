package wus

import (
	"reflect"
	"testing"
)

func TestWeightedUniformStrings(t *testing.T) {
	cases := []struct {
		name     string
		s        string
		queries  []int32
		expected []string
	}{
		{
			"Case_1",
			"abbcccdddd",
			[]int32{1, 7, 5, 4, 15},
			[]string{"Yes", "No", "No", "Yes", "No"},
		},
		{
			"Case_2",
			"abccddde",
			[]int32{1, 3, 12, 5, 9, 10},
			[]string{"Yes", "Yes", "Yes", "Yes", "No", "No"},
		},
		{
			"Case_3",
			"aaabbbbcccddd",
			[]int32{9, 7, 8, 12, 5},
			[]string{"Yes", "No", "Yes", "Yes", "No"},
		},
		{
			"Case_4",
			"abccdcdde",
			[]int32{1, 3, 12, 5, 9, 10},
			[]string{"Yes", "Yes", "No", "Yes", "No", "No"},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := weightedUniformStrings(tt.s, tt.queries)
			if !reflect.DeepEqual(tt.expected, got) {
				t.Fatalf("expected %s, got %s", tt.expected, got)
			}
		})
	}
}
