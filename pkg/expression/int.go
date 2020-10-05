package expression

import "strconv"

type IntLiteral struct {
	Value int64
}

func NewIntLiteral(value int64) *IntLiteral {
	return &IntLiteral{
		Value: value,
	}
}

func (lit *IntLiteral) String() string {
	return strconv.FormatInt(lit.Value, 10)
}

func (lit *IntLiteral) Accept(v Visitor) (Expression, error) {
	return v.VisitInt(lit)
}
