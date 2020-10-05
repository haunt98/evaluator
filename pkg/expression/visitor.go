package expression

type Visitor interface {
	Visit(expr Expression) (Expression, error)

	VisitLiteral(expr Expression) (Expression, error)
	VisitVar(expr *VarExpression) (Expression, error)
	VisitParenthesis(expr *ParenthesisExpression) (Expression, error)
	VisitUnary(expr *UnaryExpression) (Expression, error)
	VisitBinary(expr *BinaryExpression) (Expression, error)
}
