package expression

import (
	"strconv"
)

type BoolLiteral struct {
	Value bool
}

func NewBoolLiteral(value bool) Expression {
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
