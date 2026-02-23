package separatethenumbers

import (
	"reflect"
	"testing"
)

func TestSeparateTheNumbers(t *testing.T) {
	cases := []struct {
		name     string
		s        []string
		expected []string
	}{
		{
			"Case_1",
			[]string{
				"1234",
				"91011",
				"99100",
				"101103",
				"010203",
				"13",
				"1",
			},
			[]string{
				"YES 1",
				"YES 9",
				"YES 99",
				"NO",
				"NO",
				"NO",
				"NO",
			},
		},
		{
			"Case_2",
			[]string{
				"010203",
				"1",
			},
			[]string{
				"NO",
				"NO",
			},
		},
		{
			"Case_3",
			[]string{
				"10203",
				"312",
				"10001001",
				"100110021003",
				"1002100310041005",
				"9991999299939994",
				"9999999999999999",
			},
			[]string{
				"NO",
				"NO",
				"YES 1000",
				"YES 1001",
				"YES 1002",
				"YES 9991",
				"NO",
			},
		},
	}

	for _, tt := range cases {
		// clear slice
		// got = make([]string, 0)
		got = got[:0]
		t.Run(tt.name, func(t *testing.T) {
			for _, v := range tt.s {
				separateTheNumbers(v)
			}
			if !reflect.DeepEqual(tt.expected, got) {
				t.Fatalf("expected %s, got %s", tt.expected, got)
			}
		})
	}
}
