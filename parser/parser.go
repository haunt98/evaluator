// Implement Pratt parser
// https://tdop.github.io/
package parser

import (
	"fmt"
	"strings"

	"github.com/haunt98/evaluator/expression"
	"github.com/haunt98/evaluator/scanner"
	"github.com/haunt98/evaluator/token"
)

type Parser struct {
	bs *scanner.BufferScanner

	nudFns map[token.Token]nudFn
	ledFns map[token.Token]ledFn
}

// nud short for null denotation
type nudFn func(scanner.TokenText) (expression.Expression, error)

// led short for left denotation
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
	return p.parseWithPrecedence(token.LowestLevel)
}

func (p *Parser) parseWithPrecedence(precedence int) (expression.Expression, error) {
	tokenText := p.bs.Scan()
	result, err := p.nud(tokenText)
	if err != nil {
		return nil, fmt.Errorf("failed to null denotation %s: %w", tokenText, err)
	}

	for {
		if precedence >= p.bs.Peek().Token.Precedence() {
			break
		}

		tokenText = p.bs.Scan()
		result, err = p.led(tokenText, result)
		if err != nil {
			return nil, fmt.Errorf("failed to left denotation %s: %w", tokenText, err)
		}
	}

	return result, nil
}
