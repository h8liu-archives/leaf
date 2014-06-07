package types

import (
	"fmt"

	"github.com/h8liu/leaf/symbol"
)

type Struct struct {
	name string
	size uint32

	Fields   []*Field
	fieldMap map[string]*Field
}

type Field struct {
	Name   string
	Type   Type
	Offset uint32
}

var _ Type = new(Struct)

func (s *Struct) Ident() string     { return s.name }
func (s *Struct) Type() symbol.Type { return symbol.Struct }
func (s *Struct) Size() uint32      { return s.size }
func (s *Struct) Align() uint32 {
	if s.size == 0 {
		return 1
	}
	if s.size == 1 {
		return 1
	}
	if s.size == 2 {
		return 2
	}
	return 4
}

func NewStruct(name string) *Struct {
	ret := new(Struct)
	ret.name = name

	ret.fieldMap = make(map[string]*Field)
	return ret
}

func (s *Struct) AddField(name string, t Type) error {
	if s.fieldMap[name] != nil {
		return fmt.Errorf("field %q redefined", name)
	}

	f := new(Field)
	f.Name = name
	f.Type = t

	s.Fields = append(s.Fields, f)
	s.fieldMap[name] = f

	return nil
}
