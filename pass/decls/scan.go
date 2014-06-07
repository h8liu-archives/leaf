package decls

import (
	"github.com/h8liu/leaf/ast"
	"github.com/h8liu/leaf/ir/module"
	"github.com/h8liu/leaf/ir/types"
	"github.com/h8liu/leaf/scope"
)

type scanner struct {
	mod     *module.Module
	context *scope.Context
	e       error
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

func (s *scanner) define(item scope.Item) {
	if s.e != nil {
		return
	}

	s.e = s.mod.Define(item)
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
			// TODO: parse and add the fields
			s.define(t)
		case *ast.FuncDecl:
		case *ast.IfaceDecl:
		default:
			panic("bug")
		}
	}
}
