package syntax

type SyntaxTree struct {
	diagnostics    []string
	root           ExpressionSyntax
	endOfFileToken SyntaxToken
}

func NewSyntaxTree(diagnostics []string, root ExpressionSyntax, endOfFileToken SyntaxToken) *SyntaxTree {
	return &SyntaxTree{
		diagnostics:    diagnostics,
		root:           root,
		endOfFileToken: endOfFileToken,
	}
}

func (s *SyntaxTree) Diagnostics() []string {
	return s.diagnostics
}

func (s *SyntaxTree) Root() ExpressionSyntax {
	return s.root
}

func (s *SyntaxTree) EndOfFileToken() SyntaxToken {
	return s.endOfFileToken
}

func Parse(text string) SyntaxTree {
	parser := NewParser(text)
	return parser.Parse()
}
