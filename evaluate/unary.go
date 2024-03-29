package evaluate

import (
	"fmt"

	"github.com/haunt98/evaluator/expression"
)

func (v *visitor) visitNot(expr *expression.UnaryExpression) (expression.Expression, error) {
	child, err := v.Visit(expr.Child)
	if err != nil {
		return nil, err
	}

	childLit, ok := child.(*expression.BoolLiteral)
	if !ok {
		return nil, fmt.Errorf("expect bool")
	}

	return expression.NewBoolLiteral(!childLit.Value), nil
}
