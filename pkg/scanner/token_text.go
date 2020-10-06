package scanner

import (
	"fmt"

	"github.com/haunt98/evaluator/pkg/token"
)

type TokenText struct {
	Token token.Token
	Text  string
}

func (tokText TokenText) String() string {
	return fmt.Sprintf("token %s text %s", tokText.Token, tokText.Text)
}
