package wus

import (
	"fmt"

	"github.com/yossrez/prep-go/internal/hackerrank"
)

// A weighted string is a string of lowercase English
// letters where each letter has a weight.
//
// Chararter weights are 1 to 26 from 'a' to 'z'
//
// The weight of a string is a sum of the weight
// of its characters.
// A uniform string consists of a single character
// repeated zero of more times.
//
// Exp:
// - 'ccc' and 'a' are uniform strings,
// - 'bcb' and 'cd' are not.
//
// Given a string, 's', let 'U' be the set of weights
// for all possible uniform contiguous substrings
// of string 's'.
//
// There will be 'n' queries to answer where each
// query consists of a single integer.
//
// Create a return array where for each query,
// the value is "Yes" if query[i] elemOf 'U'.
// Otherwise append "No".
func weightedUniformStrings(s string, queries []int32) []string {
	queryResults := []string{}
	setElement := asciiCharsWeight()

	weightMapStr := make(map[int32]struct{})

	var r rune
	var w int32
	var uw int32
	for _, runeChar := range s {
		if runeChar == r {
			uw += w
			weightMapStr[uw] = struct{}{}
			continue
		}

		r = runeChar
		w = setElement[runeChar]
		uw = w
		weightMapStr[uw] = struct{}{}
	}

	for _, query := range queries {
		_, ok := weightMapStr[query]
		if ok {
			queryResults = append(queryResults, "Yes")
		} else {
			queryResults = append(queryResults, "No")
		}
	}

	return queryResults
}

func asciiCharsWeight() map[rune]int32 {
	weightMap := make(map[rune]int32, 26)
	asciiStr := "abcdefghijklmnopqrstuvwxyz"
	for i, rune := range asciiStr {
		weightMap[rune] = int32(i + 1)
	}
	return weightMap
}

// main
func Run() {
	// expect []string{"Yes", "No", "No", "Yes", "No"}
	res := weightedUniformStrings("abbcccdddd", []int32{1, 7, 5, 4, 15})
	fmt.Println(res)
}

func init() {
	meta := hackerrank.Meta{
		Problem:    "wus",
		Skills:     hackerrank.ProblemSolvingIntermediate,
		Difficulty: hackerrank.Easy,
		Subdomain:  hackerrank.Strings,
	}
	err := hackerrank.Registry.Register(meta.Problem, Run, meta)
	if err != nil {
		panic(err)
	}
}
