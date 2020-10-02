package expression

type Expression interface {
	Accept(v Visitor) (interface{}, error)
}
