package evaluate

import (
	"fmt"

	"github.com/haunt98/evaluator/pkg/expression"
	"github.com/haunt98/evaluator/pkg/token"
)

type evaluateVisitor struct {
	args map[string]interface{}
}

func NewEvaluateVisitor(args map[string]interface{}) expression.Visitor {
	return &evaluateVisitor{
		args: args,
	}
}

func (v *evaluateVisitor) Visit(expr expression.Expression) (expression.Expression, error) {
	return expr.Accept(v)
}

func (v *evaluateVisitor) VisitLiteral(expr expression.Expression) (expression.Expression, error) {
	return expr, nil
}

func (v *evaluateVisitor) VisitVar(expr *expression.VarExpression) (expression.Expression, error) {
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

func (v *evaluateVisitor) VisitUnary(expr *expression.UnaryExpression) (expression.Expression, error) {
	switch expr.Operator {
	case token.Not:
		return v.visitNot(expr)
	default:
		return nil, fmt.Errorf("not implement visit unary operator %s", expr.Operator)
	}
}

// TODO: add more binary handle
func (v *evaluateVisitor) VisitBinary(expr *expression.BinaryExpression) (expression.Expression, error) {
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
		return nil, fmt.Errorf("not implement visit binary operator %s", expr.Operator)
	}
}
