package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/yossrez/prep-go/internal/roadmapsh"
	_ "github.com/yossrez/prep-go/internal/roadmapsh/01-introduction/helloworld"
	_ "github.com/yossrez/prep-go/internal/roadmapsh/01-introduction/historyofgo"
	_ "github.com/yossrez/prep-go/internal/roadmapsh/01-introduction/whyusego"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println(roadmapsh.Registry.List())
		return
	}
	if strings.Contains(args[1], "_meta") {
		err := roadmapsh.Registry.Meta(args[1])
		if err != nil {
			fmt.Println(err)
		}
		return
	}
	err := roadmapsh.Registry.Execute(args[1])
	if err != nil {
		panic(err)
	}
}
