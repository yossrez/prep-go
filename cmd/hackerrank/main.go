package main

import (
	"fmt"
	"os"

	_ "github.com/yossrez/prep-go/internal/hackerrank/certify/basics/modstrings"
	_ "github.com/yossrez/prep-go/internal/hackerrank/certify/basics/parity"
	_ "github.com/yossrez/prep-go/internal/hackerrank/easy/diagonaldiff"
	_ "github.com/yossrez/prep-go/internal/hackerrank/easy/fizzbuzz"
	_ "github.com/yossrez/prep-go/internal/hackerrank/easy/gameofthrones"
	_ "github.com/yossrez/prep-go/internal/hackerrank/easy/icecreamparlor"
	_ "github.com/yossrez/prep-go/internal/hackerrank/easy/pickingnumbers"
	_ "github.com/yossrez/prep-go/internal/hackerrank/easy/savetheprisoner"
	_ "github.com/yossrez/prep-go/internal/hackerrank/easy/solvemefirst"
	_ "github.com/yossrez/prep-go/internal/hackerrank/easy/superreducedstring"
	"github.com/yossrez/prep-go/internal/hackerrank/runner"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(runner.List())
		return
	}
	err := runner.Execute(os.Args[1])
	if err != nil {
		panic(err)
	}
}
