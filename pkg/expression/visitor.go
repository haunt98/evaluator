package expression

type Visitor interface {
	Visit(expr Expression) (Expression, error)

	VisitBool(lit *BoolLiteral) (Expression, error)
	VisitInt(lit *IntLiteral) (Expression, error)
	VisitString(lit *StringLiteral) (Expression, error)

	VisitVar(expr *VarExpression) (Expression, error)
	VisitParenthesis(expr *ParenthesisExpression) (Expression, error)
	VisitArray(expr *ArrayExpression) (Expression, error)
	VisitUnary(expr *UnaryExpression) (Expression, error)
	VisitBinary(expr *BinaryExpression) (Expression, error)
}
