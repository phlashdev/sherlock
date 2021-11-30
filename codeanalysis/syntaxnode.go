package codeanalysis

type SyntaxNode interface {
	Kind() SyntaxKind
	GetChildren() []SyntaxNode
}
