package typescript

import (
	"io"
)

type File struct {
	Imports      *ImportMap
	Declarations []*Declaration
	Exports      *ExportMap
}

func (f *File) Write(w io.Writer) error {
	err := f.Imports.Write(w)
	if err != nil {
		return err
	}
	for _, d := range f.Declarations {
		err = d.Write(w)
		if err != nil {
			return err
		}
	}
	err = f.Exports.Write(w)
	if err != nil {
		return err
	}
	return nil
}
