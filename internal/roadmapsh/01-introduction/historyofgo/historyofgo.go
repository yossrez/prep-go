package historyofgo

import (
	"fmt"

	"github.com/yossrez/prep-go/internal/roadmapsh"
)

func historyOfGo() string {
	summary := `
	* History of Go
	*
	* 2007 -> Created at Google by Griesemer, Pike, and Thompson.
	* 2009 -> Announced publicly.
	* 2012 -> version 1.0.
	* 2018 -> module (Go 1.11)
	* 2022 -> generics (Go 1.18)
	*
	* Design for large-scale software development
	* combining efficiency and simplicity.
	`
	return summary
}

// main
func Run() {
	summary := historyOfGo()
	fmt.Println(summary)
}

func init() {
	err := roadmapsh.Registry.Register("historyofgo", Run, nil)
	if err != nil {
		panic(err)
	}
}
