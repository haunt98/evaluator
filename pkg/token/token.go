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

var (
	represents = map[Token]string{
		Illegal:            "Illegal",
		EOF:                "EOF",
		Ident:              "Ident",
		Bool:               "Bool",
		Int:                "Int",
		String:             "String",
		Var:                "Var",
		Or:                 "Or",
		And:                "And",
		Equal:              "==",
		NotEqual:           "!=",
		Less:               "<",
		LessOrEqual:        "<=",
		Greater:            ">",
		GreaterOrEqual:     ">=",
		In:                 "In",
		NotIn:              "NotIn",
		Not:                "!",
		OpenParenthesis:    "(",
		CloseParenthesis:   ")",
		OpenSquareBracket:  "[",
		CloseSquareBracket: "]",
		Comma:              ",",
	}

	precedences = map[Token]int{
		Or:             firstLevel,
		And:            secondLevel,
		Equal:          thirdLevel,
		NotEqual:       thirdLevel,
		Less:           thirdLevel,
		LessOrEqual:    thirdLevel,
		Greater:        thirdLevel,
		GreaterOrEqual: thirdLevel,
		In:             thirdLevel,
		NotIn:          thirdLevel,
		Not:            fourthLevel,
	}
)

func (tok Token) String() string {
	represent, ok := represents[tok]
	if !ok {
		return "Unknown"
	}

	return represent
}

func (tok Token) Precedence() int {
	precedence, ok := precedences[tok]
	if !ok {
		return LowestLevel
	}

	return precedence
}
