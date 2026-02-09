package savetheprisoner

import "github.com/yossrez/prep-go/internal/hackerrank/runner"

/*
 * Complete the saveThePrisoner func.
 * It should return an integer representing
 * the chair number of the prisoner to warn.
 */
func saveThePrisoner(n, m, s int32) int32 {
	v := m % n

	if v == 1 {
		return s
	}

	s = s + m - 1
	if s > n {
		s = s % n
		if s == 0 {
			return n
		}
	}
	return s
}

// main
func Run() {
	// expect 6
	_ = saveThePrisoner(7, 19, 2)
}

func init() {
	runner.Register(runner.HackerRankMeta{
		Problem:    "savetheprisoner",
		Skills:     runner.ProblemSolvingBasic,
		Difficulty: runner.Easy,
		Subdomain:  runner.Implementation},
		Run)
}
