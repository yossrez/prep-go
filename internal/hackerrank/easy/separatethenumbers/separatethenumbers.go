package separatethenumbers

import (
	"fmt"
	"math"
	"strconv"

	"github.com/yossrez/prep-go/internal/hackerrank"
)

// temp value for testing
var got []string

// A numeric string, 's', is beautiful if it can
// be split into a sequence of two or more positive
// integers, satisfying the following condition:
//   - a[i] - a[i-1] = 1.
//   - No a[i] contain leading 0.
//   - suquence cannot be rearranged.
//
// Given a string of numeric integers
// print 'YES x' if its a beautiful sequence
// , x is the first number,
// in the sequence. Otherwise print 'NO'
func separateTheNumbers(s string) {
	var beautiful bool
	verdict := "NO"

	sLen := len(s)
	if sLen == 1 || s[0] == '0' {
		fmt.Println(verdict)
		appendGotTemp(verdict)
		return
	}

	combination := int(math.Floor(float64(sLen / 2)))
	digit := 1
	i := digit
	x, _ := strconv.Atoi(s[:i])
	expect := x
	for {
		expect += 1
		nextLen := len(fmt.Sprintf("%d", expect))
		if i+nextLen > sLen {
			nextLen -= (i + nextLen) - sLen
		}
		next, _ := strconv.Atoi(s[i : i+nextLen])
		if next == expect {
			i += nextLen
			beautiful = true
			if i > sLen-1 {
				break
			}
			continue
		}

		digit++
		if digit > combination {
			break
		}
		x, _ = strconv.Atoi(s[:digit])
		expect = x
		i = digit
		beautiful = false
	}

	if beautiful {
		verdict = fmt.Sprintf("YES %d", x)
	}
	fmt.Println(verdict)
	appendGotTemp(verdict)
}

func appendGotTemp(val string) {
	got = append(got, val)
}

// main
func Run() {
	// expect YES 1
	separateTheNumbers("1234")
	// expect YES 10
	separateTheNumbers("101112")
	// expect YES 9
	separateTheNumbers("91011")
}

func init() {
	meta := hackerrank.Meta{
		Problem:    "separatethenumbers",
		Skills:     hackerrank.ProblemSolvingBasic,
		Difficulty: hackerrank.Easy,
		Subdomain:  hackerrank.Strings,
	}
	err := hackerrank.Registry.Register(meta.Problem, Run, meta)
	if err != nil {
		panic(err)
	}
}
