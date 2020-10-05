package expression

type StringLiteral struct {
	Value string
}

func (lit StringLiteral) Accept(v Visitor) (interface{}, error) {
	return v.VisitString(lit)
}
