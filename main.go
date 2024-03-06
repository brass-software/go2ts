package main

import (
	"fmt"
	"go/parser"
	"io/fs"
	"os"
	"strings"

	"go/token"

	"github.com/brass-software/typescript"
)

func main() {
	inDir := os.Args[1]
	outDir := os.Args[2]
	err := Exec(outDir, inDir)
	if err != nil {
		panic(err)
	}
}

func Exec(outDir, inDir string) error {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, inDir, func(fi fs.FileInfo) bool {
		return strings.HasSuffix(fi.Name(), "_test.go")
	}, parser.ParseComments)
	if err != nil {
		return err
	}
	if len(pkgs) != 1 {
		return fmt.Errorf("expected one package per dir")
	}
	for _, pkg := range pkgs {
		p, err := typescript.NewPackageFromGo(pkg)
		if err != nil {
			return err
		}
		return p.WriteToDir(outDir)
	}
	panic("unreachable")
}
