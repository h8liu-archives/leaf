package ast

type (
	Decl interface {}
	Stmt interface {}
	Expr interface {}
	Type interface {}

	Program struct { Imports []*Import; Decls []Decl }

	Import struct { Name, Path string }

	ConstDecl struct { Name string; Type Type; Expr Expr }
	VarDecl struct { Name string; Type Type; Expr Expr }

	StructDecl struct { Name string; Fields []*Field }

	Field struct { Name string; Type Type; Export bool }
	BasicType struct { T int }
	PtrType struct { Type Type }
	StructType struct { Name string }
)



