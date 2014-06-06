package ir

type Type struct {
	Size  uint32
	Align uint32

	Fields   []*TypeField
	FieldMap map[string]*TypeField
}

type TypeField struct {
	Name   string
	Offset uint32
	Type   *Type
}