package codeanalysis

type SyntaxKind string

const (
	NumberToken             SyntaxKind = "NumberToken"
	WhitespaceToken         SyntaxKind = "WhitespaceToken"
	PlusToken               SyntaxKind = "PlusToken"
	MinusToken              SyntaxKind = "MinusToken"
	StarToken               SyntaxKind = "StarToken"
	SlashToken              SyntaxKind = "SlashToken"
	OpenParenthesisToken    SyntaxKind = "OpenParenthesisToken"
	CloseParenthesisToken   SyntaxKind = "CloseParenthesisToken"
	BadToken                SyntaxKind = "BadToken"
	EndOfFileToken          SyntaxKind = "EndOfFileToken"
	NumberExpression        SyntaxKind = "NumberExpression"
	BinaryExpression        SyntaxKind = "BinaryExpression"
	ParenthesizedExpression SyntaxKind = "ParenthesizedExpression"
)
