package visitor

import (
	"fmt"

	"github.com/haunt98/evaluator/pkg/expression"
)

func (v *visitor) visitOr(expr *expression.BinaryExpression) (expression.Expression, error) {
	left, err := v.Visit(expr.Left)
	if err != nil {
		return nil, err
	}

	leftLit, ok := left.(*expression.BoolLiteral)
	if !ok {
		return nil, fmt.Errorf("expect bool literal got %s", left)
	}

	// true or any -> true
	if leftLit.Value {
		return leftLit, nil
	}

	right, err := v.Visit(expr.Right)
	if err != nil {
		return nil, err
	}

	rightLit, ok := right.(*expression.BoolLiteral)
	if !ok {
		return nil, fmt.Errorf("expect bool literal got %s", right)
	}

	return rightLit, nil
}

func (v *visitor) visitAnd(expr *expression.BinaryExpression) (expression.Expression, error) {
	left, err := v.Visit(expr.Left)
	if err != nil {
		return nil, err
	}

	leftLit, ok := left.(*expression.BoolLiteral)
	if !ok {
		return nil, fmt.Errorf("expect bool literal got %s", left)
	}

	// false and any -> false
	if !leftLit.Value {
		return leftLit, nil
	}

	right, err := v.Visit(expr.Right)
	if err != nil {
		return nil, err
	}

	rightLit, ok := right.(*expression.BoolLiteral)
	if !ok {
		return nil, fmt.Errorf("expect bool literal got %s", right)
	}

	return rightLit, nil
}

func (v *visitor) visitEqual(expr *expression.BinaryExpression) (expression.Expression, error) {
	left, err := v.Visit(expr.Left)
	if err != nil {
		return nil, err
	}

	right, err := v.Visit(expr.Right)
	if err != nil {
		return nil, err
	}

	// TODO: this is just wrong, need rewrite
	return &expression.BoolLiteral{
		Value: left == right,
	}, nil
}

func (v *visitor) visitNotEqual(expr *expression.BinaryExpression) (expression.Expression, error) {
	equalExpr, err := v.visitEqual(expr)
	if err != nil {
		return nil, err
	}

	equalLit, ok := equalExpr.(*expression.BoolLiteral)
	if !ok {
		return nil, fmt.Errorf("export bool literal got %s", equalExpr)
	}

	return &expression.BoolLiteral{
		Value: !equalLit.Value,
	}, nil
}
