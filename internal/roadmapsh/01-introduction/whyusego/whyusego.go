package whyusego

import (
	"fmt"

	"github.com/yossrez/prep-go/internal/roadmapsh"
)

func whyUseGo() string {
	summary := `
	* Why use Go?
	*
	* Go offers exceptional performance.
	* single binary deployment.
	* built-in concurrency.
	* fast compilation.
	* comprehensive standard library.
	*
	* Simple language that's easy to learn and maintain.
	* Excels at web services,
	* microservices,
	* CLI tools,
	* and system software.
	`
	return summary
}

// main
func Run() {
	summary := whyUseGo()
	fmt.Println(summary)
}

func init() {
	err := roadmapsh.Registry.Register("whyusego", Run, nil)
	if err != nil {
		panic(err)
	}
}
