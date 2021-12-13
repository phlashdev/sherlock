package binding

type boundUnaryOperatorKind string

const (
	Identity             boundUnaryOperatorKind = "Identity"
	Negation             boundUnaryOperatorKind = "Negation"
	LogicalNegation      boundUnaryOperatorKind = "LogicalNegation"
	UnknownUnaryOperator boundUnaryOperatorKind = "Unkown"
)
