package registrar

import (
	"fmt"
)

type Registry struct {
	cmd   string
	store map[string]func()
	meta  map[string]any
}

func New(cmd string) *Registry {
	return &Registry{
		cmd:   cmd,
		store: map[string]func(){},
		meta:  map[string]any{},
	}
}

func (reg *Registry) Register(args string, fn func(), meta any) error {
	_, ok := (*reg).store[args]
	if ok {
		return fmt.Errorf("exercise: %s already exist, cmd: %s", args, (*reg).cmd)
	}
	(*reg).store[args] = fn
	if meta != nil {
		(*reg).meta[fmt.Sprintf("%s_meta", args)] = meta
	}
	return nil
}

func (reg *Registry) Execute(args string) error {
	fn, ok := (*reg).store[args]
	if !ok {
		return fmt.Errorf("unknown exercise: %s, cmd: %s", args, (*reg).cmd)
	}
	fn()
	return nil
}

func (reg *Registry) Meta(args string) error {
	value, ok := (*reg).meta[args]
	if !ok {
		return fmt.Errorf("unknown exercise_meta: %s, cmd: %s", args, (*reg).cmd)
	}
	fmt.Println(value)
	return nil
}

func (reg *Registry) List() []string {
	regStore := (*reg).store
	keys := make([]string, 0, len(regStore))
	for k := range regStore {
		keys = append(keys, k)
	}
	return keys
}
