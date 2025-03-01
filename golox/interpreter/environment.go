package interpreter

import "fmt"

// envirionment contains the state of a scope
type environment struct {
	// optimization: use a more performant hashing algorithm
	values map[string]any
}

func (e *environment) get(name string) (any, error) {
	value, ok := e.values[name]
	if !ok {
		// todo: custom error type
		return nil, fmt.Errorf("variable not defined: %s", name)
	}
	return value, nil
}

func (e *environment) set(name string, value any) {
	e.values[name] = value
}
