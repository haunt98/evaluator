package parser

import (
	"fmt"
	"strconv"

	"github.com/haunt98/evaluator/pkg/expression"
	"github.com/haunt98/evaluator/pkg/scanner"
	"github.com/haunt98/evaluator/pkg/token"
)

func (p *Parser) nud(tokenText scanner.TokenText) (result expression.Expression, err error) {
	fn, ok := p.nudFns[tokenText.Token]
	if !ok {
		err = fmt.Errorf("not implement null denotation token %s text %s", tokenText.Token, tokenText.Text)
		return
	}

	result, err = fn(tokenText)

	return
}

func (p *Parser) nudBool(tokenText scanner.TokenText) (result expression.Expression, err error) {
	var value bool
	value, err = strconv.ParseBool(tokenText.Text)
	if err != nil {
		err = fmt.Errorf("failed to parse bool token %s text %s", tokenText.Token, tokenText.Text)
		return
	}

	result = expression.BoolLiteral{
		Value: value,
	}

	return
}

func (p *Parser) nudInt(tokenText scanner.TokenText) (result expression.Expression, err error) {
	var value int64
	value, err = strconv.ParseInt(tokenText.Text, 10, 64)
	if err != nil {
		err = fmt.Errorf("failed to parse int token %s text %s", tokenText.Token, tokenText.Text)
		return
	}

	result = expression.IntLiteral{
		Value: value,
	}

	return
}

func (p *Parser) nudString(tokenText scanner.TokenText) (result expression.Expression, err error) {
	result = expression.StringLiteral{
		Value: tokenText.Text,
	}

	return
}

func (p *Parser) nudVar(tokenText scanner.TokenText) (result expression.Expression, err error) {
	result = expression.VarExpression{
		Value: tokenText.Text,
	}

	return
}

func (p *Parser) nudNot(_ scanner.TokenText) (result expression.Expression, err error) {
	var expr expression.Expression
	expr, err = p.parseExpression(token.LowestLevel)
	if err != nil {
		return
	}

	result = expression.UnaryExpression{
		Operator: token.Not,
		Child:    expr,
	}

	return
}

func (p *Parser) nudOpenParenthesis(_ scanner.TokenText) (result expression.Expression, err error) {
	var expr expression.Expression
	expr, err = p.parseExpression(token.LowestLevel)
	if err != nil {
		return
	}

	if expect := p.bs.Scan(); expect.Token != token.CloseParenthesis {
		err = fmt.Errorf("expect )")
		return
	}

	result = expression.ParenthesisExpression{
		Child: expr,
	}

	return
}

func (p *Parser) nudSquareBracket(_ scanner.TokenText) (result expression.Expression, err error) {
	expr := expression.ArrayExpression{
		Children: nil,
	}

	for {
		if p.bs.Peek().Token == token.CloseSquareBracket {
			break
		}

		var child expression.Expression
		child, err = p.parseExpression(token.LowestLevel)
		if err != nil {
			return
		}

		expr.Children = append(expr.Children, child)

		if p.bs.Peek().Token != token.Comma {
			break
		}

		// skip ,
		p.bs.Scan()
	}

	if expect := p.bs.Scan(); expect.Token != token.CloseSquareBracket {
		err = fmt.Errorf("expect ] got token %s text %s", expect.Token, expect.Text)
		return
	}

	result = expr

	return
}
