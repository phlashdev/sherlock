package binding

type boundBinaryOperatorKind string

const (
	Addition              boundBinaryOperatorKind = "Addition"
	Subtraction           boundBinaryOperatorKind = "Subtraction"
	Multiplication        boundBinaryOperatorKind = "Multiplication"
	Division              boundBinaryOperatorKind = "Division"
	LogicalAnd            boundBinaryOperatorKind = "LogicalAnd"
	LogicalOr             boundBinaryOperatorKind = "LogicalOr"
	UnknownBinaryOperator boundBinaryOperatorKind = "Unkown"
)
