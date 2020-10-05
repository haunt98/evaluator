package parser

import (
	"fmt"
	"strconv"

	"github.com/haunt98/evaluator/pkg/expression"
	"github.com/haunt98/evaluator/pkg/scanner"
	"github.com/haunt98/evaluator/pkg/token"
)

const (
	defaultNumberOfChildren = 10
)

func (p *Parser) nud(tokenText scanner.TokenText) (expression.Expression, error) {
	fn, ok := p.nudFns[tokenText.Token]
	if !ok {
		return nil, fmt.Errorf("not implement null denotation token %s text %s", tokenText.Token, tokenText.Text)
	}

	return fn(tokenText)
}

func (p *Parser) nudBool(tokenText scanner.TokenText) (expression.Expression, error) {
	value, err := strconv.ParseBool(tokenText.Text)
	if err != nil {
		return nil, fmt.Errorf("failed to parse bool token %s text %s", tokenText.Token, tokenText.Text)
	}

	return &expression.BoolLiteral{
		Value: value,
	}, nil
}

func (p *Parser) nudInt(tokenText scanner.TokenText) (expression.Expression, error) {
	value, err := strconv.ParseInt(tokenText.Text, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse int token %s text %s", tokenText.Token, tokenText.Text)
	}

	return &expression.IntLiteral{
		Value: value,
	}, nil
}

func (p *Parser) nudString(tokenText scanner.TokenText) (expression.Expression, error) {
	return &expression.StringLiteral{
		Value: tokenText.Text,
	}, nil
}

func (p *Parser) nudVar(tokenText scanner.TokenText) (expression.Expression, error) {
	return &expression.VarExpression{
		Value: tokenText.Text,
	}, nil
}

func (p *Parser) nudNot(_ scanner.TokenText) (expression.Expression, error) {
	expr, err := p.parseWithPrecedence(token.LowestLevel)
	if err != nil {
		return nil, err
	}

	return &expression.UnaryExpression{
		Operator: token.Not,
		Child:    expr,
	}, nil
}

func (p *Parser) nudOpenParenthesis(_ scanner.TokenText) (expression.Expression, error) {
	expr, err := p.parseWithPrecedence(token.LowestLevel)
	if err != nil {
		return nil, err
	}

	if expect := p.bs.Scan(); expect.Token != token.CloseParenthesis {
		return nil, fmt.Errorf("expect ) got token %s text %s", expect.Token, expect.Text)
	}

	return &expression.ParenthesisExpression{
		Child: expr,
	}, nil
}

func (p *Parser) nudSquareBracket(_ scanner.TokenText) (expression.Expression, error) {
	children := make([]expression.Expression, 0, defaultNumberOfChildren)

	for {
		if p.bs.Peek().Token == token.CloseSquareBracket {
			break
		}

		var child expression.Expression
		child, err := p.parseWithPrecedence(token.LowestLevel)
		if err != nil {
			return nil, err
		}

		children = append(children, child)

		if p.bs.Peek().Token != token.Comma {
			break
		}

		// skip ,
		p.bs.Scan()
	}

	if expect := p.bs.Scan(); expect.Token != token.CloseSquareBracket {
		return nil, fmt.Errorf("expect ] got token %s text %s", expect.Token, expect.Text)
	}

	return &expression.ArrayExpression{
		Children: children,
	}, nil
}
