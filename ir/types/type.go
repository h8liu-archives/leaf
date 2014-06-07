package types

type Type interface {
	Size() uint32
	Align() uint32
	Ident() string
}
