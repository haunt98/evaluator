package expression

import (
	"strings"
)

var _ Expression = (*ArrayExpression)(nil)

type ArrayExpression struct {
	Children []Expression
}

func NewArrayExpression(children ...Expression) *ArrayExpression {
	if len(children) == 0 {
		return &ArrayExpression{
			Children: []Expression{},
		}
	}

	return &ArrayExpression{
		Children: children,
	}
}

func (expr *ArrayExpression) String() string {
	childrenRepresent := make([]string, len(expr.Children))
	for i, child := range expr.Children {
		childrenRepresent[i] = child.String()
	}

	return "[" + strings.Join(childrenRepresent, " ,") + "]"
}

func (expr *ArrayExpression) Accept(v Visitor) (Expression, error) {
	return v.VisitLiteral(expr)
}
