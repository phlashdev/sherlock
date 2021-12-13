package syntax

type SyntaxToken struct {
	kind     SyntaxKind
	position int
	text     string
	value    interface{}
}

func NewSyntaxToken(kind SyntaxKind, position int, text string, value interface{}) *SyntaxToken {
	return &SyntaxToken{
		kind:     kind,
		position: position,
		text:     text,
		value:    value,
	}
}

func (t *SyntaxToken) Kind() SyntaxKind {
	return t.kind
}

func (t *SyntaxToken) Position() int {
	return t.position
}

func (t *SyntaxToken) Text() string {
	return t.text
}

func (t *SyntaxToken) Value() interface{} {
	return t.value
}

func (t *SyntaxToken) GetChildren() []SyntaxNode {
	return []SyntaxNode{}
}
