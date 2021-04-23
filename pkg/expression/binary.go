package expression

import (
	"github.com/haunt98/evaluator/pkg/token"
)

var _ Expression = (*BinaryExpression)(nil)

type BinaryExpression struct {
	Operator    token.Token
	Left, Right Expression
}

func NewBinaryExpression(op token.Token, left, right Expression) *BinaryExpression {
	return &BinaryExpression{
		Operator: op,
		Left:     left,
		Right:    right,
	}
}

func (expr *BinaryExpression) String() string {
	return expr.Left.String() + " " + expr.Operator.String() + " " + expr.Right.String()
}

func (expr *BinaryExpression) Accept(v Visitor) (Expression, error) {
	return v.VisitBinary(expr)
}
