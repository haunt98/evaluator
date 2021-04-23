package evaluate

import (
	"fmt"

	"github.com/haunt98/evaluator/pkg/expression"
	"github.com/haunt98/evaluator/pkg/token"
)

var _ expression.Visitor = (*visitor)(nil)

type visitor struct {
	args map[string]interface{}
}

func NewVisitor(args map[string]interface{}) *visitor {
	return &visitor{
		args: args,
	}
}

func (v *visitor) Visit(expr expression.Expression) (expression.Expression, error) {
	return expr.Accept(v)
}

func (v *visitor) VisitLiteral(expr expression.Expression) (expression.Expression, error) {
	return expr, nil
}

func (v *visitor) VisitVar(expr *expression.VarExpression) (expression.Expression, error) {
	value, ok := v.args[expr.Value]
	if !ok {
		return nil, fmt.Errorf("args missing %s", expr.Value)
	}

	// TODO: add more types
	switch v := value.(type) {
	case bool:
		return expression.NewBoolLiteral(v), nil
	case int:
		return expression.NewIntLiteral(int64(v)), nil
	case int64:
		return expression.NewIntLiteral(v), nil
	case string:
		return expression.NewStringLiteral(v), nil
	default:
		return nil, fmt.Errorf("not implement var type %T", v)
	}
}

func (v *visitor) VisitUnary(expr *expression.UnaryExpression) (expression.Expression, error) {
	switch expr.Operator {
	case token.Not:
		return v.visitNot(expr)
	default:
		return nil, fmt.Errorf("not implement visit unary operator %s", expr.Operator)
	}
}

// TODO: add more binary handle
func (v *visitor) VisitBinary(expr *expression.BinaryExpression) (expression.Expression, error) {
	switch expr.Operator {
	case token.Or:
		return v.visitOr(expr)
	case token.And:
		return v.visitAnd(expr)
	case token.Equal:
		return v.visitEqual(expr)
	case token.NotEqual:
		return v.visitNotEqual(expr)
	case token.Less:
		return v.visitLess(expr)
	case token.LessOrEqual:
		return v.visitLessOrEqual(expr)
	case token.Greater:
		return v.visitGreater(expr)
	case token.GreaterOrEqual:
		return v.visitGreaterOrEqual(expr)
	case token.In:
		return v.visitIn(expr)
	case token.NotIn:
		return v.visitNotIn(expr)
	default:
		return nil, fmt.Errorf("not implement visit binary operator %s", expr.Operator)
	}
}
