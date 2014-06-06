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

var (
	Int32Type  = newBasic("int32", 4)
	Uint32Type = newBasic("uint32", 4)
	Int8Type   = newBasic("int8", 1)
	Uint8Type  = newBasic("uint8", 1)
)
