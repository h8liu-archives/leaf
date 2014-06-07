package scope

import (
	"fmt"

	"github.com/h8liu/leaf/ident"
)

// A scope like a namespace where identifiers are put into it
type Item interface {
	Ident() string
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
	name := i.Ident()
	if !ident.IsValid(name) {
		return fmt.Errorf("%q is not a valid identifier", name)
	}
	if s.items[name] != nil {
		return fmt.Errorf("%q already defined in scope", name)
	}

	s.items[name] = i
	return nil
}

func (s *Scope) Query(name string) Item {
	return s.items[name]
}

func (s *Scope) List() []Item {
	ret := make([]Item, 0, len(s.items))
	for _, item := range s.items {
		ret = append(ret, item)
	}

	return ret
}
