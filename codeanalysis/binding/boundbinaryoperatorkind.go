package binding

type boundBinaryOperatorKind string

const (
	Addition              boundBinaryOperatorKind = "Addition"
	Subtraction           boundBinaryOperatorKind = "Subtraction"
	Multiplication        boundBinaryOperatorKind = "Multiplication"
	Division              boundBinaryOperatorKind = "Division"
	UnknownBinaryOperator boundBinaryOperatorKind = "Unkown"
)
