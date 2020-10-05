package expression

type ParenthesisExpression struct {
	Child Expression
}

func (expr *ParenthesisExpression) String() string {
	return "(" + expr.Child.String() + ")"
}

func (expr *ParenthesisExpression) Accept(v Visitor) (interface{}, error) {
	return v.VisitParenthesis(expr)
}
