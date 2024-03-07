package typescript

type Statement struct {
	IsReturn bool
	Return   *Expression
	IsDecl   bool
	Decl     *Declaration
	IsIf     bool
	If       *IfStatement
}
