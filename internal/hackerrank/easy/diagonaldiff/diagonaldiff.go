package diagonaldiff

import (
	"github.com/yossrez/prep-go/internal/hackerrank"
)

/*
 * Given a square matrix,
 * calculate the absolute difference
 * between the sums of its diagonals.
 */
func diagonalDiff(arr [][]int32) int32 {
	l := 0
	r := len(arr[0]) - 1
	var dl int32
	var dr int32
	for _, list := range arr {
		dl += list[l]
		dr += list[r]
		l++
		r--
	}
	if dl < dr {
		return dr - dl
	}
	return dl - dr
}

// main
func Run() {
	// expect 2
	matrix := [][]int32{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}
	_ = diagonalDiff(matrix)
}

func init() {
	meta := hackerrank.Meta{
		Problem:    "diagonaldiff",
		Skills:     hackerrank.ProblemSolvingBasic,
		Difficulty: hackerrank.Easy,
		Subdomain:  hackerrank.Warmup,
	}
	err := hackerrank.Registry.Register(meta.Problem, Run, meta)
	if err != nil {
		panic(err)
	}
}
