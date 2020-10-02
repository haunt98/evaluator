// Implement Pratt parser
// https://tdop.github.io/
package parser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/haunt98/evaluator/pkg/expression"
	"github.com/haunt98/evaluator/pkg/scanner"
	"github.com/haunt98/evaluator/pkg/token"
)

type Parser struct {
	bs     *scanner.BufferScanner
	nudFns map[token.Token]nudFn // nud short for null denotation
	ledFns map[token.Token]ledFn // led short for left denotation
}

type nudFn func(scanner.TokenText) (expression.Expression, error)
type ledFn func(scanner.TokenText, expression.Expression) (expression.Expression, error)

func NewParser(input string) *Parser {
	s := scanner.NewScanner(strings.NewReader(input))
	bs := scanner.NewBufferScanner(s)

	p := &Parser{
		bs: bs,
	}

	p.nudFns = map[token.Token]nudFn{
		token.Bool:              p.nudBool,
		token.Int:               p.nudInt,
		token.String:            p.nudString,
		token.Var:               p.nudVar,
		token.Not:               p.nudNot,
		token.OpenParenthesis:   p.nudOpenParenthesis,
		token.OpenSquareBracket: p.nudSquareBracket,
	}

	p.ledFns = map[token.Token]ledFn{
		token.Or:             p.ledInfix,
		token.And:            p.ledInfix,
		token.Equal:          p.ledInfix,
		token.NotEqual:       p.ledInfix,
		token.Less:           p.ledInfix,
		token.LessOrEqual:    p.ledInfix,
		token.Greater:        p.ledInfix,
		token.GreaterOrEqual: p.ledInfix,
		token.In:             p.ledInfix,
		token.NotIn:          p.ledInfix,
	}

	return p
}

func (p *Parser) Parse() (expression.Expression, error) {
	return p.parseExpression(token.LowestLevel)
}

func (p *Parser) parseExpression(precedence int) (result expression.Expression, err error) {
	tokenText := p.bs.Scan()
	result, err = p.nullDenotation(tokenText)
	if err != nil {
		return
	}

	for {
		if precedence >= p.bs.Peek().Token.Precedence() {
			break
		}

		tokenText = p.bs.Scan()
		result, err = p.leftDenotation(tokenText, result)
		if err != nil {
			return
		}
	}

	return
}

func (p *Parser) nullDenotation(tokenText scanner.TokenText) (result expression.Expression, err error) {
	fn, ok := p.nudFns[tokenText.Token]
	if !ok {
		err = fmt.Errorf("not implement null denotation")
		return
	}

	result, err = fn(tokenText)

	return
}

func (p *Parser) nudBool(tokenText scanner.TokenText) (result expression.Expression, err error) {
	var value bool
	value, err = strconv.ParseBool(tokenText.Text)
	if err != nil {
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
		err = fmt.Errorf("expect ]")
		return
	}

	result = expr

	return
}

func (p *Parser) leftDenotation(tokenText scanner.TokenText, expr expression.Expression) (result expression.Expression, err error) {
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
