package evaluate

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/haunt98/evaluator/pkg/expression"
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

// TODO: test visit binary

func TestEvaluateVisitorVisit(t *testing.T) {
	var tests []testCase
	tests = append(tests, generateTestCaseLiteral()...)
	tests = append(tests, generateTestCaseVar()...)

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v := NewEvaluateVisitor(tc.inputArgs)

			gotResult, gotErr := v.Visit(tc.inputExpr)
			if diff := cmp.Diff(tc.wantErr, gotErr); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
			if tc.wantErr != nil {
				return
			}
			if diff := cmp.Diff(tc.wantResult, gotResult); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
