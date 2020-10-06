package scanner

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/haunt98/evaluator/pkg/token"
)

type scannerTestCase struct {
	name  string
	input string
	want  TokenText
}

func generateTestCaseLiteral() []scannerTestCase {
	return []scannerTestCase{
		{
			name:  "ident",
			input: "x",
			want: TokenText{
				Token: token.Ident,
				Text:  "x",
			},
		},
		{
			name:  "bool",
			input: "true",
			want: TokenText{
				Token: token.Bool,
				Text:  "true",
			},
		},
		{
			name:  "bool",
			input: "false",
			want: TokenText{
				Token: token.Bool,
				Text:  "false",
			},
		},
		{
			name:  "int",
			input: "1",
			want: TokenText{
				Token: token.Int,
				Text:  "1",
			},
		},
		{
			name:  "string",
			input: `"a"`,
			want: TokenText{
				Token: token.String,
				Text:  "a",
			},
		},
	}
}

func generateTestCaseVar() []scannerTestCase {
	return []scannerTestCase{
		{
			name:  "var",
			input: "$x",
			want: TokenText{
				Token: token.Var,
				Text:  "x",
			},
		},
		{
			name:  "var",
			input: "$1",
			want: TokenText{
				Token: token.Var,
				Text:  "1",
			},
		},
		{
			name:  "var",
			input: "$_",
			want: TokenText{
				Token: token.Var,
				Text:  "_",
			},
		},
	}
}

func generateTestCaseOperator() []scannerTestCase {
	return []scannerTestCase{
		{
			name:  "equal",
			input: "==",
			want: TokenText{
				Token: token.Equal,
				Text:  "==",
			},
		},
		{
			name:  "not equal",
			input: "!=",
			want: TokenText{
				Token: token.NotEqual,
				Text:  "!=",
			},
		},
		{
			name:  "less",
			input: "<",
			want: TokenText{
				Token: token.Less,
				Text:  "<",
			},
		},
		{
			name:  "less or equal",
			input: "<=",
			want: TokenText{
				Token: token.LessOrEqual,
				Text:  "<=",
			},
		},
		{
			name:  "greater",
			input: ">",
			want: TokenText{
				Token: token.Greater,
				Text:  ">",
			},
		},
		{
			name:  "greater or equal",
			input: ">=",
			want: TokenText{
				Token: token.GreaterOrEqual,
				Text:  ">=",
			},
		},
		{
			name:  "in",
			input: "in",
			want: TokenText{
				Token: token.In,
				Text:  "in",
			},
		},
		{
			name:  "not in",
			input: "notin",
			want: TokenText{
				Token: token.NotIn,
				Text:  "notin",
			},
		},
		{
			name:  "not",
			input: "!",
			want: TokenText{
				Token: token.Not,
				Text:  "!",
			},
		},
		{
			name:  "or",
			input: "or",
			want: TokenText{
				Token: token.Or,
				Text:  "or",
			},
		},
		{
			name:  "and",
			input: "and",
			want: TokenText{
				Token: token.And,
				Text:  "and",
			},
		},
		{
			name:  "in uppercase",
			input: "IN",
			want: TokenText{
				Token: token.In,
				Text:  "in",
			},
		},
		{
			name:  "not in mixed case",
			input: "NotIn",
			want: TokenText{
				Token: token.NotIn,
				Text:  "notin",
			},
		},
	}
}

func generateTestCaseOthers() []scannerTestCase {
	return []scannerTestCase{
		{
			name:  "open parenthesis",
			input: "(",
			want: TokenText{
				Token: token.OpenParenthesis,
				Text:  "(",
			},
		},
		{
			name:  "close parenthesis",
			input: ")",
			want: TokenText{
				Token: token.CloseParenthesis,
				Text:  ")",
			},
		},
		{
			name:  "open square bracket",
			input: "[",
			want: TokenText{
				Token: token.OpenSquareBracket,
				Text:  "[",
			},
		},
		{
			name:  "close square bracket",
			input: "]",
			want: TokenText{
				Token: token.CloseSquareBracket,
				Text:  "]",
			},
		},
		{
			name:  "comma",
			input: ",",
			want: TokenText{
				Token: token.Comma,
				Text:  ",",
			},
		},
		{
			name:  "EOF",
			input: "",
			want: TokenText{
				Token: token.EOF,
				Text:  "",
			},
		},
		{
			name:  "ident lowercase",
			input: "water",
			want: TokenText{
				Token: token.Ident,
				Text:  "water",
			},
		},
		{
			name:  "ident uppercase",
			input: "WATER",
			want: TokenText{
				Token: token.Ident,
				Text:  "WATER",
			},
		},
		{
			name:  "ident mixed case",
			input: "Water",
			want: TokenText{
				Token: token.Ident,
				Text:  "Water",
			},
		},
	}
}

func generateTestCaseIllegal() []scannerTestCase {
	return []scannerTestCase{
		{
			name:  "=!",
			input: "=!",
			want: TokenText{
				Token: token.Illegal,
				Text:  "=!",
			},
		},
	}
}

func TestScanner_Scan(t *testing.T) {
	var tests []scannerTestCase
	tests = append(tests, generateTestCaseLiteral()...)
	tests = append(tests, generateTestCaseVar()...)
	tests = append(tests, generateTestCaseOperator()...)
	tests = append(tests, generateTestCaseOthers()...)
	tests = append(tests, generateTestCaseIllegal()...)

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := NewScanner(strings.NewReader(tc.input))

			got := s.Scan()
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
