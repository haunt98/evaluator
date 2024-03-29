package parser

import (
	"testing"

	"github.com/haunt98/evaluator/expression"
	"github.com/haunt98/evaluator/token"
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
			name:     "bool",
			input:    "true",
			wantExpr: expression.NewBoolLiteral(true),
		},
		{
			name:     "bool",
			input:    "false",
			wantExpr: expression.NewBoolLiteral(false),
		},
		{
			name:     "int",
			input:    "1",
			wantExpr: expression.NewIntLiteral(1),
		},
		{
			name:     "string",
			input:    `"a"`,
			wantExpr: expression.NewStringLiteral("a"),
		},
	}
}

func generateTestCaseVar() []testCase {
	return []testCase{
		{
			name:     "var",
			input:    "$x",
			wantExpr: expression.NewVarExpression("x"),
		},
		{
			name:     "var",
			input:    "$1",
			wantExpr: expression.NewVarExpression("1"),
		},
		{
			name:     "var",
			input:    "$_",
			wantExpr: expression.NewVarExpression("_"),
		},
	}
}

func generateTestCaseUnary() []testCase {
	return []testCase{
		{
			name:     "not",
			input:    "!true",
			wantExpr: expression.NewUnaryExpression(token.Not, expression.NewBoolLiteral(true)),
		},
	}
}

func generateTestCaseParenthesis() []testCase {
	return []testCase{
		{
			name:     "parenthesis single bool",
			input:    "(true)",
			wantExpr: expression.NewBoolLiteral(true),
		},
		{
			name:     "parenthesis single int",
			input:    "(1)",
			wantExpr: expression.NewIntLiteral(1),
		},
		{
			name:     "parenthesis single string",
			input:    `("a")`,
			wantExpr: expression.NewStringLiteral("a"),
		},
		{
			name:  "parenthesis and",
			input: `(true and false)`,
			wantExpr: expression.NewBinaryExpression(token.And,
				expression.NewBoolLiteral(true),
				expression.NewBoolLiteral(false),
			),
		},
	}
}

func generateTestCaseArray() []testCase {
	return []testCase{
		{
			name:     "array empty",
			input:    "[]",
			wantExpr: expression.NewArrayExpression(),
		},
		{
			name:     "array single item",
			input:    "[true]",
			wantExpr: expression.NewArrayExpression(expression.NewBoolLiteral(true)),
		},
		{
			name:  "array multi items",
			input: "[true, false]",
			wantExpr: expression.NewArrayExpression(
				expression.NewBoolLiteral(true),
				expression.NewBoolLiteral(false),
			),
		},
		{
			name:  "array multi complex items",
			input: `[true, 1, "a"]`,
			wantExpr: expression.NewArrayExpression(
				expression.NewBoolLiteral(true),
				expression.NewIntLiteral(1),
				expression.NewStringLiteral("a"),
			),
		},
	}
}

func generateTestCaseBinary() []testCase {
	return []testCase{
		{
			name:  "or",
			input: "true or false",
			wantExpr: expression.NewBinaryExpression(token.Or,
				expression.NewBoolLiteral(true),
				expression.NewBoolLiteral(false),
			),
		},
		{
			name:  "and",
			input: "false and true",
			wantExpr: expression.NewBinaryExpression(token.And,
				expression.NewBoolLiteral(false),
				expression.NewBoolLiteral(true),
			),
		},
		{
			name:  "equal",
			input: "1 == 2",
			wantExpr: expression.NewBinaryExpression(token.Equal,
				expression.NewIntLiteral(1),
				expression.NewIntLiteral(2),
			),
		},
		{
			name:  "not equal",
			input: "1 != 2",
			wantExpr: expression.NewBinaryExpression(token.NotEqual,
				expression.NewIntLiteral(1),
				expression.NewIntLiteral(2),
			),
		},
		{
			name:  "less",
			input: "1 < 2",
			wantExpr: expression.NewBinaryExpression(token.Less,
				expression.NewIntLiteral(1),
				expression.NewIntLiteral(2),
			),
		},
		{
			name:  "less or equal",
			input: "1 <= 2",
			wantExpr: expression.NewBinaryExpression(token.LessOrEqual,
				expression.NewIntLiteral(1),
				expression.NewIntLiteral(2),
			),
		},
		{
			name:  "greater",
			input: "1 > 2",
			wantExpr: expression.NewBinaryExpression(token.Greater,
				expression.NewIntLiteral(1),
				expression.NewIntLiteral(2),
			),
		},
		{
			name:  "greater or equal",
			input: "1 >= 2",
			wantExpr: expression.NewBinaryExpression(token.GreaterOrEqual,
				expression.NewIntLiteral(1),
				expression.NewIntLiteral(2),
			),
		},
		{
			name:  "in",
			input: "1 in [1, 2]",
			wantExpr: expression.NewBinaryExpression(token.In,
				expression.NewIntLiteral(1),
				expression.NewArrayExpression(
					expression.NewIntLiteral(1),
					expression.NewIntLiteral(2),
				),
			),
		},
		{
			name:  "not in",
			input: "1 notin [1, 2]",
			wantExpr: expression.NewBinaryExpression(token.NotIn,
				expression.NewIntLiteral(1),
				expression.NewArrayExpression(
					expression.NewIntLiteral(1),
					expression.NewIntLiteral(2),
				),
			),
		},
	}
}

func generateTestCaseComplex() []testCase {
	return []testCase{
		{
			name:  "complex",
			input: "$x or $y or $z",
			wantExpr: expression.NewBinaryExpression(token.Or,
				expression.NewBinaryExpression(token.Or,
					expression.NewVarExpression("x"),
					expression.NewVarExpression("y"),
				),
				expression.NewVarExpression("z"),
			),
		},
		{
			name:  "complex",
			input: "$x or $y and $z",
			wantExpr: expression.NewBinaryExpression(token.Or,
				expression.NewVarExpression("x"),
				expression.NewBinaryExpression(token.And,
					expression.NewVarExpression("y"),
					expression.NewVarExpression("z"),
				),
			),
		},
		{
			name:  "complex",
			input: "($x or $y) and $z",
			wantExpr: expression.NewBinaryExpression(token.And,
				expression.NewBinaryExpression(token.Or,
					expression.NewVarExpression("x"),
					expression.NewVarExpression("y"),
				),
				expression.NewVarExpression("z"),
			),
		},
		{
			name:  "complex",
			input: "(true and $x) or (($x or false) and true)",
			wantExpr: expression.NewBinaryExpression(token.Or,
				expression.NewBinaryExpression(token.And,
					expression.NewBoolLiteral(true),
					expression.NewVarExpression("x"),
				),
				expression.NewBinaryExpression(token.And,
					expression.NewBinaryExpression(token.Or,
						expression.NewVarExpression("x"),
						expression.NewBoolLiteral(false),
					),
					expression.NewBoolLiteral(true),
				),
			),
		},
	}
}

func TestParserParse(t *testing.T) {
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
			if tc.wantErr != nil {
				return
			}
			assert.Equal(t, tc.wantExpr, gotExpr)
		})
	}
}
