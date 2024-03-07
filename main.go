package main

import (
	"flag"
	"go/parser"
	"os"
	"path/filepath"
	"strings"

	"go/token"

	"github.com/brass-software/go2ts/pkg/typescript"
)

func main() {
	flag.Parse()
	dir := flag.Arg(0)
	if dir == "" {
		dir = "."
	}
	err := go2ts(dir)
	if err != nil {
		panic(err)
	}
}

func go2ts(path string) error {
	fi, err := os.Stat(path)
	if err != nil {
		return err
	}
	if !fi.IsDir() {
		fset := token.NewFileSet()
		goFileAST, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if err != nil {
			return err
		}
		importer := typescript.GoImporter{}
		tsFile, err := importer.NewFile(goFileAST)
		if err != nil {
			return err
		}
		goFileName := filepath.Base(path)
		dir := filepath.Dir(path)
		tsFileName := strings.TrimSuffix(goFileName, ".go") + ".ts"
		f, err := os.Create(filepath.Join(dir, tsFileName))
		if err != nil {
			return err
		}
		return tsFile.Write(f)
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	for _, e := range entries {
		err = go2ts(filepath.Join(path, e.Name()))
		if err != nil {
			return err
		}
	}
	return nil
}
