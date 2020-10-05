package visitor

import (
	"fmt"

	"github.com/haunt98/evaluator/pkg/expression"
	"github.com/haunt98/evaluator/pkg/token"
)

type visitor struct {
	args map[string]interface{}
}

func NewVisitor(args map[string]interface{}) expression.Visitor {
	return &visitor{
		args: args,
	}
}

func (v visitor) Visit(expr expression.Expression) (expression.Expression, error) {
	return expr.Accept(v)
}

func (v visitor) VisitBool(lit *expression.BoolLiteral) (expression.Expression, error) {
	return lit, nil
}

func (v visitor) VisitInt(lit *expression.IntLiteral) (expression.Expression, error) {
	return lit, nil
}

func (v visitor) VisitString(lit *expression.StringLiteral) (expression.Expression, error) {
	return lit, nil
}

func (v visitor) VisitVar(expr *expression.VarExpression) (expression.Expression, error) {
	value, ok := v.args[expr.Value]
	if !ok {
		return nil, fmt.Errorf("args missing %s", expr.Value)
	}

	switch v := value.(type) {
	case bool:
		return &expression.BoolLiteral{
			Value: v,
		}, nil
	case int:
		return &expression.IntLiteral{
			Value: int64(v),
		}, nil
	case string:
		return expression.NewStringLiteral(v), nil
	// TODO: add more types
	default:
		return nil, fmt.Errorf("not implement var type %T", v)
	}
}

func (v visitor) VisitParenthesis(expr *expression.ParenthesisExpression) (expression.Expression, error) {
	return v.Visit(expr.Child)
}

func (v visitor) VisitArray(expr *expression.ArrayExpression) (expression.Expression, error) {
	return expr, nil
}

func (v visitor) VisitUnary(expr *expression.UnaryExpression) (expression.Expression, error) {
	switch expr.Operator {
	case token.Not:
		return v.visitNot(expr)
	default:
		return nil, fmt.Errorf("not implement visit unary")
	}
}

func (v visitor) VisitBinary(expr *expression.BinaryExpression) (expression.Expression, error) {
	switch expr.Operator {
	case token.Or:
		return v.visitOr(expr)
	case token.And:
		return v.visitAnd(expr)
	case token.Equal:
		return v.visitEqual(expr)
	case token.NotEqual:
		return v.visitNotEqual(expr)
	default:
		return nil, fmt.Errorf("not implement visit binary")
	}
}
