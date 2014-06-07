// Type representation
package types

import (
	"github.com/h8liu/leaf/symbol"
)

type Basic struct {
	name  string // type name
	size  uint32
	align uint32
}

func newBasic(name string, size uint32) *Basic {
	return &Basic{name, size, size}
}
func (b *Basic) Ident() string     { return b.name }
func (b *Basic) Type() symbol.Type { return symbol.Builtin }
func (b *Basic) Align() uint32     { return b.align }
func (b *Basic) Size() uint32      { return b.size }

var _ Type = new(Basic)

var (
	Int  = newBasic("int", 4)
	Uint = newBasic("uint", 4)
	Byte = newBasic("byte", 1)
	Char = newBasic("char", 1)

	Ptr  = newBasic("ptr", 4)
	Bool = newBasic("bool", 1)

	Basics = []*Basic{Int, Uint, Byte, Char, Ptr, Bool}
)
