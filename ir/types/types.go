// Type representation
package types

type Type struct {
	Name    string // type name
	IsBasic bool   // if it is basic type
	Size    uint32
	Align   uint32

	Fields   []*TypeField
	FieldMap map[string]*TypeField
}

type TypeField struct {
	Name   string
	Offset uint32
	Type   *Type
}
