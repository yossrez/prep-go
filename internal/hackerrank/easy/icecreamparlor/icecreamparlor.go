package icecreamparlor

import "github.com/yossrez/prep-go/internal/hackerrank"

/*
 * Given a list of prices for the flavors of
 * ice cream, select the two that will cost
 * all of the money they have.
 */
// use 1-based indexing for the result
func iceCreamParlor(m int32, arr []int32) []int32 {
	var flavor [2]int32
	cache := map[int32]int32{}
	for i, cost := range arr {
		idxOneBased := int32(i) + 1
		diff := m - cost
		if diff < 0 {
			continue
		}

		pos, ok := cache[diff]
		if ok {
			flavor[0] = pos
			flavor[1] = idxOneBased
			break
		} else {
			cache[cost] = idxOneBased
		}
	}
	return flavor[:]
}

// main
func Run() {
	// expect [1, 4]
	_ = iceCreamParlor(6, []int32{1, 3, 4, 5, 6})
}

func init() {
	meta := hackerrank.Meta{
		Problem:    "icecreamparlor",
		Skills:     hackerrank.ProblemSolvingIntermediate,
		Difficulty: hackerrank.Easy,
		Subdomain:  hackerrank.Search,
	}
	err := hackerrank.Registry.Register(meta.Problem, Run, meta)
	if err != nil {
		panic(err)
	}
}
