package scope

import (
	"fmt"

	"github.com/h8liu/leaf/ident"
	"github.com/h8liu/leaf/symbol"
)

type Scope struct {
	syms map[string]symbol.Symbol
}

func New() *Scope {
	ret := new(Scope)
	ret.syms = make(map[string]symbol.Symbol)
	return ret
}

func (s *Scope) Define(sym symbol.Symbol) error {
	name := sym.Ident()
	if !ident.IsValid(name) {
		return fmt.Errorf("%q is not a valid identifier", name)
	}
	if s.syms[name] != nil {
		return fmt.Errorf("%q already defined in scope", name)
	}

	s.syms[name] = sym
	return nil
}

func (s *Scope) Query(name string) symbol.Symbol {
	return s.syms[name]
}

func (s *Scope) List() []symbol.Symbol {
	ret := make([]symbol.Symbol, 0, len(s.syms))
	for _, s := range s.syms {
		ret = append(ret, s)
	}

	symbol.Sort(ret)

	return ret
}
