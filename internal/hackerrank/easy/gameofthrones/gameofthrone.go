package gameofthrones

import "github.com/yossrez/prep-go/internal/hackerrank/runner"

func gameOfThrones(s string) string {
	token := map[rune]int{}
	for _, r := range s {
		token[r] += 1
	}
	oddOne := false
	for _, count := range token {
		if count%2 == 1 {
			if oddOne {
				return "NO"
			} else {
				oddOne = true
			}
		}
	}
	return "YES"
}

func Run() {
	// expect YES
	_ = gameOfThrones("cdcdcdcdeeeef")
}

func init() {
	runner.Register(
		runner.HackerRankMeta{
			Problem:    "gameofthrones",
			Skills:     runner.ProblemSolvingBasic,
			Difficulty: runner.Easy,
			Subdomain:  runner.Strings,
		},
		Run)
}
