package scanner

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/haunt98/evaluator/pkg/token"
)

func TestBufferScannerScan(t *testing.T) {
	tests := []struct {
		name  string
		input string
		wants []TokenText
	}{
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
				if diff := cmp.Diff(want, got); diff != "" {
					t.Errorf("mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
