package roadmapsh

import "github.com/yossrez/prep-go/pkg/registrar"

var Registry *registrar.Registry

func init() {
	Registry = registrar.New("roadmapsh")
}
