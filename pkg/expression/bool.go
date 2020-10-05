package expression

import "strconv"

type BoolLiteral struct {
	Value bool
}

func (lit *BoolLiteral) String() string {
	return strconv.FormatBool(lit.Value)
}

func (lit *BoolLiteral) Accept(v Visitor) (interface{}, error) {
	return v.VisitBool(lit)
}
