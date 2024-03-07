package typescript

import (
	"fmt"
	"strings"
)

type Import struct {
	From    string
	Default string
	Named   []string
}

func (i *Import) String() string {
	if len(i.Named) == 0 {
		return fmt.Sprintf("import %s from \"%s\";\n", i.Default, i.From)
	}
	if i.Default == "" {
		return fmt.Sprintf("import { %s } from \"%s\";\n", strings.Join(i.Named, ", "), i.From)
	}
	return fmt.Sprintf("import %s, { %s } from \"%s\";\n", i.Default, strings.Join(i.Named, ", "), i.From)
}
