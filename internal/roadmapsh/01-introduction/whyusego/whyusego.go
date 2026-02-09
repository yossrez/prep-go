package whyusego

import (
	"fmt"

	"github.com/yossrez/prep-go/internal/roadmapsh/runner"
)

func whyUseGo() string {
	summary := `
	* Why use Go?
	*
	* Go offers exceptional performance with single binary
	* deployment, built-in concurrency, fast compilation,
	* and comprehensive standard library. Simple language
	* that's easy to learn and maintain. Excels at web services,
	* microservices, CLI tools, and system software.
	`
	return summary
}

// main
func Run() {
	summary := whyUseGo()
	fmt.Println(summary)
}

func init() {
	runner.Register("whyusego", Run)
}
