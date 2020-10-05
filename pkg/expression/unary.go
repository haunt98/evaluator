package expression

import "github.com/haunt98/evaluator/pkg/token"

type UnaryExpression struct {
	Operator token.Token
	Child    Expression
}

func (expr *UnaryExpression) String() string {
	return expr.Operator.String() + expr.Child.String()
}

func (expr *UnaryExpression) Accept(v Visitor) (interface{}, error) {
	return v.VisitUnary(expr)
}
