package ast

const (
	Uint32 = iota
	Int32
	Uint8
	Int8

	Uint = Uint32
	Int  = Int32
	Ptr  = Uint32
	Char = Int8
	Byte = Uint8
)

type (
	Decl interface{}
	Type interface{}
	Stmt interface{}
	Expr interface{}

	Module struct {
		Imports []*Import
		Decls   []Decl
	}

	Import struct {
		Name, Path string
	}

	ConstDecl struct {
		Name string
		Type Type
		Expr Expr
	}
	VarDecl struct {
		Name string
		Type Type
		Expr Expr
	}

	StructDecl struct {
		Name   string
		Fields []*Field
	}
	Field struct {
		Name   string
		Type   Type
		Export bool
	}

	IfaceDecl struct {
		Name  string
		Funcs []*FuncSig
	}

	// types
	BasicType  struct{ T int }
	PtrType    struct{ Type Type }
	StructType struct{ Pack, Name string }
	ArrayType  struct {
		Size Expr
		Type Type
	}

	FuncSig struct {
		Name  string
		Paras []*Para
		Type  Type
	}

	FuncDecl struct {
		*FuncSig
		Body *BlockStmt
	}

	Para struct {
		Name string
		Type Type
	}

	// statements
	BlockStmt struct{ Stmts []Stmt }
	DeclStmt  struct{ Decl Decl }
	IfStmt    struct {
		Cond Expr
		Then *BlockStmt
		Else *BlockStmt
	}
	ForStmt struct {
		Pre  Stmt
		Cond Expr
		Post Stmt
		Body *BlockStmt
	}
	BreakStmt    struct{ Label string }
	ContStmt     struct{ Label string }
	FallThruStmt struct{}
	ReturnStmt   struct{ Expr Expr }
	GotoStmt     struct{ Label string }
	LabeledStmt  struct {
		Label string
		Stmt  Stmt
	}
	EmptyStmt   struct{}
	IncStmt     struct{ Expr Expr }
	DecStmt     struct{ Expr Expr }
	AssignStmt  struct{ Left, Right Expr }
	VarDeclStmt struct {
		Name string
		Type Type
		Expr Expr
	}
	CallStmt struct{ Call *CallExpr }

	// expressions
	CallExpr struct {
		Func Expr
		Args []Expr
	}
	StringLit struct{ Value string }
	IntLit    struct{ Value int64 }
	Ident     struct{ Name string }
	UnaryOp   struct {
		Op   string
		Expr Expr
	}
	BinaryOp struct {
		Op          string
		Left, Right Expr
	}
	SelectExpr struct {
		Expr Expr
		Name string
	}
	SubExpr  struct{ Expr, Sub Expr }
	TypeExpr struct{ Type Type }
)
