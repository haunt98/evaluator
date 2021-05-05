package parser

import (
	"fmt"

	"github.com/haunt98/evaluator/expression"
	"github.com/haunt98/evaluator/scanner"
)

func (p *Parser) led(tokenText scanner.TokenText, expr expression.Expression) (expression.Expression, error) {
	fn, ok := p.ledFns[tokenText.Token]
	if !ok {
		return nil, fmt.Errorf("not implement left denotation token %s text %s", tokenText.Token, tokenText.Text)
	}

	return fn(tokenText, expr)
}

func (p *Parser) ledInfix(tokenText scanner.TokenText, expr expression.Expression) (expression.Expression, error) {
	rightExpr, err := p.parseWithPrecedence(tokenText.Token.Precedence())
	if err != nil {
		return nil, err
	}

	return expression.NewBinaryExpression(tokenText.Token, expr, rightExpr), nil
}
