package beutifulpairs

import (
	"math"

	"github.com/yossrez/prep-go/internal/hackerrank"
)

// Complete the beutifulPairs function below.
// It should return an integer that represent
// the maximum number of pairwise disjoint
// beutiful pairs that can be formed.
//
// A pair of indices (i,j) is beutiful
// if the i-th element of array 'A' is
// equal to the j-th element of array 'B'.
// A set containing beutiful pairs is called
// a beutiful set.
//
// A beutiful set is calles pairwise disjoint
// if for every pair (l[i],r[i]) belonging to the
// set there is no repetition of either l[i] or r[i]
// values.
//
// Exp:
//   - set [(1,2),(1,3),(3,4)] is not pairwise disjoint
//     as there is repetition of 1, that is l[0][0] = r[1][0]
//     or (1,2) = (1,3)
//
// Note:
//   - You must first change 1 element in B, and your choice
//     of element must be optimal
func beutifulPairs(A, B []int32) int32 {
	var maxPair int

	lMap := make(map[int]struct{})
	rMap := make(map[int]struct{})

	n := len(A)
	const denominator = 2
	var tasks [denominator]int
	tasks[0] = int(math.Floor((float64(n / denominator))))
	tasks[1] = n - tasks[0]

	n = 0
	count := 0
	for _, t := range tasks {
		for i, v := range A {
			for j := range t {
				j += n
				if v == B[j] {
					if _, ok := lMap[i]; ok {
						continue
					}
					if _, ok := rMap[j]; ok {
						continue
					}
					lMap[i] = struct{}{}
					rMap[j] = struct{}{}
					count++
				}
			}
			if count == t {
				break
			}
		}
		count = 0
		n = t
	}

	// I think we don't really need to change 1 element
	// in B array, just use little logic we know the max pairs.
	n = len(A)
	maxPair = len(lMap)
	if maxPair < n {
		maxPair++
	} else if maxPair == n {
		maxPair--
	}
	return int32(maxPair)
}

// main
func Run() {
	// expect 4
	_ = beutifulPairs(
		[]int32{1, 2, 3, 4},
		[]int32{1, 2, 3, 3},
	)
}

func init() {
	meta := hackerrank.Meta{
		Problem:    "beutifulpairs",
		Skills:     hackerrank.ProblemSolvingIntermediate,
		Difficulty: hackerrank.Easy,
		Subdomain:  hackerrank.Greedy,
	}
	err := hackerrank.Registry.Register(meta.Problem, Run, meta)
	if err != nil {
		panic(err)
	}
}
