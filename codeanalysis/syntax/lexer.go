package syntax

import (
	"reflect"
	"strconv"
	"unicode"

	"github.com/phlashdev/sherlock/codeanalysis/diagnostic"
)

type lexer struct {
	text        string
	runes       []rune
	position    int
	diagnostics []diagnostic.Diagnostic
}

func NewLexer(text string) *lexer {
	return &lexer{
		text:  text,
		runes: []rune(text),
	}
}

func (l *lexer) Diagnostics() []diagnostic.Diagnostic {
	return l.diagnostics
}

func (l *lexer) current() rune {
	return l.peek(0)
}

func (l *lexer) lookahead() rune {
	return l.peek(1)
}

func (l *lexer) peek(offset int) rune {
	index := l.position + offset

	if index >= len(l.runes) {
		return 0
	}

	return l.runes[index]
}

func (l *lexer) next() {
	l.position++
}

func (l *lexer) Lex() SyntaxToken {
	if l.position >= len(l.runes) {
		return *NewSyntaxToken(EndOfFileToken, l.position, string(rune(0)), nil)
	}

	start := l.position

	if unicode.IsDigit(l.current()) {
		for unicode.IsDigit(l.current()) {
			l.next()
		}

		text := string(l.runes[start:l.position])
		value, err := strconv.Atoi(text)
		if err != nil {
			length := l.position - start
			span := *diagnostic.NewTextSpan(start, length)
			report := reportInvalidNumber(span, l.text, reflect.TypeOf(0))
			l.diagnostics = append(l.diagnostics, report)
		}
		return *NewSyntaxToken(NumberToken, start, text, value)
	}

	if unicode.IsSpace(l.current()) {
		for unicode.IsSpace(l.current()) {
			l.next()
		}

		text := string(l.runes[start:l.position])
		return *NewSyntaxToken(WhitespaceToken, start, text, nil)
	}

	if unicode.IsLetter(l.current()) {
		for unicode.IsLetter(l.current()) {
			l.next()
		}

		text := string(l.runes[start:l.position])
		kind := getKeywordKind(text)
		return *NewSyntaxToken(kind, start, text, nil)
	}

	var token *SyntaxToken
	switch l.current() {
	case '+':
		token = NewSyntaxToken(PlusToken, l.position, "+", nil)
	case '-':
		token = NewSyntaxToken(MinusToken, l.position, "-", nil)
	case '*':
		token = NewSyntaxToken(StarToken, l.position, "*", nil)
	case '/':
		token = NewSyntaxToken(SlashToken, l.position, "/", nil)
	case '(':
		token = NewSyntaxToken(OpenParenthesisToken, l.position, "(", nil)
	case ')':
		token = NewSyntaxToken(CloseParenthesisToken, l.position, ")", nil)
	case '&':
		if l.lookahead() == '&' {
			token = NewSyntaxToken(AmpersandAmpersandToken, l.position, "&&", nil)
		}
	case '|':
		if l.lookahead() == '|' {
			token = NewSyntaxToken(PipePipeToken, l.position, "||", nil)
		}
	case '=':
		if l.lookahead() == '=' {
			token = NewSyntaxToken(EqualsEqualsToken, l.position, "==", nil)
		}
	case '!':
		if l.lookahead() == '=' {
			token = NewSyntaxToken(BangEqualsToken, l.position, "!=", nil)
		} else {
			token = NewSyntaxToken(BangToken, l.position, "!", nil)
		}
	}

	if token == nil {
		report := reportBadCharacter(l.position, l.current())
		l.diagnostics = append(l.diagnostics, report)

		text := string(l.runes[l.position])
		token = NewSyntaxToken(BadToken, l.position, text, nil)
	}

	switch token.kind {
	case AmpersandAmpersandToken, PipePipeToken, EqualsEqualsToken, BangEqualsToken:
		l.position += 2
	default:
		l.position++
	}

	return *token
}
