package expression

type BoolLiteral struct {
	Value bool
}

func (lit *BoolLiteral) Accept(v Visitor) (interface{}, error) {
	return v.VisitBool(lit)
}
