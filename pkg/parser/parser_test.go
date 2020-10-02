package parser

import (
	"testing"

	"github.com/haunt98/evaluator/pkg/expression"
	"github.com/haunt98/evaluator/pkg/token"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name     string
	input    string
	wantExpr expression.Expression
	wantErr  error
}

func generateTestCaseLiteral() []testCase {
	return []testCase{
		{
			name:  "bool",
			input: "true",
			wantExpr: expression.BoolLiteral{
				Value: true,
			},
			wantErr: nil,
		},
		{
			name:  "bool",
			input: "false",
			wantExpr: expression.BoolLiteral{
				Value: false,
			},
			wantErr: nil,
		},
		{
			name:  "int",
			input: "1",
			wantExpr: expression.IntLiteral{
				Value: 1,
			},
			wantErr: nil,
		},
		{
			name:  "string",
			input: `"a"`,
			wantExpr: expression.StringLiteral{
				Value: "a",
			},
			wantErr: nil,
		},
	}
}

func generateTestCaseVar() []testCase {
	return []testCase{
		{
			name:  "var",
			input: "$x",
			wantExpr: expression.VarExpression{
				Value: "x",
			},
			wantErr: nil,
		},
		{
			name:  "var",
			input: "$1",
			wantExpr: expression.VarExpression{
				Value: "1",
			},
			wantErr: nil,
		},
		{
			name:  "var",
			input: "$_",
			wantExpr: expression.VarExpression{
				Value: "_",
			},
			wantErr: nil,
		},
	}
}

func generateTestCaseUnary() []testCase {
	return []testCase{
		{
			name:  "not",
			input: "!true",
			wantExpr: expression.UnaryExpression{
				Operator: token.Not,
				Child: expression.BoolLiteral{
					Value: true,
				},
			},
			wantErr: nil,
		},
	}
}

func generateTestCaseParenthesis() []testCase {
	return []testCase{
		{
			name:  "parenthesis",
			input: "(true)",
			wantExpr: expression.ParenthesisExpression{
				Child: expression.BoolLiteral{
					Value: true,
				},
			},
			wantErr: nil,
		},
		{
			name:  "parenthesis",
			input: "(1)",
			wantExpr: expression.ParenthesisExpression{
				Child: expression.IntLiteral{
					Value: 1,
				},
			},
			wantErr: nil,
		},
		{
			name:  "parenthesis",
			input: `("a")`,
			wantExpr: expression.ParenthesisExpression{
				Child: expression.StringLiteral{
					Value: "a",
				},
			},
			wantErr: nil,
		},
	}
}

func generateTestCaseArray() []testCase {
	return []testCase{
		{
			name:  "array",
			input: "[true]",
			wantExpr: expression.ArrayExpression{
				Children: []expression.Expression{
					expression.BoolLiteral{
						Value: true,
					},
				},
			},
			wantErr: nil,
		},
		{
			name:  "array",
			input: "[true, false]",
			wantExpr: expression.ArrayExpression{
				Children: []expression.Expression{
					expression.BoolLiteral{
						Value: true,
					},
					expression.BoolLiteral{
						Value: false,
					},
				},
			},
			wantErr: nil,
		},
		{
			name:  "array",
			input: `[true, 1, "a"]`,
			wantExpr: expression.ArrayExpression{
				Children: []expression.Expression{
					expression.BoolLiteral{
						Value: true,
					},
					expression.IntLiteral{
						Value: 1,
					},
					expression.StringLiteral{
						Value: "a",
					},
				},
			},
			wantErr: nil,
		},
	}
}

// TODO: all binary
func generateTestCaseBinary() []testCase {
	return []testCase{
		{
			name:  "binary or",
			input: "true or false",
			wantExpr: expression.BinaryExpression{
				Operator: token.Or,
				Left: expression.BoolLiteral{
					Value: true,
				},
				Right: expression.BoolLiteral{
					Value: false,
				},
			},
			wantErr: nil,
		},
	}
}

func generateTestCaseComplex() []testCase {
	return []testCase{
		{
			name:  "complex",
			input: "$x or $y or $z",
			wantExpr: expression.BinaryExpression{
				Operator: token.Or,
				Left: expression.BinaryExpression{
					Operator: token.Or,
					Left: expression.VarExpression{
						Value: "x",
					},
					Right: expression.VarExpression{
						Value: "y",
					},
				},
				Right: expression.VarExpression{
					Value: "z",
				},
			},
			wantErr: nil,
		},
		{
			name:  "complex",
			input: "$x or $y and $z",
			wantExpr: expression.BinaryExpression{
				Operator: token.Or,
				Left: expression.VarExpression{
					Value: "x",
				},
				Right: expression.BinaryExpression{
					Operator: token.And,
					Left: expression.VarExpression{
						Value: "y",
					},
					Right: expression.VarExpression{
						Value: "z",
					},
				},
			},
			wantErr: nil,
		},
		{
			name:  "complex",
			input: "($x or $y) and $z",
			wantExpr: expression.BinaryExpression{
				Operator: token.And,
				Left: expression.ParenthesisExpression{
					Child: expression.BinaryExpression{
						Operator: token.Or,
						Left: expression.VarExpression{
							Value: "x",
						},
						Right: expression.VarExpression{
							Value: "y",
						},
					},
				},
				Right: expression.VarExpression{
					Value: "z",
				},
			},
			wantErr: nil,
		},
	}
}

func TestParser_Parse(t *testing.T) {
	var tests []testCase
	tests = append(tests, generateTestCaseLiteral()...)
	tests = append(tests, generateTestCaseVar()...)
	tests = append(tests, generateTestCaseUnary()...)
	tests = append(tests, generateTestCaseParenthesis()...)
	tests = append(tests, generateTestCaseArray()...)
	tests = append(tests, generateTestCaseBinary()...)
	tests = append(tests, generateTestCaseComplex()...)

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			p := NewParser(tc.input)

			gotExpr, gotErr := p.Parse()

			assert.Equal(t, tc.wantErr, gotErr)
			assert.Equal(t, tc.wantExpr, gotExpr)
		})
	}
}

//goos: linux
//goarch: amd64
//pkg: evaluator/pkg/parser
//BenchmarkParser_Parse
//BenchmarkParser_Parse-4   	  172706	      6449 ns/op
func BenchmarkParser_Parse(b *testing.B) {
	input := `!($x == true or $y != 1) and $z == "a" or $t in [true, 1, "a"]`

	var expr expression.Expression
	var err error
	for n := 0; n < b.N; n++ {
		p := NewParser(input)
		expr, err = p.Parse()
	}

	_ = err
	_ = expr
}
