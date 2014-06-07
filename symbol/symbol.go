package symbol

import (
	"fmt"
)

type Type int

const (
	Const = iota
	Iface
	Struct
	Func
	Var
	Builtin
	NA
)

var typeNames = map[Type]string{
	Const:  "const",
	Iface:  "iface",
	Struct: "struct",
	Func:   "func",
	Var:    "var",
}

func (t Type) String() string {
	if ret, found := typeNames[t]; found {
		return ret
	}
	return ""
}

type Symbol interface {
	Ident() string
	Type() Type
}

func String(s Symbol) string {
	return fmt.Sprintf("%5s %s", s.Type().String(), s.Ident())
}
