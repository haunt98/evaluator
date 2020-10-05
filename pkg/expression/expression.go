package expression

type Expression interface {
	String() string
	Accept(v Visitor) (interface{}, error)
}
