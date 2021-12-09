package binding

type boundUnaryOperatorKind string

const (
	Identity             boundUnaryOperatorKind = "Identity"
	Negation             boundUnaryOperatorKind = "Negation"
	UnknownUnaryOperator boundUnaryOperatorKind = "Unkown"
)
