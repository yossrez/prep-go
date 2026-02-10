package hackerrank

import "github.com/yossrez/prep-go/pkg/registrar"

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

type Meta struct {
	Problem    string
	Skills     SkillLevel
	Difficulty Difficulty
	Subdomain  Subdomain
}

var Registry *registrar.Registry

func init() {
	Registry = registrar.New("hackerrank")
}
