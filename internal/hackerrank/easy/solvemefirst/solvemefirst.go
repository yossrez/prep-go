package solvemefirst

import (
	"fmt"

	"github.com/yossrez/prep-go/internal/hackerrank/runner"
)

func solveMeFirst(a uint32, b uint32) uint32 {
	// Hint: Type return (a+b) below
	return a + b
}

// main
func Run() {
	var a, b, res uint32
	fmt.Scanf("%v\n%v", &a, &b)
	res = solveMeFirst(a, b)
	fmt.Println(res)
}

func init() {
	runner.Register(
		runner.HackerRankMeta{
			Problem:    "solvemefirst",
			Skills:     runner.ProblemSolvingBasic,
			Difficulty: runner.Easy,
			Subdomain:  runner.Warmup,
		},
		Run)
}
