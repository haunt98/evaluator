package expression

type IntLiteral struct {
	Value int64
}

func (lit *IntLiteral) Accept(v Visitor) (interface{}, error) {
	return v.VisitInt(lit)
}
