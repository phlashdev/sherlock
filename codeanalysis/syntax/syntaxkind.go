package syntax

type SyntaxKind string

// Tokens
const (
	BadToken              SyntaxKind = "BadToken"
	EndOfFileToken        SyntaxKind = "EndOfFileToken"
	WhitespaceToken       SyntaxKind = "WhitespaceToken"
	NumberToken           SyntaxKind = "NumberToken"
	PlusToken             SyntaxKind = "PlusToken"
	MinusToken            SyntaxKind = "MinusToken"
	StarToken             SyntaxKind = "StarToken"
	SlashToken            SyntaxKind = "SlashToken"
	OpenParenthesisToken  SyntaxKind = "OpenParenthesisToken"
	CloseParenthesisToken SyntaxKind = "CloseParenthesisToken"
)

// Expressions
const (
	LiteralExpression       SyntaxKind = "LiteralExpression"
	UnaryExpression         SyntaxKind = "UnaryExpression"
	BinaryExpression        SyntaxKind = "BinaryExpression"
	ParenthesizedExpression SyntaxKind = "ParenthesizedExpression"
)
