package binding

import (
	"fmt"
	"reflect"

	"github.com/phlashdev/sherlock/codeanalysis/syntax"
)

var (
	binaryOperators []boundBinaryOperator = []boundBinaryOperator{
		*newBoundBinaryOperator(syntax.PlusToken, Addition, reflect.TypeOf(0)),
		*newBoundBinaryOperator(syntax.MinusToken, Subtraction, reflect.TypeOf(0)),
		*newBoundBinaryOperator(syntax.StarToken, Multiplication, reflect.TypeOf(0)),
		*newBoundBinaryOperator(syntax.SlashToken, Division, reflect.TypeOf(0)),

		*newBoundBinaryOperatorWithSameOperandType(syntax.EqualsEqualsToken, Equals, reflect.TypeOf(0), reflect.TypeOf(false)),
		*newBoundBinaryOperatorWithSameOperandType(syntax.BangEqualsToken, NotEquals, reflect.TypeOf(0), reflect.TypeOf(false)),

		*newBoundBinaryOperatorWithDifferentTypes(syntax.AmpersandAmpersandToken, LogicalAnd, reflect.TypeOf(false), reflect.TypeOf(false), reflect.TypeOf(false)),
		*newBoundBinaryOperatorWithDifferentTypes(syntax.PipePipeToken, LogicalOr, reflect.TypeOf(false), reflect.TypeOf(false), reflect.TypeOf(false)),
		*newBoundBinaryOperator(syntax.EqualsEqualsToken, Equals, reflect.TypeOf(false)),
		*newBoundBinaryOperator(syntax.BangEqualsToken, NotEquals, reflect.TypeOf(false)),
	}
)

type boundBinaryOperator struct {
	syntaxKind syntax.SyntaxKind
	kind       boundBinaryOperatorKind
	leftType   reflect.Type
	rightType  reflect.Type
	resultType reflect.Type
}

func newBoundBinaryOperator(syntaxKind syntax.SyntaxKind, kind boundBinaryOperatorKind, operatorType reflect.Type) *boundBinaryOperator {
	return newBoundBinaryOperatorWithDifferentTypes(syntaxKind, kind, operatorType, operatorType, operatorType)
}

func newBoundBinaryOperatorWithSameOperandType(syntaxKind syntax.SyntaxKind, kind boundBinaryOperatorKind, operandType reflect.Type, resultType reflect.Type) *boundBinaryOperator {
	return newBoundBinaryOperatorWithDifferentTypes(syntaxKind, kind, operandType, operandType, resultType)
}

func newBoundBinaryOperatorWithDifferentTypes(syntaxKind syntax.SyntaxKind, kind boundBinaryOperatorKind, leftType reflect.Type, rightType reflect.Type, resultType reflect.Type) *boundBinaryOperator {
	return &boundBinaryOperator{
		syntaxKind: syntaxKind,
		kind:       kind,
		leftType:   leftType,
		rightType:  rightType,
		resultType: resultType,
	}
}

func (b *boundBinaryOperator) SyntaxKind() syntax.SyntaxKind {
	return b.syntaxKind
}

func (b *boundBinaryOperator) Kind() boundBinaryOperatorKind {
	return b.kind
}

func (b *boundBinaryOperator) LeftType() reflect.Type {
	return b.leftType
}

func (b *boundBinaryOperator) RightType() reflect.Type {
	return b.rightType
}

func (b *boundBinaryOperator) ResultType() reflect.Type {
	return b.resultType
}

func bindBinaryOperator(syntaxKind syntax.SyntaxKind, leftType reflect.Type, rightType reflect.Type) (boundBinaryOperator, error) {
	for _, op := range binaryOperators {
		if op.syntaxKind == syntaxKind && op.leftType == leftType && op.rightType == rightType {
			return op, nil
		}
	}

	return boundBinaryOperator{}, fmt.Errorf("operator not found")
}
