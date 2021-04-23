package expression

type StringLiteral struct {
	Value string
}

var _ Expression = (*StringLiteral)(nil)

func NewStringLiteral(value string) *StringLiteral {
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
