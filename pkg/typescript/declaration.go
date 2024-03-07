package typescript

import "io"

type Declaration struct {
	Name    string
	IsConst bool
	Const   *Const
	IsVar   bool
	Var     *Var
	IsFunc  bool
	Func    *Func
	IsType  bool
	Type    *Type
}

func (decl *Declaration) Write(w io.Writer) error {
	return nil
}
