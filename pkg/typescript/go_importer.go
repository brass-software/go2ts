package typescript

import (
	"bytes"
	"go/ast"
	"go/printer"
	"go/token"
)

type GoImporter struct {
	LibMap *LibraryMapping
}

func (importer *GoImporter) NewFile(pkg *ast.File) (*File, error) {
	return nil, nil
}

func (importer *GoImporter) NewPackage(pkg *ast.Package) (*Package, error) {
	p := &Package{}
	for _, obj := range pkg.Scope.Objects {
		switch obj.Kind {
		case ast.Con: // constant
			spec := obj.Decl.(ast.ValueSpec)
			for i, name := range spec.Names {
				d := &Declaration{
					Name:    name.Name,
					IsConst: true,
					Const: &Const{
						Type:  importer.NewType(spec.Type),
						Value: importer.NewExpression(spec.Values[i]),
					},
				}
				p.Declarations = append(p.Declarations, d)
			}
		case ast.Var: // variable
			spec := obj.Decl.(ast.ValueSpec)
			for i, name := range spec.Names {
				d := &Declaration{
					Name:  name.Name,
					IsVar: true,
					Var: &Var{
						Type:  importer.NewType(spec.Type),
						Value: importer.NewExpression(spec.Values[i]),
					},
				}
				p.Declarations = append(p.Declarations, d)
			}
		case ast.Typ: // type
			spec := obj.Decl.(ast.TypeSpec)
			d := &Declaration{
				Name:   spec.Name.Name,
				IsType: true,
				Type:   importer.NewType(spec.Type),
			}
			p.Declarations = append(p.Declarations, d)
		case ast.Fun: // function or method
			//TODO
		}
	}
	return p, nil
}

func (importer *GoImporter) NewExpression(expr any) *Expression {
	switch e := expr.(type) {
	case ast.SelectorExpr:
		return &Expression{
			IsName: true,
			Value:  str(expr),
		}
	case ast.IndexExpr:
		return &Expression{
			IsName: true,
			Value:  str(expr),
		}
	case ast.IndexListExpr:
		return &Expression{
			IsName: true,
			Value:  str(expr),
		}
	case ast.SliceExpr:
		return &Expression{
			IsName: true,
			Value:  str(expr),
		}
	case ast.CallExpr:
		return &Expression{
			IsCall: true,
			Fn:     str(e.Fun),
			Inputs: importer.NewArgList(e.Args),
		}
	case ast.StarExpr:
		return &Expression{
			IsName: true,
			Value:  str(expr),
		}
	case ast.UnaryExpr:
		panic("unsupported type")
	case ast.BinaryExpr:
		panic("unsupported type")
	case ast.KeyValueExpr:
		panic("unsupported type")
	default:
		panic("unsupported type")
	}
}

func (importer *GoImporter) NewType(expr ast.Expr) *Type {
	switch expr.(type) {
	default:
		panic("unknown type: " + str(expr))
	}
}

func str(expr any) string {
	buf := bytes.NewBuffer(nil)
	fset := token.NewFileSet()
	err := printer.Fprint(buf, fset, expr)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func (importer *GoImporter) NewArgList(args []ast.Expr) []*Expression {
	res := []*Expression{}
	for _, arg := range args {
		res = append(res, importer.NewExpression(arg))
	}
	return res
}
