package superreducedstring

import (
	"strings"

	"github.com/yossrez/prep-go/internal/hackerrank"
)

/*
 * Given a string of lowercase characters
 * in range ascii['a',..'z']. Reduce it by
 * deleting a pair of adjacent letters that match.
 *
 * Do this operation until there's no more
 * pair match. Return 'Empty String' if the
 * final string is empty otherwise return
 * whats left.
 */
func superReducedString(s string) string {
	if len(s) == 1 {
		return s
	}

	strBuilder := strings.Builder{}
	strBuilder.WriteString(s)
	l := 0
	r := l + 1
	for {
		str := strBuilder.String()
		lastIdx := len(str) - 1
		strByte := str[l]
		strAdjByte := str[r]
		if strByte != strAdjByte {
			l++
			r++
			if r > lastIdx {
				return str
			}
			continue
		}

		strBuilder.Reset()
		if l > 0 {
			strBuilder.WriteString(str[:l])
		}
		if r < lastIdx {
			strBuilder.WriteString(str[r+1:])
		}

		lenStr := strBuilder.Len()
		if lenStr == 0 {
			break
		}
		if lenStr == 1 {
			return strBuilder.String()
		}

		if r == lastIdx {
			l -= 2
		} else {
			l--
		}
		if l < 0 {
			l = 0
		}
		r = l + 1
	}
	return "Empty String"
}

// main
func Run() {
	// expect 'abd
	_ = superReducedString("aaabccddd")
}

func init() {
	meta := hackerrank.Meta{
		Problem:    "superreducedstring",
		Skills:     hackerrank.ProblemSolvingIntermediate,
		Difficulty: hackerrank.Easy,
		Subdomain:  hackerrank.Strings,
	}
	err := hackerrank.Registry.Register(meta.Problem, Run, meta)
	if err != nil {
		panic(err)
	}
}
