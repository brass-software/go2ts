package typescript

import (
	"io"
	"os"
	"path/filepath"
)

type Package struct {
	Imports      *ImportMap
	Declarations []*Declaration
	Exports      *ExportMap
}

func (p *Package) WriteToDir(dir string) error {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}
	f, err := os.Create(filepath.Join(dir, "index.ts"))
	if err != nil {
		return err
	}
	return p.WriteToFile(f)
}

func (p *Package) WriteToFile(w io.Writer) error {
	err := p.Imports.Write(w)
	if err != nil {
		return err
	}
	for _, d := range p.Declarations {
		err = d.Write(w)
		if err != nil {
			return err
		}
	}
	err = p.Exports.Write(w)
	if err != nil {
		return err
	}
	return nil
}
