package syntax

type SyntaxNode interface {
	Kind() SyntaxKind
	GetChildren() []SyntaxNode
}
