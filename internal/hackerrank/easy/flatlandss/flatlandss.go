package flatlandss

import (
	"math"
	"slices"

	"github.com/yossrez/prep-go/internal/hackerrank"
)

/*
 * Flatland is a country with a number of cities,
 * some of which have space station. Cities are numbered
 * consecutively and each has a road of 1km length connecting
 * it to the next city.
 *
 * It is not a circular route, so the first ciry doesn't
 * connect with the last city.
 *
 * Determine the maximum distance from any city to its
 * nearest space station.
 */
// Params:
// 	- n: the number of cites
// 	- c: the indices of cities with a space station
func flatlandSpaceStation(n int32, c []int32) int32 {
	var maxDist int32
	cNum := int32(len(c))
	if n == cNum {
		return maxDist
	}

	if n-cNum == 1 {
		maxDist = 1
		return maxDist
	}

	slices.Sort(c)

	var currCity int32
	var cIdx int32
	var currStation int32
	var hasNextStation bool
	for currCity < n {
		currStation = c[cIdx]
		if currCity == currStation {
			currCity++
			continue
		}
		cIdx++
		hasNextStation = cIdx < cNum
		if hasNextStation && currCity == c[cIdx] {
			currCity++
			continue
		} else {
			cIdx--
		}

		// calc distance
		var a int32 = math.MaxInt32
		var b int32 = math.MaxInt32
		if cIdx-1 >= 0 {
			a = c[cIdx-1] - currCity
			if a < 0 {
				a = -a
			}
		}
		if hasNextStation {
			b = c[cIdx+1] - currCity
		}
		x := currStation - currCity
		if x < 0 {
			x = -x
		}
		nearest := min(min(a, b), x)
		if maxDist < nearest {
			maxDist = nearest
		}
		currCity++
	}

	return maxDist
}

// main
func Run() {
	// expect 2
	_ = flatlandSpaceStation(5, []int32{4, 0})
}

func init() {
	meta := hackerrank.Meta{
		Problem:    "flatlandss",
		Skills:     hackerrank.ProblemSolvingIntermediate,
		Difficulty: hackerrank.Easy,
		Subdomain:  hackerrank.Implementation,
	}
	err := hackerrank.Registry.Register(meta.Problem, Run, meta)
	if err != nil {
		panic(err)
	}
}
