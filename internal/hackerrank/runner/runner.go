package runner

import (
	"fmt"
)

type SkillLevel int

const (
	Unrated SkillLevel = iota
	ProblemSolvingBasic
	ProblemSolvingIntermediate
	ProblemSolvingAdvanced
)

type Difficulty int

const (
	Unrank Difficulty = iota
	Easy
	Medium
	Hard
)

type Subdomain int

const (
	None Subdomain = iota
	Warmup
	Implementation
	Strings
	Sorting
	Search
	GraphTheory
	Greedy
	DynamicProgramming
	ConstructiveAlgorithms
	BitManipulation
	Recursion
	GameTheory
	NPComplete
	Debugging
)

type HackerRankMeta struct {
	Problem    string
	Skills     SkillLevel
	Difficulty Difficulty
	Subdomain  Subdomain
}

var registry = map[string]func(){}

func Register(meta HackerRankMeta, fn func()) {
	registry[meta.Problem] = fn
}

func Execute(name string) error {
	fn, ok := registry[name]
	if !ok {
		return fmt.Errorf("unknown exercise: %s, cmd: hackerrank", name)
	}
	fn()
	return nil
}

func List() []string {
	keys := make([]string, 0, len(registry))
	for k := range registry {
		keys = append(keys, k)
	}
	return keys
}
