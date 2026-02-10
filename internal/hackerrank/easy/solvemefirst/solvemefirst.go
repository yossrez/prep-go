package solvemefirst

import (
	"fmt"

	"github.com/yossrez/prep-go/internal/hackerrank"
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
	meta := hackerrank.Meta{
		Problem:    "solvemefirst",
		Skills:     hackerrank.ProblemSolvingBasic,
		Difficulty: hackerrank.Easy,
		Subdomain:  hackerrank.Warmup,
	}
	err := hackerrank.Registry.Register(meta.Problem, Run, meta)
	if err != nil {
		panic(err)
	}
}
