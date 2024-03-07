package typescript

import (
	"fmt"
	"io"
)

type ExportMap struct {
	Default *Expression
	Named   []*NamedExport
}

func (exports *ExportMap) Write(w io.Writer) error {
	if len(exports.Named) > 0 {
		_, err := fmt.Fprintf(w, "export {")
		if err != nil {
			return err
		}
		for _, ex := range exports.Named {
			_, err = fmt.Fprintf(w, "\t%s: %s,\n", ex.Name, ex.Value.String())
			if err != nil {
				return err
			}
		}
		_, err = fmt.Fprintf(w, "}\n")
		if err != nil {
			return err
		}
	}
	if exports.Default != nil {
		_, err := fmt.Fprintf(w, "export default %s;\n", exports.Default.String())
		if err != nil {
			return err
		}
	}
	return nil
}
