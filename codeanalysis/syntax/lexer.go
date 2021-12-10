package syntax

import (
	"fmt"
	"strconv"
	"unicode"
)

type lexer struct {
	text        string
	runes       []rune
	position    int
	diagnostics []string
}

func NewLexer(text string) *lexer {
	return &lexer{
		text:  text,
		runes: []rune(text),
	}
}

func (l *lexer) Diagnostics() []string {
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

	if unicode.IsLetter(l.current()) {
		start := l.position

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
	case '!':
		token = NewSyntaxToken(BangToken, l.position, "!", nil)
	case '&':
		if l.lookahead() == '&' {
			token = NewSyntaxToken(AmpersandAmpersandToken, l.position, "&&", nil)
		}
	case '|':
		if l.lookahead() == '|' {
			token = NewSyntaxToken(PipePipeToken, l.position, "||", nil)
		}
	}

	if token == nil {
		l.diagnostics = append(l.diagnostics, fmt.Sprintf("ERROR: bad character input: %q", l.current()))

		text := string(l.runes[l.position])
		token = NewSyntaxToken(BadToken, l.position, text, nil)
	}

	if token.kind == AmpersandAmpersandToken || token.kind == PipePipeToken {
		l.position += 2
	} else {
		l.position++
	}

	return *token
}
