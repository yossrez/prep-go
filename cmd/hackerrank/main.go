package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/yossrez/prep-go/internal/hackerrank"
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
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println(hackerrank.Registry.List())
		return
	}
	if strings.Contains(args[1], "_meta") {
		err := hackerrank.Registry.Meta(args[1])
		if err != nil {
			panic(err)
		}
		return
	}
	err := hackerrank.Registry.Execute(args[1])
	if err != nil {
		panic(err)
	}
}
