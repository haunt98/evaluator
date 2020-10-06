package scanner

import (
	"io"
	"strings"
	"text/scanner"

	"github.com/haunt98/evaluator/pkg/token"
)

type Scanner struct {
	textScanner *scanner.Scanner
}

func NewScanner(r io.Reader) *Scanner {
	textScanner := &scanner.Scanner{}
	textScanner.Mode = scanner.ScanIdents | scanner.ScanStrings | scanner.ScanInts
	textScanner.Init(r)

	return &Scanner{
		textScanner: textScanner,
	}
}

func (s *Scanner) Scan() (result TokenText) {
	ch := s.textScanner.Scan()
	text := s.textScanner.TokenText()

	result.Text = strings.ToLower(text)

	switch ch {
	case scanner.EOF:
		result.Token = token.EOF
	case scanner.Ident:
		switch result.Text {
		case "true", "false":
			result.Token = token.Bool
		case "or":
			result.Token = token.Or
		case "and":
			result.Token = token.And
		case "in":
			result.Token = token.In
		case "notin":
			result.Token = token.NotIn
		default:
			result.Token = token.Ident
		}
	case scanner.Int:
		result.Token = token.Int
	case scanner.String:
		result.Token = token.String
		// remove ""
		// "abc" -> abc
		result.Text = strings.Trim(result.Text, `"`)
	case '$':
		result.Token = token.Var
		// consume next
		s.textScanner.Scan()
		result.Text = s.textScanner.TokenText()
		return
	case '=':
		if expect := s.textScanner.Scan(); expect != '=' {
			result.Token = token.Illegal
			result.Text += s.textScanner.TokenText()
			return
		}

		result.Token = token.Equal
		result.Text += s.textScanner.TokenText()
	case '!':
		if expect := s.textScanner.Peek(); expect == '=' {
			result.Token = token.NotEqual
			// consume =
			s.textScanner.Scan()
			result.Text += s.textScanner.TokenText()
			return
		}

		result.Token = token.Not
	case '<':
		if expect := s.textScanner.Peek(); expect == '=' {
			result.Token = token.LessOrEqual
			// consume =
			_ = s.textScanner.Scan()
			result.Text += s.textScanner.TokenText()
			return
		}

		result.Token = token.Less
	case '>':
		if expect := s.textScanner.Peek(); expect == '=' {
			result.Token = token.GreaterOrEqual
			// consume =
			_ = s.textScanner.Scan()
			result.Text += s.textScanner.TokenText()
			return
		}

		result.Token = token.Greater
	case '(':
		result.Token = token.OpenParenthesis
	case ')':
		result.Token = token.CloseParenthesis
	case '[':
		result.Token = token.OpenSquareBracket
	case ']':
		result.Token = token.CloseSquareBracket
	case ',':
		result.Token = token.Comma
	default:
		result.Token = token.Illegal
	}

	return
}
