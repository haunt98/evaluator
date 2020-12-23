package scanner

import (
	"strings"
	"testing"

	"github.com/haunt98/evaluator/pkg/token"
	"github.com/stretchr/testify/assert"
)

type bufferScannerTestCase struct {
	name  string
	input string
	wants []TokenText
}

func TestBufferScannerScan(t *testing.T) {
	tests := []bufferScannerTestCase{
		{
			name:  "or and",
			input: "or and",
			wants: []TokenText{
				{
					Token: token.Or,
					Text:  "or",
				},
				{
					Token: token.And,
					Text:  "and",
				},
				{
					Token: token.EOF,
					Text:  "",
				},
			},
		},
		{
			name:  "in notin >= <= ! !=",
			input: "in notin >= <= ! !=",
			wants: []TokenText{
				{
					Token: token.In,
					Text:  "in",
				},
				{
					Token: token.NotIn,
					Text:  "notin",
				},
				{
					Token: token.GreaterOrEqual,
					Text:  ">=",
				},
				{
					Token: token.LessOrEqual,
					Text:  "<=",
				},
				{
					Token: token.Not,
					Text:  "!",
				},
				{
					Token: token.NotEqual,
					Text:  "!=",
				},
				{
					Token: token.EOF,
					Text:  "",
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := NewScanner(strings.NewReader(tc.input))
			bufferScanner := NewBufferScanner(s)

			for _, want := range tc.wants {
				got := bufferScanner.Scan()
				assert.Equal(t, want, got)
			}
		})
	}
}

func TestBufferScannerPeek(t *testing.T) {
	tests := []bufferScannerTestCase{
		{
			name:  "([])",
			input: "([])",
			wants: []TokenText{
				{
					Token: token.OpenParenthesis,
					Text:  "(",
				},
				{
					Token: token.OpenParenthesis,
					Text:  "(",
				},
				{
					Token: token.OpenParenthesis,
					Text:  "(",
				},
				{
					Token: token.OpenParenthesis,
					Text:  "(",
				},
			},
		},
		{
			name:  "==",
			input: "==",
			wants: []TokenText{
				{
					Token: token.Equal,
					Text:  "==",
				},
				{
					Token: token.Equal,
					Text:  "==",
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := NewScanner(strings.NewReader(tc.input))
			bufferScanner := NewBufferScanner(s)

			for _, want := range tc.wants {
				got := bufferScanner.Peek()
				assert.Equal(t, want, got)
			}
		})
	}
}
