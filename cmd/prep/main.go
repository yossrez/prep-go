package main

import (
	"fmt"
	"os"

	_ "github.com/yossrez/prep-go/internal/roadmapsh/01-introduction/helloworld"
	_ "github.com/yossrez/prep-go/internal/roadmapsh/01-introduction/historyofgo"
	_ "github.com/yossrez/prep-go/internal/roadmapsh/01-introduction/whyusego"
	"github.com/yossrez/prep-go/internal/roadmapsh/runner"
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
