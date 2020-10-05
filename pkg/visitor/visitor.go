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

func (v visitor) Visit(expr expression.Expression) (interface{}, error) {
	return expr.Accept(v)
}

func (v visitor) VisitBool(lit expression.BoolLiteral) (interface{}, error) {
	return lit.Value, nil
}

func (v visitor) VisitInt(lit expression.IntLiteral) (interface{}, error) {
	return lit.Value, nil
}

func (v visitor) VisitString(lit expression.StringLiteral) (interface{}, error) {
	return lit.Value, nil
}

func (v visitor) VisitVar(expr expression.VarExpression) (interface{}, error) {
	value, ok := v.args[expr.Value]
	if !ok {
		return nil, fmt.Errorf("args missing %s", expr.Value)
	}

	return value, nil
}

func (v visitor) VisitParenthesis(expr expression.ParenthesisExpression) (interface{}, error) {
	return v.Visit(expr.Child)
}

func (v visitor) VisitArray(expr expression.ArrayExpression) (interface{}, error) {
	results := make([]interface{}, 0, len(expr.Children))

	for _, child := range expr.Children {
		result, err := v.Visit(child)
		if err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	return results, nil
}

func (v visitor) VisitUnary(expr expression.UnaryExpression) (interface{}, error) {
	switch expr.Operator {
	case token.Not:
		return v.visitNot(expr)
	default:
		return nil, fmt.Errorf("not implement visit unary")
	}
}

func (v visitor) VisitBinary(expr expression.BinaryExpression) (interface{}, error) {
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
