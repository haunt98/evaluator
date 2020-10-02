package token

type Token int

const (
	Illegal Token = iota
	EOF

	Ident
	Bool
	Int
	String
	Var

	Or
	And
	Equal
	NotEqual
	Less
	LessOrEqual
	Greater
	GreaterOrEqual
	In
	NotIn
	Not

	OpenParenthesis
	CloseParenthesis
	OpenSquareBracket
	CloseSquareBracket
	Comma
)

const (
	LowestLevel = 0
	firstLevel  = 1
	secondLevel = 2
	thirdLevel  = 3
	fourthLevel = 4
)

func (tok Token) Precedence() int {
	switch tok {
	case Or:
		return firstLevel
	case And:
		return secondLevel
	case Equal, NotEqual, Less, LessOrEqual, Greater, GreaterOrEqual, In, NotIn:
		return thirdLevel
	case Not:
		return fourthLevel
	default:
		return LowestLevel
	}
}
