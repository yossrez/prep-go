package helloworld

import (
	"fmt"
	"os/user"

	"github.com/yossrez/prep-go/internal/roadmapsh"
)

func helloWorld() string {
	currentUser, err := user.Current()
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("Hello %s!", currentUser.Username)
}

// main
func Run() {
	message := helloWorld()
	fmt.Println(message)
}

func init() {
	err := roadmapsh.Registry.Register("helloworld", Run, nil)
	if err != nil {
		panic(err)
	}
}
