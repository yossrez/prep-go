package fizzbuzz

import (
	"fmt"

	"github.com/yossrez/prep-go/internal/hackerrank/runner"
)

// for testing purposes
var testBucket []string

/*
 * Given an integers,
 * print 'Fizz' if the number divisible by 3
 * print 'Buzz' if the number divisible by 5
 * print 'FizzBuzz' if the number divisible by 3 and 5
 * otherwise print the number itself.
 */
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
	runner.Register(
		runner.HackerRankMeta{
			Problem:    "fizzbuzz",
			Skills:     runner.ProblemSolvingBasic,
			Difficulty: runner.Easy,
			Subdomain:  runner.Warmup,
		},
		Run)
}
