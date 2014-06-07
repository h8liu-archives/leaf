package main

import (
	"github.com/h8liu/leaf/ir/module"
	"github.com/h8liu/leaf/pass/decls"
)

func main() {
	mod := module.New("main")
	decls.Scan(example, mod)
	mod.PrintDecls()
}
