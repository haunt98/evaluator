package expression

type Expression interface {
	String() string
	Accept(v Visitor) (Expression, error)
}
