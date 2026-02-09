package pickingnumbers

import (
	"github.com/yossrez/prep-go/internal/hackerrank/runner"
)

/*
 * Given an array of integers,
 * find the longest subarray
 * where the absolute difference between
 * any two elements is less than or equal to 1.
 */
func pickingNumbers(a []int32) int32 {
	var longSubArrCount int32
	token := map[int32]int32{}

	for _, val := range a {
		token[val] += 1
	}

	for key, val := range token {
		var count int32 = val
		var sub int32
		for _, n := range [2]int32{1, -1} {
			pair, ok := token[key+n]
			if ok {
				if pair > sub {
					sub = pair
				}
			}
			if val == 1 {
				break
			}
		}
		count += sub
		if longSubArrCount < count {
			longSubArrCount = count
		}
	}
	// fmt.Println(token)
	return longSubArrCount
}

// main
func Run() {
	// expect 5
	_ = pickingNumbers([]int32{1, 2, 2, 3, 1, 2})
}

func init() {
	runner.Register(
		runner.HackerRankMeta{
			Problem:    "pickingnumbers",
			Skills:     runner.ProblemSolvingBasic,
			Difficulty: runner.Easy,
			Subdomain:  runner.Implementation,
		},
		Run)
}
