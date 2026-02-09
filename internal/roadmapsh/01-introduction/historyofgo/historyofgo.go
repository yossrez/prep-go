package historyofgo

import (
	"fmt"

	"github.com/yossrez/prep-go/internal/roadmapsh/runner"
)

func historyOfGo() string {
	summary := `
	* History of Go
	*
	* Created at Google in 2007 by Griesemer, Pike, and Thompson.
	* Announced publicly in 2009, version 1.0 in 2012. Key milestones
	* include modules (Go 1.11) and generics (Go 1.18). Design for
	* large-scale software development combining efficiency and simplicity.
	`
	return summary
}

// main
func Run() {
	summary := historyOfGo()
	fmt.Println(summary)
}

func init() {
	runner.Register("historyofgo", Run)
}
