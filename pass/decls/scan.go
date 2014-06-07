package decls

import (
	"github.com/h8liu/leaf/ast"
	"github.com/h8liu/leaf/ir/module"
	"github.com/h8liu/leaf/ir/types"
	"github.com/h8liu/leaf/scope"
	"github.com/h8liu/leaf/symbol"
)

type scanner struct {
	mod     *module.Module
	context *scope.Context
	errors  []error
}

func ne(e error) {
	if e != nil {
		panic(e)
	}
}

func baseContext() *scope.Context {
	ret := scope.NewContext()
	for _, t := range types.Basics {
		ne(ret.Define(t))
	}

	return ret
}

func newScanner(m *module.Module) *scanner {
	ret := new(scanner)
	ret.mod = m
	ret.context = baseContext()
	return ret
}

func (s *scanner) record(e error) {
	if e != nil {
		s.errors = append(s.errors, e)
	}
}

func (s *scanner) define(sym symbol.Symbol) {
	s.record(s.mod.Define(sym))
}

func Scan(m *ast.Module, mod *module.Module) {
	// TODO: add import
	s := newScanner(mod)

	for _, d := range m.Decls {
		switch d := d.(type) {
		case *ast.ConstDecl:
		case *ast.VarDecl:
		case *ast.StructDecl:
			t := types.NewStruct(d.Name)
			s.define(t)
		case *ast.FuncDecl:
		case *ast.IfaceDecl:
		default:
			panic("bug")
		}
	}
}
