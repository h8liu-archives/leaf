// Type representation
package types

type Basic struct {
	Name  string // type name
	Size  uint32
	Align uint32
}

func newBasic(name string, size uint32) *Basic {
	return &Basic{name, size, size}
}

func (b *Basic) Ident() string {
	return b.Name
}

var (
	Int  = newBasic("int", 4)
	Uint = newBasic("uint", 4)
	Byte = newBasic("byte", 1)
	Char = newBasic("char", 1)

	Ptr  = newBasic("ptr", 4)
	Bool = newBasic("bool", 1)

	Basics = []*Basic{Int, Uint, Byte, Char, Ptr, Bool}
)
