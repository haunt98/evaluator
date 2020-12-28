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
			inputArgs:  nil,
			wantResult: expression.NewBoolLiteral(true),
			wantErr:    nil,
		},
		{
			name:       "int",
			inputExpr:  expression.NewIntLiteral(1),
			inputArgs:  nil,
			wantResult: expression.NewIntLiteral(1),
			wantErr:    nil,
		},
		{
			name:       "string",
			inputExpr:  expression.NewStringLiteral("a"),
			inputArgs:  nil,
			wantResult: expression.NewStringLiteral("a"),
			wantErr:    nil,
		},
		{
			name: "array",
			inputExpr: expression.NewArrayExpression(
				expression.NewBoolLiteral(true),
				expression.NewBoolLiteral(false),
			),
			inputArgs: nil,
			wantResult: expression.NewArrayExpression(
				expression.NewBoolLiteral(true),
				expression.NewBoolLiteral(false),
			),
			wantErr: nil,
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
			wantErr:    nil,
		},
		{
			name:      "var int",
			inputExpr: expression.NewVarExpression("x"),
			inputArgs: map[string]interface{}{
				"x": 1,
			},
			wantResult: expression.NewIntLiteral(1),
			wantErr:    nil,
		},
		{
			name:      "var int64",
			inputExpr: expression.NewVarExpression("x"),
			inputArgs: map[string]interface{}{
				"x": int64(1),
			},
			wantResult: expression.NewIntLiteral(1),
			wantErr:    nil,
		},
		{
			name:      "var string",
			inputExpr: expression.NewVarExpression("x"),
			inputArgs: map[string]interface{}{
				"x": "xxx",
			},
			wantResult: expression.NewStringLiteral("xxx"),
			wantErr:    nil,
		},
	}
}

func generateTestCaseUnary() []testCase {
	return []testCase{
		{
			name:       "not true",
			inputExpr:  expression.NewUnaryExpression(token.Not, expression.NewBoolLiteral(true)),
			inputArgs:  nil,
			wantResult: expression.NewBoolLiteral(false),
			wantErr:    nil,
		},
		{
			name:       "not false",
			inputExpr:  expression.NewUnaryExpression(token.Not, expression.NewBoolLiteral(false)),
			inputArgs:  nil,
			wantResult: expression.NewBoolLiteral(true),
			wantErr:    nil,
		},
	}
}

// TODO: add more binary unittest
func generateTestCaseBinary() []testCase {
	return []testCase{
		{
			name: "or",
			inputExpr: expression.NewBinaryExpression(token.Or,
				expression.NewBoolLiteral(true),
				expression.NewBoolLiteral(false),
			),
			inputArgs:  nil,
			wantResult: expression.NewBoolLiteral(true),
			wantErr:    nil,
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
			wantErr:    nil,
		},
		{
			name: "and",
			inputExpr: expression.NewBinaryExpression(token.And,
				expression.NewBoolLiteral(true),
				expression.NewBoolLiteral(false),
			),
			inputArgs:  nil,
			wantResult: expression.NewBoolLiteral(false),
			wantErr:    nil,
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
			wantErr:    nil,
		},
		{
			name: "equal bool",
			inputExpr: expression.NewBinaryExpression(token.Equal,
				expression.NewBoolLiteral(true),
				expression.NewBoolLiteral(true),
			),
			inputArgs:  nil,
			wantResult: expression.NewBoolLiteral(true),
			wantErr:    nil,
		},
		{
			name: "equal bool",
			inputExpr: expression.NewBinaryExpression(token.Equal,
				expression.NewBoolLiteral(true),
				expression.NewBoolLiteral(false),
			),
			inputArgs:  nil,
			wantResult: expression.NewBoolLiteral(false),
			wantErr:    nil,
		},
		{
			name: "equal int",
			inputExpr: expression.NewBinaryExpression(token.Equal,
				expression.NewIntLiteral(1),
				expression.NewIntLiteral(1),
			),
			inputArgs:  nil,
			wantResult: expression.NewBoolLiteral(true),
			wantErr:    nil,
		},
		{
			name: "equal int",
			inputExpr: expression.NewBinaryExpression(token.Equal,
				expression.NewIntLiteral(1),
				expression.NewIntLiteral(2),
			),
			inputArgs:  nil,
			wantResult: expression.NewBoolLiteral(false),
			wantErr:    nil,
		},
		{
			name: "equal string",
			inputExpr: expression.NewBinaryExpression(token.Equal,
				expression.NewStringLiteral("a"),
				expression.NewStringLiteral("a"),
			),
			inputArgs:  nil,
			wantResult: expression.NewBoolLiteral(true),
			wantErr:    nil,
		},
		{
			name: "equal string",
			inputExpr: expression.NewBinaryExpression(token.Equal,
				expression.NewStringLiteral("a"),
				expression.NewStringLiteral("b"),
			),
			inputArgs:  nil,
			wantResult: expression.NewBoolLiteral(false),
			wantErr:    nil,
		},
		{
			name: "not equal bool",
			inputExpr: expression.NewBinaryExpression(token.NotEqual,
				expression.NewBoolLiteral(true),
				expression.NewBoolLiteral(true),
			),
			inputArgs:  nil,
			wantResult: expression.NewBoolLiteral(false),
			wantErr:    nil,
		},
		{
			name: "not equal bool",
			inputExpr: expression.NewBinaryExpression(token.NotEqual,
				expression.NewBoolLiteral(true),
				expression.NewBoolLiteral(false),
			),
			inputArgs:  nil,
			wantResult: expression.NewBoolLiteral(true),
			wantErr:    nil,
		},
		{
			name: "not equal int",
			inputExpr: expression.NewBinaryExpression(token.NotEqual,
				expression.NewIntLiteral(1),
				expression.NewIntLiteral(1),
			),
			inputArgs:  nil,
			wantResult: expression.NewBoolLiteral(false),
			wantErr:    nil,
		},
		{
			name: "not equal int",
			inputExpr: expression.NewBinaryExpression(token.NotEqual,
				expression.NewIntLiteral(1),
				expression.NewIntLiteral(2),
			),
			inputArgs:  nil,
			wantResult: expression.NewBoolLiteral(true),
			wantErr:    nil,
		},
		{
			name: "not equal string",
			inputExpr: expression.NewBinaryExpression(token.NotEqual,
				expression.NewStringLiteral("a"),
				expression.NewStringLiteral("a"),
			),
			inputArgs:  nil,
			wantResult: expression.NewBoolLiteral(false),
			wantErr:    nil,
		},
		{
			name: "not equal string",
			inputExpr: expression.NewBinaryExpression(token.NotEqual,
				expression.NewStringLiteral("a"),
				expression.NewStringLiteral("b"),
			),
			inputArgs:  nil,
			wantResult: expression.NewBoolLiteral(true),
			wantErr:    nil,
		},
		{
			name: "less",
			inputExpr: expression.NewBinaryExpression(token.Less,
				expression.NewIntLiteral(1),
				expression.NewIntLiteral(2),
			),
			inputArgs:  nil,
			wantResult: expression.NewBoolLiteral(true),
			wantErr:    nil,
		},
		{
			name: "less",
			inputExpr: expression.NewBinaryExpression(token.Less,
				expression.NewIntLiteral(1),
				expression.NewIntLiteral(1),
			),
			inputArgs:  nil,
			wantResult: expression.NewBoolLiteral(false),
			wantErr:    nil,
		},
		{
			name: "less",
			inputExpr: expression.NewBinaryExpression(token.Less,
				expression.NewIntLiteral(1),
				expression.NewIntLiteral(0),
			),
			inputArgs:  nil,
			wantResult: expression.NewBoolLiteral(false),
			wantErr:    nil,
		},
		{
			name: "less or equal",
			inputExpr: expression.NewBinaryExpression(token.Less,
				expression.NewIntLiteral(1),
				expression.NewIntLiteral(2),
			),
			inputArgs:  nil,
			wantResult: expression.NewBoolLiteral(true),
			wantErr:    nil,
		},
		{
			name: "less or equal",
			inputExpr: expression.NewBinaryExpression(token.LessOrEqual,
				expression.NewIntLiteral(1),
				expression.NewIntLiteral(1),
			),
			inputArgs:  nil,
			wantResult: expression.NewBoolLiteral(true),
			wantErr:    nil,
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
			inputArgs:  nil,
			wantResult: expression.NewBoolLiteral(false),
			wantErr:    nil,
		},
		{
			name: "greater",
			inputExpr: expression.NewBinaryExpression(token.Greater,
				expression.NewIntLiteral(1),
				expression.NewIntLiteral(0),
			),
			inputArgs:  nil,
			wantResult: expression.NewBoolLiteral(true),
			wantErr:    nil,
		},
		{
			name: "greater or equal",
			inputExpr: expression.NewBinaryExpression(token.GreaterOrEqual,
				expression.NewIntLiteral(1),
				expression.NewIntLiteral(2),
			),
			inputArgs:  nil,
			wantResult: expression.NewBoolLiteral(false),
			wantErr:    nil,
		},
		{
			name: "greater or equal",
			inputExpr: expression.NewBinaryExpression(token.GreaterOrEqual,
				expression.NewIntLiteral(1),
				expression.NewIntLiteral(1),
			),
			inputArgs:  nil,
			wantResult: expression.NewBoolLiteral(true),
			wantErr:    nil,
		},
		{
			name: "greater or equal",
			inputExpr: expression.NewBinaryExpression(token.GreaterOrEqual,
				expression.NewIntLiteral(1),
				expression.NewIntLiteral(0),
			),
			inputArgs:  nil,
			wantResult: expression.NewBoolLiteral(true),
			wantErr:    nil,
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
			inputArgs:  nil,
			wantResult: expression.NewBoolLiteral(true),
			wantErr:    nil,
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
			inputArgs:  nil,
			wantResult: expression.NewBoolLiteral(false),
			wantErr:    nil,
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
			inputArgs:  nil,
			wantResult: expression.NewBoolLiteral(true),
			wantErr:    nil,
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
			inputArgs:  nil,
			wantResult: expression.NewBoolLiteral(false),
			wantErr:    nil,
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
			inputArgs:  nil,
			wantResult: expression.NewBoolLiteral(false),
			wantErr:    nil,
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
			inputArgs:  nil,
			wantResult: expression.NewBoolLiteral(true),
			wantErr:    nil,
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
			inputArgs:  nil,
			wantResult: expression.NewBoolLiteral(false),
			wantErr:    nil,
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
			inputArgs:  nil,
			wantResult: expression.NewBoolLiteral(true),
			wantErr:    nil,
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
