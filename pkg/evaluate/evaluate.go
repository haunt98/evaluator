package evaluate

import (
	"fmt"

	"github.com/haunt98/evaluator/pkg/expression"
	"github.com/haunt98/evaluator/pkg/visitor"
)

func Evaluate(expr expression.Expression, args map[string]interface{}) (bool, error) {
	v := visitor.NewVisitor(args)

	expr, err := v.Visit(expr)
	if err != nil {
		return false, err
	}

	lit, ok := expr.(*expression.BoolLiteral)
	if !ok {
		return false, fmt.Errorf("expect bool literal got %s", expr)
	}

	return lit.Value, nil
}
