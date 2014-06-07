package scope

import (
	"fmt"
)

// A scope like a namespace where identifiers are put into it
type Item interface {
	Name() string
}

type Scope struct {
	items map[string]Item
}

func New() *Scope {
	ret := new(Scope)
	ret.items = make(map[string]Item)
	return ret
}

func (s *Scope) Define(i Item) error {
	name := i.Name()
	if s.items[name] != nil {
		return fmt.Errorf("%q already defined in scope", name)
	}

	s.items[name] = i
	return nil
}

func (s *Scope) Query(name string) Item {
	return s.items[name]
}
