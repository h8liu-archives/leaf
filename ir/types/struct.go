package types

import (
	"github.com/h8liu/leaf/symbol"
)

type Struct struct {
	Name string
	Size uint32

	Fields []*Field
}

type Field struct {
	Name   string
	Type   Type
	Offset uint32
}

func (s *Struct) Ident() string {
	return s.Name
}

func (s *Struct) Type() symbol.Type {
	return symbol.Struct
}

func NewStruct(name string) *Struct {
	ret := new(Struct)
	ret.Name = name
	return ret
}
