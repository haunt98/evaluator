package expression

import "strconv"

type IntLiteral struct {
	Value int64
}

func (lit *IntLiteral) String() string {
	return strconv.FormatInt(lit.Value, 10)
}

func (lit *IntLiteral) Accept(v Visitor) (interface{}, error) {
	return v.VisitInt(lit)
}
