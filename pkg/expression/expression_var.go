package expression

type VarExpression struct {
	Value string
}

func (expr VarExpression) Accept(v Visitor) (interface{}, error) {
	return v.VisitVar(expr)
}
