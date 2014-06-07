// Type representation
package types

type TypeBox struct {
	Name  string // type name
	Size  uint32
	Align uint32
}

type BasicType struct {
	*TypeBox
}

func newBasic(name string, size uint32) *BasicType {
	return &BasicType{&TypeBox{name, size, size}}
}

func (b *TypeBox) Ident() string {
	return b.Name
}

var (
	Int  = newBasic("int", 4)
	Uint = newBasic("uint", 4)
	Byte = newBasic("byte", 1)
	Char = newBasic("char", 1)

	Ptr = newBasic("ptr", 4)
)
