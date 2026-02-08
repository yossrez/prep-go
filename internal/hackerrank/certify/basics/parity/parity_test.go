package parity

import (
	"reflect"
	"sync"
	"testing"
)

func TestParity_Table(t *testing.T) {
	cases := []struct {
		name     string
		size     int32
		arr      []int32
		expected []int32
	}{
		{"Case_1", 5, []int32{1, 2, 3, 4, 5}, []int32{1, 3, 5, 2, 4}},
	}

	var wg sync.WaitGroup

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			oddChan := make(chan int32)
			evenChan := make(chan int32)
			serverChan = make(chan in, len(tt.arr))
			for idx := 0; idx < len(tt.arr); idx++ {
				i := idx
				serverChan <- in{tt.arr[i], oddChan, evenChan}
			}

			odds, evens := []int32{}, []int32{}
			wg.Add(len(tt.arr))
			go func() {
				for newOdd := range oddChan {
					odds = append(odds, newOdd)
					wg.Done()
				}
			}()
			go func() {
				for newEven := range evenChan {
					evens = append(evens, newEven)
					wg.Done()
				}
			}()
			go Server()
			wg.Wait()

			var got []int32
			for _, resultItem := range odds {
				got = append(got, resultItem)
			}
			for _, resultItem := range evens {
				got = append(got, resultItem)
			}

			if !reflect.DeepEqual(tt.expected, got) {
				t.Fatalf("expected %d, got %d", tt.expected, got)
			}
		})
	}
}
