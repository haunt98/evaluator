package expression

type ParenthesisExpression struct {
	Child Expression
}

func (expr *ParenthesisExpression) Accept(v Visitor) (interface{}, error) {
	return v.VisitParenthesis(expr)
}
