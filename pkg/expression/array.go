package expression

type ArrayExpression struct {
	Children []Expression
}

func (expr *ArrayExpression) Accept(v Visitor) (interface{}, error) {
	return v.VisitArray(expr)
}
