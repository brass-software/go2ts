package typescript

import (
	"fmt"
	"strings"
)

type Expression struct {
	IsLiteral bool
	IsName    bool
	Value     string

	IsCall bool
	Fn     string
	Inputs []*Expression
}

func (expr *Expression) String() string {
	if expr.IsName || expr.IsLiteral {
		return expr.Value
	}
	args := []string{}
	for _, in := range expr.Inputs {
		args = append(args, in.String())
	}
	return fmt.Sprintf("%s(%s)", expr.Fn, strings.Join(args, ", "))
}
