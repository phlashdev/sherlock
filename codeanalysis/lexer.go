package codeanalysis

import (
	"fmt"
	"strconv"
	"unicode"
)

type Lexer struct {
	text        string
	runes       []rune
	position    int
	diagnostics []string
}

func NewLexer(text string) *Lexer {
	return &Lexer{
		text:  text,
		runes: []rune(text),
	}
}

func (l *Lexer) Diagnostics() []string {
	return l.diagnostics
}

func (l *Lexer) current() rune {
	if l.position >= len(l.runes) {
		return 0
	}

	return l.runes[l.position]
}

func (l *Lexer) next() {
	l.position++
}

func (l *Lexer) NextToken() SyntaxToken {
	if l.position >= len(l.runes) {
		return *NewSyntaxToken(EndOfFileToken, l.position, string(rune(0)), nil)
	}

	if unicode.IsDigit(l.current()) {
		start := l.position

		for unicode.IsDigit(l.current()) {
			l.next()
		}

		text := string(l.runes[start:l.position])
		value, err := strconv.Atoi(text)
		if err != nil {
			l.diagnostics = append(l.diagnostics, fmt.Sprintf("The number %v isn't valid int", value))
		}
		return *NewSyntaxToken(NumberToken, start, text, value)
	}

	if unicode.IsSpace(l.current()) {
		start := l.position

		for unicode.IsSpace(l.current()) {
			l.next()
		}

		text := string(l.runes[start:l.position])
		return *NewSyntaxToken(WhitespaceToken, start, text, nil)
	}

	if l.current() == '+' {
		token := *NewSyntaxToken(PlusToken, l.position, "+", nil)
		l.next()
		return token
	} else if l.current() == '-' {
		token := *NewSyntaxToken(MinusToken, l.position, "-", nil)
		l.next()
		return token
	} else if l.current() == '*' {
		token := *NewSyntaxToken(StarToken, l.position, "*", nil)
		l.next()
		return token
	} else if l.current() == '/' {
		token := *NewSyntaxToken(SlashToken, l.position, "/", nil)
		l.next()
		return token
	} else if l.current() == '(' {
		token := *NewSyntaxToken(OpenParenthesisToken, l.position, "(", nil)
		l.next()
		return token
	} else if l.current() == ')' {
		l.next()
		token := *NewSyntaxToken(CloseParenthesisToken, l.position, ")", nil)
		l.next()
		return token
	}

	l.diagnostics = append(l.diagnostics, fmt.Sprintf("ERROR: bad character input: %q", l.current()))

	text := string(l.runes[l.position])
	token := *NewSyntaxToken(BadToken, l.position, text, nil)
	l.next()
	return token
}
