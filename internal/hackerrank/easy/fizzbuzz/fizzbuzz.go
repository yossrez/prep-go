package fizzbuzz

import (
	"fmt"

	"github.com/yossrez/prep-go/internal/hackerrank/runner"
)

// for testing purposes
var testBucket []string

func fizzBuzz(n int32) {
	// Write your code here
	var bucket []string = make([]string, n)
	var i int32 = 0
	f, b := false, false
	for i < n {
		i++
		f = i%3 == 0
		b = i%5 == 0

		if !f && !b {
			bucket[i-1] = fmt.Sprintf("%d", i)
			f, b = false, false
			continue
		}
		if f {
			bucket[i-1] += "Fizz"
		}
		if b {
			bucket[i-1] += "Buzz"
		}
	}

	for _, val := range bucket {
		fmt.Println(val)
	}
	// for testing purposes
	testBucket = bucket
}

// main
func Run() {
	fizzBuzz(15)
}

// ran automatically when the package imported
func init() {
	meta := runner.HackerRankMeta{}
	meta.Problem = "fizzbuzz"
	meta.Skills = runner.ProblemSolvingBasic
	meta.Difficulty = runner.Easy
	meta.Subdomain = runner.Warmup
	runner.Register(meta, Run)
}
