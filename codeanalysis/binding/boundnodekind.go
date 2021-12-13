package binding

type boundNodeKind string

const (
	LiteralExpression boundNodeKind = "LiteralExpression"
	UnaryExpression   boundNodeKind = "UnaryExpression"
	BinaryExpression  boundNodeKind = "BinaryExpression"
)
