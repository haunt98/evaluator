package evaluate

import (
	"fmt"

	"github.com/haunt98/evaluator/pkg/expression"
	"github.com/haunt98/evaluator/pkg/token"
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

	// TODO: handle more types
	switch l := left.(type) {
	case *expression.BoolLiteral:
		switch r := right.(type) {
		case *expression.BoolLiteral:
			return expression.NewBoolLiteral(l.Value == r.Value), nil
		default:
			return nil, fmt.Errorf("expect bool literal got %T", r)
		}
	case *expression.IntLiteral:
		switch r := right.(type) {
		case *expression.IntLiteral:
			return expression.NewBoolLiteral(l.Value == r.Value), nil
		default:
			return nil, fmt.Errorf("expect int literal got %T", r)
		}
	case *expression.StringLiteral:
		switch r := right.(type) {
		case *expression.StringLiteral:
			return expression.NewBoolLiteral(l.Value == r.Value), nil
		default:
			return nil, fmt.Errorf("expect string literal got %T", r)
		}
	default:
		return nil, fmt.Errorf("not implement visit equal %T", l)
	}
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

	return expression.NewBoolLiteral(!equalLit.Value), nil
}

func (v *visitor) visitLess(expr *expression.BinaryExpression) (expression.Expression, error) {
	left, err := v.Visit(expr.Left)
	if err != nil {
		return nil, err
	}

	leftLit, ok := left.(*expression.IntLiteral)
	if !ok {
		return nil, fmt.Errorf("expect int literal got %s", left)
	}

	right, err := v.Visit(expr.Right)
	if err != nil {
		return nil, err
	}

	rightLit, ok := right.(*expression.IntLiteral)
	if !ok {
		return nil, fmt.Errorf("expect int literal got %s", right)
	}

	return expression.NewBoolLiteral(leftLit.Value < rightLit.Value), nil
}

func (v *visitor) visitLessOrEqual(expr *expression.BinaryExpression) (expression.Expression, error) {
	left, err := v.Visit(expr.Left)
	if err != nil {
		return nil, err
	}

	leftLit, ok := left.(*expression.IntLiteral)
	if !ok {
		return nil, fmt.Errorf("expect int literal got %s", left)
	}

	right, err := v.Visit(expr.Right)
	if err != nil {
		return nil, err
	}

	rightLit, ok := right.(*expression.IntLiteral)
	if !ok {
		return nil, fmt.Errorf("expect int literal got %s", right)
	}

	return expression.NewBoolLiteral(leftLit.Value <= rightLit.Value), nil
}

func (v *visitor) visitGreater(expr *expression.BinaryExpression) (expression.Expression, error) {
	left, err := v.Visit(expr.Left)
	if err != nil {
		return nil, err
	}

	leftLit, ok := left.(*expression.IntLiteral)
	if !ok {
		return nil, fmt.Errorf("expect int literal got %s", left)
	}

	right, err := v.Visit(expr.Right)
	if err != nil {
		return nil, err
	}

	rightLit, ok := right.(*expression.IntLiteral)
	if !ok {
		return nil, fmt.Errorf("expect int literal got %s", right)
	}

	return expression.NewBoolLiteral(leftLit.Value > rightLit.Value), nil
}

func (v *visitor) visitGreaterOrEqual(expr *expression.BinaryExpression) (expression.Expression, error) {
	left, err := v.Visit(expr.Left)
	if err != nil {
		return nil, err
	}

	leftLit, ok := left.(*expression.IntLiteral)
	if !ok {
		return nil, fmt.Errorf("expect int literal got %s", left)
	}

	right, err := v.Visit(expr.Right)
	if err != nil {
		return nil, err
	}

	rightLit, ok := right.(*expression.IntLiteral)
	if !ok {
		return nil, fmt.Errorf("expect int literal got %s", right)
	}

	return expression.NewBoolLiteral(leftLit.Value >= rightLit.Value), nil
}

func (v *visitor) visitIn(expr *expression.BinaryExpression) (expression.Expression, error) {
	left, err := v.Visit(expr.Left)
	if err != nil {
		return nil, err
	}

	right, err := v.Visit(expr.Right)
	if err != nil {
		return nil, err
	}

	rightArr, ok := right.(*expression.ArrayExpression)
	if !ok {
		return nil, fmt.Errorf("expect array expression got %s", right)
	}

	// compare left to all children of right
	for _, child := range rightArr.Children {
		equalExpr, err := v.visitEqual(expression.NewBinaryExpression(token.Equal, left, child))
		if err != nil {
			continue
		}

		equalLit, ok := equalExpr.(*expression.BoolLiteral)
		if !ok {
			continue
		}

		if equalLit.Value {
			return expression.NewBoolLiteral(true), nil
		}
	}

	return expression.NewBoolLiteral(false), nil
}

func (v *visitor) visitNotIn(expr *expression.BinaryExpression) (expression.Expression, error) {
	equalExpr, err := v.visitIn(expr)
	if err != nil {
		return nil, err
	}

	equalLit, ok := equalExpr.(*expression.BoolLiteral)
	if !ok {
		return nil, fmt.Errorf("expect bool literal got %s", equalLit)
	}

	return expression.NewBoolLiteral(!equalLit.Value), nil
}
