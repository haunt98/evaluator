package expression

import (
	"strconv"
)

var _ Expression = (*BoolLiteral)(nil)

type BoolLiteral struct {
	Value bool
}

func NewBoolLiteral(value bool) *BoolLiteral {
	return &BoolLiteral{
		Value: value,
	}
}

func (lit *BoolLiteral) String() string {
	return strconv.FormatBool(lit.Value)
}

func (lit *BoolLiteral) Accept(v Visitor) (Expression, error) {
	return v.VisitLiteral(lit)
}
