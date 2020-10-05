package visitor

import (
	"fmt"

	"github.com/haunt98/evaluator/pkg/expression"
)

func (v *visitor) visitNot(expr *expression.UnaryExpression) (interface{}, error) {
	rawChildResult, err := v.Visit(expr.Child)
	if err != nil {
		return nil, err
	}

	childResult, ok := rawChildResult.(bool)
	if !ok {
		return nil, fmt.Errorf("expect bool")
	}

	return !childResult, nil
}
