package decls

import (
	"github.com/h8liu/leaf/ast"
	"github.com/h8liu/leaf/ir/types"
	"github.com/h8liu/leaf/scope"
)

func BuiltIn() *scope.Scope {
	ret := scope.New()
	ret.Define(types.Int)
	ret.Define(types.Uint)
	ret.Define(types.Byte)
	ret.Define(types.Char)
	ret.Define(types.Ptr)

	return ret
}

func ScanPackage(p *ast.Package, priv *scope.Scope, pub *scope.Scope) {

}
