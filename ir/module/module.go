package module

import (
	"fmt"

	"github.com/h8liu/leaf/scope"
	"github.com/h8liu/leaf/symbol"
)

type Module struct {
	Name string
	*scope.Scope
}

func New(name string) *Module {
	ret := new(Module)
	ret.Name = name
	ret.Scope = scope.New()

	return ret
}

func (m *Module) PrintDecls() {
	syms := m.Scope.List()

	for _, s := range syms {
		fmt.Println(symbol.String(s))
	}
}
