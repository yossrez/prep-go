package gameofthrones

import "github.com/yossrez/prep-go/internal/hackerrank"

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
	meta := hackerrank.Meta{
		Problem:    "gameofthrones",
		Skills:     hackerrank.ProblemSolvingBasic,
		Difficulty: hackerrank.Easy,
		Subdomain:  hackerrank.Strings,
	}
	err := hackerrank.Registry.Register(meta.Problem, Run, meta)
	if err != nil {
		panic(err)
	}
}
