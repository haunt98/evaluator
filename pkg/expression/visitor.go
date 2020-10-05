package expression

type Visitor interface {
	Visit(expr Expression) (interface{}, error)

	VisitBool(lit *BoolLiteral) (interface{}, error)
	VisitInt(lit *IntLiteral) (interface{}, error)
	VisitString(lit *StringLiteral) (interface{}, error)

	VisitVar(expr *VarExpression) (interface{}, error)
	VisitParenthesis(expr *ParenthesisExpression) (interface{}, error)
	VisitArray(expr *ArrayExpression) (interface{}, error)
	VisitUnary(expr *UnaryExpression) (interface{}, error)
	VisitBinary(expr *BinaryExpression) (interface{}, error)
}
