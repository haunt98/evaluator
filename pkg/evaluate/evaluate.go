package evaluate

import (
	"fmt"

	"github.com/haunt98/evaluator/pkg/expression"
	"github.com/haunt98/evaluator/pkg/visitor"
)

func Evaluate(expr expression.Expression, args map[string]interface{}) (bool, error) {
	v := visitor.NewVisitor(args)

	rawResult, err := v.Visit(expr)
	if err != nil {
		return false, err
	}

	result, ok := rawResult.(bool)
	if !ok {
		return false, fmt.Errorf("expect bool")
	}

	return result, nil
}
