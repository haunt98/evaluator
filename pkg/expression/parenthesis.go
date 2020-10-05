package expression

type ParenthesisExpression struct {
	Child Expression
}

func NewParenthesisExpression(child Expression) *ParenthesisExpression {
	return &ParenthesisExpression{
		Child: child,
	}
}

func (expr *ParenthesisExpression) String() string {
	return "(" + expr.Child.String() + ")"
}

func (expr *ParenthesisExpression) Accept(v Visitor) (Expression, error) {
	return v.VisitParenthesis(expr)
}
