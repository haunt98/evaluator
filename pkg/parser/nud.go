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
		return nil, fmt.Errorf("not implement null denotation %s", tokenText)
	}

	return fn(tokenText)
}

func (p *Parser) nudBool(tokenText scanner.TokenText) (expression.Expression, error) {
	value, err := strconv.ParseBool(tokenText.Text)
	if err != nil {
		return nil, fmt.Errorf("failed to parse bool %s: %w", tokenText, err)
	}

	return expression.NewBoolLiteral(value), nil
}

func (p *Parser) nudInt(tokenText scanner.TokenText) (expression.Expression, error) {
	value, err := strconv.ParseInt(tokenText.Text, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse int token %s: %w", tokenText, err)
	}

	return expression.NewIntLiteral(value), nil
}

func (p *Parser) nudString(tokenText scanner.TokenText) (expression.Expression, error) {
	return expression.NewStringLiteral(tokenText.Text), nil
}

func (p *Parser) nudVar(tokenText scanner.TokenText) (expression.Expression, error) {
	return expression.NewVarExpression(tokenText.Text), nil
}

func (p *Parser) nudNot(_ scanner.TokenText) (expression.Expression, error) {
	expr, err := p.parseWithPrecedence(token.LowestLevel)
	if err != nil {
		return nil, err
	}

	return expression.NewUnaryExpression(token.Not, expr), nil
}

func (p *Parser) nudOpenParenthesis(_ scanner.TokenText) (expression.Expression, error) {
	expr, err := p.parseWithPrecedence(token.LowestLevel)
	if err != nil {
		return nil, err
	}

	if expect := p.bs.Scan(); expect.Token != token.CloseParenthesis {
		return nil, fmt.Errorf("expect %s got %s", token.CloseParenthesis, expect)
	}

	return expr, nil
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
		return nil, fmt.Errorf("expect %s got %s", token.CloseSquareBracket, expect)
	}

	return expression.NewArrayExpression(children...), nil
}
