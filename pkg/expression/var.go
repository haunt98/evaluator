package expression

type VarExpression struct {
	Value string
}

func (expr *VarExpression) String() string {
	return "$" + expr.Value
}

func (expr *VarExpression) Accept(v Visitor) (interface{}, error) {
	return v.VisitVar(expr)
}
