package runner

import (
	"fmt"
)

var registry = map[string]func(){}

func Register(name string, fn func()) {
	registry[name] = fn
}

func Execute(name string) error {
	fn, ok := registry[name]
	if !ok {
		return fmt.Errorf("unknown exercise: %s, cmd: prep", name)
	}
	fn()
	return nil
}

func List() []string {
	keys := make([]string, 0, len(registry))
	for k := range registry {
		keys = append(keys, k)
	}
	return keys
}
