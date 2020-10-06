package expression

import (
	"github.com/haunt98/evaluator/pkg/token"
)

type UnaryExpression struct {
	Operator token.Token
	Child    Expression
}

func NewUnaryExpression(op token.Token, child Expression) Expression {
	return &UnaryExpression{
		Operator: op,
		Child:    child,
	}
}

func (expr *UnaryExpression) String() string {
	return expr.Operator.String() + expr.Child.String()
}

func (expr *UnaryExpression) Accept(v Visitor) (Expression, error) {
	return v.VisitUnary(expr)
}
