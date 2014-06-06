package ast

/*

import "fmt"

const msg = "Hello world"

func main() {
	fmt.Println(msg)
	fmt.Println(fabo(5))
}

func fabo(i int) {
	if i <= 1 {
		return 1
	}
	return fabo(i-1) + fabo(i-2)
}

*/

var example = &Package{
	Imports: []*Import{
		{"fmt", "fmt"},
	},

	Decls: []Decl{
		&ConstDecl{
			Name: "msg",
			Type: &PtrType{&BasicType{Char}},
			Expr: &StringLit{"Hello World"},
		},
		&FuncDecl{
			Name: "main",
			Body: &BlockStmt{
				[]Stmt{
					&CallStmt{
						&CallExpr{
							Func: &SelectExpr{
								&Ident{"fmt"},
								"Println",
							},
							Args: []Expr{
								&Ident{"msg"},
							},
						},
					},
					&CallStmt{
						&CallExpr{
							Func: &SelectExpr{
								&Ident{"fmt"},
								"Println",
							},
							Args: []Expr{
								&CallExpr{
									Func: &Ident{"fabo"},
									Args: []Expr{&IntLit{5}},
								},
							},
						},
					},
				},
			},
		},
		&FuncDecl{
			Name: "fabo",
			Paras: []*Para{
				{"i", &BasicType{Int}},
			},
			Body: &BlockStmt{
				[]Stmt{
					&IfStmt{
						Cond: &BinaryOp{
							Op:    "<=",
							Left:  &Ident{"i"},
							Right: &IntLit{1},
						},
						Then: &BlockStmt{
							[]Stmt{
								&ReturnStmt{&IntLit{1}},
							},
						},
					},
					&ReturnStmt{
						&BinaryOp{
							Op: "+",
							Left: &CallExpr{
								Func: &Ident{"fabo"},
								Args: []Expr{
									&BinaryOp{
										Op:    "-",
										Left:  &Ident{"i"},
										Right: &IntLit{1},
									},
								},
							},
							Right: &CallExpr{
								Func: &Ident{"fabo"},
								Args: []Expr{
									&BinaryOp{
										Op:    "-",
										Left:  &Ident{"i "},
										Right: &IntLit{2},
									},
								},
							},
						},
					},
				},
			},
		},
	},
}
