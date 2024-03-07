package typescript

import (
	"fmt"
	"strings"
)

type Type struct {
	IsScalar bool
	Scalar   string
	IsMap    bool
	IsArray  bool
	ElemType *Type
	IsStruct bool
	Fields   []*Field
}

func (t *Type) String(prefix, indent string) string {
	if t.IsScalar {
		return t.Scalar
	}
	if t.IsMap {
		return fmt.Sprintf("{[key: string]: %s}", t.ElemType.String(prefix, indent))
	}
	if t.IsArray {
		return fmt.Sprintf("Array<%s>", t.ElemType.String(prefix, indent))
	}
	if t.IsStruct {
		if len(t.Fields) == 0 {
			return "{}"
		}
		b := strings.Builder{}
		b.WriteString("{\n")
		for _, f := range t.Fields {
			b.WriteString(prefix)
			b.WriteString(indent)
			b.WriteString(f.Name)
			b.WriteString(": ")
			b.WriteString(f.Type.String(prefix+indent, indent))
			b.WriteString(";\n")
		}
		b.WriteString("}\n")
		return b.String()
	}
	panic("unreachable")
}

func (t *Type) IsPrimitive() bool {
	if t.IsScalar {
		if t.Scalar == "string" ||
			t.Scalar == "int" ||
			t.Scalar == "bool" ||
			t.Scalar == "float" {
			return true
		}
	}
	return false
}

// func (t *Type) Imports() []*nodejs.Import {
// 	if t.IsScalar {
// 		if t.IsPrimitive() {
// 			return []*nodejs.Import{}
// 		}
// 		return []*nodejs.Import{
// 			{
// 				From:    "./" + t.Scalar,
// 				Default: t.Scalar,
// 			},
// 		}
// 	}
// 	if t.IsArray || t.IsMap {
// 		return []*nodejs.Import{
// 			{
// 				From:    "./" + t.ElemType,
// 				Default: t.ElemType,
// 			},
// 		}
// 	}
// 	if t.IsStruct {
// 		imports := []*nodejs.Import{}
// 		for _, f := range t.Fields {
// 			if unicode.IsUpper(firstRune(f.Type)) {
// 				imports = append(imports, &nodejs.Import{
// 					From:    "./" + f.Type,
// 					Default: f.Type,
// 				})
// 			}
// 		}
// 		imports = dedup(imports)
// 		return imports
// 	}
// 	panic("unknown type")
// }

// func (t *Type) WriteTS(w io.Writer) error {
// 	tpl := `{{ range .Imports }}{{ .String }}{{ end }}
// {{ .TSString }}`
// 	tmpl, err := template.New("type").Parse(tpl)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return tmpl.Execute(w, t)
// }

// func (t *Type) TSString() string {
// 	if t.IsScalar {
// 		return fmt.Sprintf("type %s = %s;\n\nexport default %s;\n", t.Name, t.Scalar, t.Name)
// 	}
// 	if t.IsArray {
// 		return fmt.Sprintf("type %s = %s[];\n\nexport default %s;\n", t.Name, t.ElemType, t.Name)
// 	}
// 	if t.IsMap {
// 		return fmt.Sprintf("type %s = {[key: string]: %s};\n\nexport default %s;\n", t.Name, t.ElemType, t.Name)
// 	}
// 	if t.IsStruct {
// 		res := "export default interface " + t.Name + " {\n"
// 		for _, f := range t.Fields {
// 			res += fmt.Sprintf("\t%s: %s;\n", f.Name, f.Type)
// 		}
// 		res += "}\n"
// 		return res
// 	}
// 	panic("unknown type")
// }
