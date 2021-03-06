package syntax

type SyntaxKind string

// Tokens
const (
	BadToken                SyntaxKind = "BadToken"
	EndOfFileToken          SyntaxKind = "EndOfFileToken"
	WhitespaceToken         SyntaxKind = "WhitespaceToken"
	NumberToken             SyntaxKind = "NumberToken"
	PlusToken               SyntaxKind = "PlusToken"
	MinusToken              SyntaxKind = "MinusToken"
	StarToken               SyntaxKind = "StarToken"
	SlashToken              SyntaxKind = "SlashToken"
	BangToken               SyntaxKind = "BangToken"
	AmpersandAmpersandToken SyntaxKind = "AmpersandAmpersandToken"
	PipePipeToken           SyntaxKind = "PipePipeToken"
	EqualsEqualsToken       SyntaxKind = "EqualsEqualsToken"
	BangEqualsToken         SyntaxKind = "BangEqualstoken"
	OpenParenthesisToken    SyntaxKind = "OpenParenthesisToken"
	CloseParenthesisToken   SyntaxKind = "CloseParenthesisToken"
	IdentifierToken         SyntaxKind = "IdentifierToken"
)

// Keywords
const (
	TrueKeyword  SyntaxKind = "TrueKeyword"
	FalseKeyword SyntaxKind = "FalseKeyword"
)

// Expressions
const (
	LiteralExpression       SyntaxKind = "LiteralExpression"
	UnaryExpression         SyntaxKind = "UnaryExpression"
	BinaryExpression        SyntaxKind = "BinaryExpression"
	ParenthesizedExpression SyntaxKind = "ParenthesizedExpression"
)
