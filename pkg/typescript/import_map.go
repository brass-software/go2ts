package typescript

import (
	"fmt"
	"io"
)

type ImportMap struct {
	Imports []*Import
}

func (m *ImportMap) Add(i *Import) error {
	for _, imp := range m.Imports {
		if imp.From == i.From {
			if i.Default != "" {
				if imp.Default != "" && imp.Default != i.Default {
					return fmt.Errorf("default import %s already exists from %s", i.Default, i.From)
				}
				imp.Default = i.Default
			}
			return nil
		}
	}
	m.Imports = append(m.Imports, i)
	return nil
}

func (m *ImportMap) Write(w io.Writer) error {
	for _, imp := range m.Imports {
		_, err := fmt.Fprintf(w, imp.String())
		if err != nil {
			return err
		}
	}
	return nil
}
