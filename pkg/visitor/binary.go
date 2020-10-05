package visitor

import (
	"fmt"

	"github.com/haunt98/evaluator/pkg/expression"
)

func (v *visitor) visitOr(expr expression.BinaryExpression) (interface{}, error) {
	rawLeftResult, err := v.Visit(expr.Left)
	if err != nil {
		return nil, err
	}

	leftResult, ok := rawLeftResult.(bool)
	if !ok {
		return nil, fmt.Errorf("expect bool")
	}

	// true or ? = true
	if leftResult {
		return true, nil
	}

	rawRightResult, err := v.Visit(expr.Right)
	if err != nil {
		return nil, err
	}

	rightResult, ok := rawRightResult.(bool)
	if !ok {
		return nil, fmt.Errorf("expect bool")
	}

	return rightResult, nil
}

func (v *visitor) visitAnd(expr expression.BinaryExpression) (interface{}, error) {
	rawLeftResult, err := v.Visit(expr.Left)
	if err != nil {
		return nil, err
	}

	leftResult, ok := rawLeftResult.(bool)
	if !ok {
		return nil, fmt.Errorf("expect bool")
	}

	// false and ? = false
	if !leftResult {
		return false, nil
	}

	rawRightResult, err := v.Visit(expr.Right)
	if err != nil {
		return nil, err
	}

	rightResult, ok := rawRightResult.(bool)
	if !ok {
		return nil, fmt.Errorf("expect bool")
	}

	return rightResult, nil
}

func (v *visitor) visitEqual(expr expression.BinaryExpression) (interface{}, error) {
	leftResult, err := v.Visit(expr.Left)
	if err != nil {
		return nil, err
	}

	rightResult, err := v.Visit(expr.Right)
	if err != nil {
		return nil, err
	}

	return leftResult == rightResult, nil
}

func (v *visitor) visitNotEqual(expr expression.BinaryExpression) (interface{}, error) {
	leftResult, err := v.Visit(expr.Left)
	if err != nil {
		return nil, err
	}

	rightResult, err := v.Visit(expr.Right)
	if err != nil {
		return nil, err
	}

	return leftResult != rightResult, nil
}
