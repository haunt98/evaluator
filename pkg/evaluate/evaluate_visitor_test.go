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
			name: "bool",
			inputExpr: &expression.BoolLiteral{
				Value: true,
			},
			inputArgs:  nil,
			wantResult: expression.NewBoolLiteral(true),
			wantErr:    nil,
		},
		{
			name: "int",
			inputExpr: &expression.IntLiteral{
				Value: 1,
			},
			inputArgs:  nil,
			wantResult: expression.NewIntLiteral(1),
			wantErr:    nil,
		},
		{
			name: "string",
			inputExpr: &expression.StringLiteral{
				Value: "a",
			},
			inputArgs:  nil,
			wantResult: expression.NewStringLiteral("a"),
			wantErr:    nil,
		},
	}
}

func generateTestCaseVar() []testCase {
	return []testCase{
		{
			name: "var",
			inputExpr: &expression.VarExpression{
				Value: "x",
			},
			inputArgs: map[string]interface{}{
				"x": "xxx",
			},
			wantResult: expression.NewStringLiteral("xxx"),
			wantErr:    nil,
		},
	}
}

// TODO: more array, include var
func generateTestCaseArray() []testCase {
	return []testCase{
		{
			name: "array",
			inputExpr: &expression.ArrayExpression{
				Children: []expression.Expression{
					&expression.BoolLiteral{
						Value: true,
					},
				},
			},
			inputArgs:  nil,
			wantResult: expression.NewArrayExpression(expression.NewBoolLiteral(true)),
			wantErr:    nil,
		},
		{
			name: "array",
			inputExpr: &expression.ArrayExpression{
				Children: []expression.Expression{
					&expression.BoolLiteral{
						Value: true,
					},
					&expression.BoolLiteral{
						Value: false,
					},
				},
			},
			inputArgs:  nil,
			wantResult: expression.NewArrayExpression(expression.NewBoolLiteral(true), expression.NewBoolLiteral(false)),
			wantErr:    nil,
		},
	}
}

// TODO: test visit binary

func TestVisitor_Visit(t *testing.T) {
	var tests []testCase
	tests = append(tests, generateTestCaseLiteral()...)
	tests = append(tests, generateTestCaseVar()...)
	tests = append(tests, generateTestCaseArray()...)

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