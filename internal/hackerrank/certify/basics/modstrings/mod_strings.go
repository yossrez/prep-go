package modstrings

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/yossrez/prep-go/internal/hackerrank/runner"
)

/*
 * Complete the 'ModifyString' function below and add imports if needed.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING str as parameter.
 */

func ModifyString(str string) string {
	whitespaceByte := byte(' ')
	builder := strings.Builder{}
	n := len(str) - 1

	for i := n; i >= 0; i-- {
		if str[i] == whitespaceByte {
			continue
		}
		_, err := strconv.Atoi(string(str[i]))
		if err == nil {
			continue
		}
		builder.WriteByte(str[i])
	}

	return builder.String()
}

// main
func Run() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	// stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	stdout, err := os.Create("./cmd/hackerrank/OUTPUT_PATH")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	str := readLine(reader)

	result := ModifyString(str)

	fmt.Fprintf(writer, "%s\n", result)

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

// ran automatically when the package imported
func init() {
	runner.Register(
		runner.HackerRankMeta{
			Problem:    "modstrings",
			Skills:     runner.ProblemSolvingBasic,
			Difficulty: runner.Easy,
			Subdomain:  runner.Strings,
		},
		Run)
}
