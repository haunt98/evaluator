package expression

import "github.com/haunt98/evaluator/pkg/token"

type VarExpression struct {
	Value string
}

func NewVarExpression(value string) *VarExpression {
	return &VarExpression{
		Value: value,
	}
}

func (expr *VarExpression) String() string {
	return token.Var.String() + expr.Value
}

func (expr *VarExpression) Accept(v Visitor) (Expression, error) {
	return v.VisitVar(expr)
}
