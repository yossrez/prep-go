package twocharacters

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/yossrez/prep-go/internal/hackerrank"
)

func alternate(s string) int32 {
	var longestTwoChars int32

	// dedup s
	sMap := make(map[rune]struct{})
	for _, v := range s {
		_, ok := sMap[v]
		if !ok {
			sMap[v] = struct{}{}
		}
	}

	if len(sMap) == 1 {
		return 0
	}

	var i int
	sKeys := make([]rune, len(sMap))
	for key := range sMap {
		sKeys[i] = key
		delete(sMap, key)
		i++
	}
	i = 0

	var twoChars [2]rune
	strBuilder := strings.Builder{}

	sLen := len(s) - 1
	l := 0
	r := l + 1

	var expect int
	var strRune rune
	for l <= len(sKeys)-2 {
		twoChars[0] = sKeys[l]
		twoChars[1] = sKeys[r]

		for i = range sLen {
			strRune = rune(s[i])
			if strRune == twoChars[0] {
				expect = 1
				break
			} else if strRune == twoChars[1] {
				expect = 0
				break
			}
		}
		strBuilder.WriteRune(strRune)

		valid := true
		for i += 1; i <= sLen; i++ {
			strRune = rune(s[i])
			if strRune == twoChars[0] || strRune == twoChars[1] {
				if strRune == twoChars[expect] {
					if expect == 0 {
						expect = 1
					} else {
						expect = 0
					}
					strBuilder.WriteRune(strRune)
				} else {
					valid = false
					break
				}
			}
		}

		if valid && longestTwoChars < int32(strBuilder.Len()) {
			longestTwoChars = int32(strBuilder.Len())
		}
		strBuilder.Reset()
		if r < len(sKeys)-1 {
			r++
		} else {
			l++
			r = l + 1
		}
	}

	return longestTwoChars
}

func Run() {
	// expect 5
	// res := alternate("beabeefeab")
	// fmt.Println(res)

	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create("./cmd/hackerrank/OUTPUT_PATH")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	lTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	l := int32(lTemp)

	s := readLine(reader)

	if l >= 1 && l <= 1000 && l != int32(len(s)) {
		fmt.Fprint(writer, errors.New("not match"))
	}
	result := alternate(s)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	meta := hackerrank.Meta{
		Problem:    "twocharacters",
		Skills:     hackerrank.ProblemSolvingIntermediate,
		Difficulty: hackerrank.Easy,
		Subdomain:  hackerrank.Strings,
	}
	err := hackerrank.Registry.Register(meta.Problem, Run, meta)
	if err != nil {
		panic(err)
	}
}
