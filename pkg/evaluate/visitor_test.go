package evaluate

import (
	"testing"

	"github.com/haunt98/evaluator/pkg/expression"
	"github.com/haunt98/evaluator/pkg/token"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name       string
	inputExpr  expression.Expression
	inputArgs  map[string]interface{}
	wantResult expression.Expression
	wantErr    error
}

func generateTestCaseLiteral() []testCase {
	return []testCase{
		{
			name:       "bool",
			inputExpr:  expression.NewBoolLiteral(true),
			wantResult: expression.NewBoolLiteral(true),
		},
		{
			name:       "int",
			inputExpr:  expression.NewIntLiteral(1),
			wantResult: expression.NewIntLiteral(1),
		},
		{
			name:       "string",
			inputExpr:  expression.NewStringLiteral("a"),
			wantResult: expression.NewStringLiteral("a"),
		},
		{
			name: "array",
			inputExpr: expression.NewArrayExpression(
				expression.NewBoolLiteral(true),
				expression.NewBoolLiteral(false),
			),
			wantResult: expression.NewArrayExpression(
				expression.NewBoolLiteral(true),
				expression.NewBoolLiteral(false),
			),
		},
	}
}

func generateTestCaseVar() []testCase {
	return []testCase{
		{
			name:      "var bool",
			inputExpr: expression.NewVarExpression("x"),
			inputArgs: map[string]interface{}{
				"x": true,
			},
			wantResult: expression.NewBoolLiteral(true),
		},
		{
			name:      "var int",
			inputExpr: expression.NewVarExpression("x"),
			inputArgs: map[string]interface{}{
				"x": 1,
			},
			wantResult: expression.NewIntLiteral(1),
		},
		{
			name:      "var int64",
			inputExpr: expression.NewVarExpression("x"),
			inputArgs: map[string]interface{}{
				"x": int64(1),
			},
			wantResult: expression.NewIntLiteral(1),
		},
		{
			name:      "var string",
			inputExpr: expression.NewVarExpression("x"),
			inputArgs: map[string]interface{}{
				"x": "xxx",
			},
			wantResult: expression.NewStringLiteral("xxx"),
		},
	}
}

func generateTestCaseUnary() []testCase {
	return []testCase{
		{
			name:       "not true",
			inputExpr:  expression.NewUnaryExpression(token.Not, expression.NewBoolLiteral(true)),
			wantResult: expression.NewBoolLiteral(false),
		},
		{
			name:       "not false",
			inputExpr:  expression.NewUnaryExpression(token.Not, expression.NewBoolLiteral(false)),
			wantResult: expression.NewBoolLiteral(true),
		},
		{
			name:      "not var true",
			inputExpr: expression.NewUnaryExpression(token.Not, expression.NewVarExpression("x")),
			inputArgs: map[string]interface{}{
				"x": true,
			},
			wantResult: expression.NewBoolLiteral(false),
		},
		{
			name:      "not var false",
			inputExpr: expression.NewUnaryExpression(token.Not, expression.NewVarExpression("x")),
			inputArgs: map[string]interface{}{
				"x": false,
			},
			wantResult: expression.NewBoolLiteral(true),
		},
	}
}

func generateTestCaseBinary() []testCase {
	return []testCase{
		{
			name: "or",
			inputExpr: expression.NewBinaryExpression(token.Or,
				expression.NewBoolLiteral(true),
				expression.NewBoolLiteral(false),
			),
			wantResult: expression.NewBoolLiteral(true),
		},
		{
			name: "or with args",
			inputExpr: expression.NewBinaryExpression(token.Or,
				expression.NewVarExpression("x"),
				expression.NewBoolLiteral(false),
			),
			inputArgs: map[string]interface{}{
				"x": false,
			},
			wantResult: expression.NewBoolLiteral(false),
		},
		{
			name: "and",
			inputExpr: expression.NewBinaryExpression(token.And,
				expression.NewBoolLiteral(true),
				expression.NewBoolLiteral(false),
			),
			wantResult: expression.NewBoolLiteral(false),
		},
		{
			name: "and with args",
			inputExpr: expression.NewBinaryExpression(token.And,
				expression.NewVarExpression("x"),
				expression.NewBoolLiteral(true),
			),
			inputArgs: map[string]interface{}{
				"x": true,
			},
			wantResult: expression.NewBoolLiteral(true),
		},
		{
			name: "equal bool",
			inputExpr: expression.NewBinaryExpression(token.Equal,
				expression.NewBoolLiteral(true),
				expression.NewBoolLiteral(true),
			),
			wantResult: expression.NewBoolLiteral(true),
		},
		{
			name: "equal bool",
			inputExpr: expression.NewBinaryExpression(token.Equal,
				expression.NewBoolLiteral(true),
				expression.NewBoolLiteral(false),
			),
			wantResult: expression.NewBoolLiteral(false),
		},
		{
			name: "equal int",
			inputExpr: expression.NewBinaryExpression(token.Equal,
				expression.NewIntLiteral(1),
				expression.NewIntLiteral(1),
			),
			wantResult: expression.NewBoolLiteral(true),
		},
		{
			name: "equal int",
			inputExpr: expression.NewBinaryExpression(token.Equal,
				expression.NewIntLiteral(1),
				expression.NewIntLiteral(2),
			),
			wantResult: expression.NewBoolLiteral(false),
		},
		{
			name: "equal string",
			inputExpr: expression.NewBinaryExpression(token.Equal,
				expression.NewStringLiteral("a"),
				expression.NewStringLiteral("a"),
			),
			wantResult: expression.NewBoolLiteral(true),
		},
		{
			name: "equal string",
			inputExpr: expression.NewBinaryExpression(token.Equal,
				expression.NewStringLiteral("a"),
				expression.NewStringLiteral("b"),
			),
			wantResult: expression.NewBoolLiteral(false),
		},
		{
			name: "not equal bool",
			inputExpr: expression.NewBinaryExpression(token.NotEqual,
				expression.NewBoolLiteral(true),
				expression.NewBoolLiteral(true),
			),
			wantResult: expression.NewBoolLiteral(false),
		},
		{
			name: "not equal bool",
			inputExpr: expression.NewBinaryExpression(token.NotEqual,
				expression.NewBoolLiteral(true),
				expression.NewBoolLiteral(false),
			),
			wantResult: expression.NewBoolLiteral(true),
		},
		{
			name: "not equal int",
			inputExpr: expression.NewBinaryExpression(token.NotEqual,
				expression.NewIntLiteral(1),
				expression.NewIntLiteral(1),
			),
			wantResult: expression.NewBoolLiteral(false),
		},
		{
			name: "not equal int",
			inputExpr: expression.NewBinaryExpression(token.NotEqual,
				expression.NewIntLiteral(1),
				expression.NewIntLiteral(2),
			),
			wantResult: expression.NewBoolLiteral(true),
		},
		{
			name: "not equal string",
			inputExpr: expression.NewBinaryExpression(token.NotEqual,
				expression.NewStringLiteral("a"),
				expression.NewStringLiteral("a"),
			),
			wantResult: expression.NewBoolLiteral(false),
		},
		{
			name: "not equal string",
			inputExpr: expression.NewBinaryExpression(token.NotEqual,
				expression.NewStringLiteral("a"),
				expression.NewStringLiteral("b"),
			),
			wantResult: expression.NewBoolLiteral(true),
		},
		{
			name: "less",
			inputExpr: expression.NewBinaryExpression(token.Less,
				expression.NewIntLiteral(1),
				expression.NewIntLiteral(2),
			),
			wantResult: expression.NewBoolLiteral(true),
		},
		{
			name: "less",
			inputExpr: expression.NewBinaryExpression(token.Less,
				expression.NewIntLiteral(1),
				expression.NewIntLiteral(1),
			),
			wantResult: expression.NewBoolLiteral(false),
		},
		{
			name: "less",
			inputExpr: expression.NewBinaryExpression(token.Less,
				expression.NewIntLiteral(1),
				expression.NewIntLiteral(0),
			),
			wantResult: expression.NewBoolLiteral(false),
		},
		{
			name: "less or equal",
			inputExpr: expression.NewBinaryExpression(token.Less,
				expression.NewIntLiteral(1),
				expression.NewIntLiteral(2),
			),
			wantResult: expression.NewBoolLiteral(true),
		},
		{
			name: "less or equal",
			inputExpr: expression.NewBinaryExpression(token.LessOrEqual,
				expression.NewIntLiteral(1),
				expression.NewIntLiteral(1),
			),
			wantResult: expression.NewBoolLiteral(true),
		},
		{
			name: "less or equal",
			inputExpr: expression.NewBinaryExpression(token.LessOrEqual,
				expression.NewIntLiteral(1),
				expression.NewIntLiteral(0),
			),
			inputArgs:  nil,
			wantResult: expression.NewBoolLiteral(false),
			wantErr:    nil,
		},
		{
			name: "greater",
			inputExpr: expression.NewBinaryExpression(token.Greater,
				expression.NewIntLiteral(1),
				expression.NewIntLiteral(2),
			),
			inputArgs:  nil,
			wantResult: expression.NewBoolLiteral(false),
			wantErr:    nil,
		},
		{
			name: "greater",
			inputExpr: expression.NewBinaryExpression(token.Greater,
				expression.NewIntLiteral(1),
				expression.NewIntLiteral(1),
			),
			wantResult: expression.NewBoolLiteral(false),
		},
		{
			name: "greater",
			inputExpr: expression.NewBinaryExpression(token.Greater,
				expression.NewIntLiteral(1),
				expression.NewIntLiteral(0),
			),
			wantResult: expression.NewBoolLiteral(true),
		},
		{
			name: "greater or equal",
			inputExpr: expression.NewBinaryExpression(token.GreaterOrEqual,
				expression.NewIntLiteral(1),
				expression.NewIntLiteral(2),
			),
			wantResult: expression.NewBoolLiteral(false),
		},
		{
			name: "greater or equal",
			inputExpr: expression.NewBinaryExpression(token.GreaterOrEqual,
				expression.NewIntLiteral(1),
				expression.NewIntLiteral(1),
			),
			wantResult: expression.NewBoolLiteral(true),
		},
		{
			name: "greater or equal",
			inputExpr: expression.NewBinaryExpression(token.GreaterOrEqual,
				expression.NewIntLiteral(1),
				expression.NewIntLiteral(0),
			),
			wantResult: expression.NewBoolLiteral(true),
		},
		{
			name: "in",
			inputExpr: expression.NewBinaryExpression(token.In,
				expression.NewBoolLiteral(true),
				expression.NewArrayExpression(
					expression.NewBoolLiteral(true),
					expression.NewIntLiteral(1),
					expression.NewStringLiteral("a"),
				),
			),
			wantResult: expression.NewBoolLiteral(true),
		},
		{
			name: "in",
			inputExpr: expression.NewBinaryExpression(token.In,
				expression.NewBoolLiteral(false),
				expression.NewArrayExpression(
					expression.NewBoolLiteral(true),
					expression.NewIntLiteral(1),
					expression.NewStringLiteral("a"),
				),
			),
			wantResult: expression.NewBoolLiteral(false),
		},
		{
			name: "in",
			inputExpr: expression.NewBinaryExpression(token.In,
				expression.NewStringLiteral("a"),
				expression.NewArrayExpression(
					expression.NewBoolLiteral(true),
					expression.NewIntLiteral(1),
					expression.NewStringLiteral("a"),
				),
			),
			wantResult: expression.NewBoolLiteral(true),
		},
		{
			name: "in",
			inputExpr: expression.NewBinaryExpression(token.In,
				expression.NewStringLiteral("b"),
				expression.NewArrayExpression(
					expression.NewBoolLiteral(true),
					expression.NewIntLiteral(1),
					expression.NewStringLiteral("a"),
				),
			),
			wantResult: expression.NewBoolLiteral(false),
		},
		{
			name: "not in",
			inputExpr: expression.NewBinaryExpression(token.NotIn,
				expression.NewBoolLiteral(true),
				expression.NewArrayExpression(
					expression.NewBoolLiteral(true),
					expression.NewIntLiteral(1),
					expression.NewStringLiteral("a"),
				),
			),
			wantResult: expression.NewBoolLiteral(false),
		},
		{
			name: "not in",
			inputExpr: expression.NewBinaryExpression(token.NotIn,
				expression.NewBoolLiteral(false),
				expression.NewArrayExpression(
					expression.NewBoolLiteral(true),
					expression.NewIntLiteral(1),
					expression.NewStringLiteral("a"),
				),
			),
			wantResult: expression.NewBoolLiteral(true),
		},
		{
			name: "not in",
			inputExpr: expression.NewBinaryExpression(token.NotIn,
				expression.NewStringLiteral("a"),
				expression.NewArrayExpression(
					expression.NewBoolLiteral(true),
					expression.NewIntLiteral(1),
					expression.NewStringLiteral("a"),
				),
			),
			wantResult: expression.NewBoolLiteral(false),
		},
		{
			name: "not in",
			inputExpr: expression.NewBinaryExpression(token.NotIn,
				expression.NewStringLiteral("b"),
				expression.NewArrayExpression(
					expression.NewBoolLiteral(true),
					expression.NewIntLiteral(1),
					expression.NewStringLiteral("a"),
				),
			),
			wantResult: expression.NewBoolLiteral(true),
		},
	}
}

func TestEvaluateVisitorVisit(t *testing.T) {
	var tests []testCase
	tests = append(tests, generateTestCaseLiteral()...)
	tests = append(tests, generateTestCaseVar()...)
	tests = append(tests, generateTestCaseUnary()...)
	tests = append(tests, generateTestCaseBinary()...)

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := NewVisitor(tc.inputArgs)

			gotResult, gotErr := v.Visit(tc.inputExpr)
			assert.Equal(t, tc.wantErr, gotErr)
			if tc.wantErr != nil {
				return
			}
			assert.Equal(t, tc.wantResult, gotResult)
		})
	}
}
