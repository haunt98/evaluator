package parser

import (
	"fmt"

	"github.com/haunt98/evaluator/pkg/expression"
	"github.com/haunt98/evaluator/pkg/scanner"
)

func (p *Parser) led(tokenText scanner.TokenText, expr expression.Expression) (result expression.Expression, err error) {
	fn, ok := p.ledFns[tokenText.Token]
	if !ok {
		err = fmt.Errorf("not implement left denotation")
		return
	}

	result, err = fn(tokenText, expr)

	return
}

func (p *Parser) ledInfix(tokenText scanner.TokenText, expr expression.Expression) (result expression.Expression, err error) {
	var rightExpr expression.Expression
	rightExpr, err = p.parseExpression(tokenText.Token.Precedence())
	if err != nil {
		return
	}

	result = expression.BinaryExpression{
		Operator: tokenText.Token,
		Left:     expr,
		Right:    rightExpr,
	}

	return
}
