package expression

type StringLiteral struct {
	Value string
}

func NewStringLiteral(value string) Expression {
	return &StringLiteral{
		Value: value,
	}
}

func (lit *StringLiteral) String() string {
	return `"` + lit.Value + `"`
}

func (lit *StringLiteral) Accept(v Visitor) (Expression, error) {
	return v.VisitLiteral(lit)
}
