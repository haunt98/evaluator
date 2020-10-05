package expression

import "github.com/haunt98/evaluator/pkg/token"

type BinaryExpression struct {
	Operator    token.Token
	Left, Right Expression
}

func (expr *BinaryExpression) String() string {
	return expr.Left.String() + expr.Operator.String() + expr.Right.String()
}

func (expr *BinaryExpression) Accept(v Visitor) (interface{}, error) {
	return v.VisitBinary(expr)
}
