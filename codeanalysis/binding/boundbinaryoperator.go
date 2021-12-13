package binding

import (
	"fmt"
	"reflect"

	"github.com/phlashdev/sherlock/codeanalysis/syntax"
)

var (
	binaryOperators []boundBinaryOperator = []boundBinaryOperator{
		*newBoundBinaryOperatorWithSameType(syntax.PlusToken, Addition, reflect.TypeOf(0)),
		*newBoundBinaryOperatorWithSameType(syntax.MinusToken, Subtraction, reflect.TypeOf(0)),
		*newBoundBinaryOperatorWithSameType(syntax.StarToken, Multiplication, reflect.TypeOf(0)),
		*newBoundBinaryOperatorWithSameType(syntax.SlashToken, Division, reflect.TypeOf(0)),

		*newBoundBinaryOperator(syntax.AmpersandAmpersandToken, LogicalAnd, reflect.TypeOf(false), reflect.TypeOf(false), reflect.TypeOf(false)),
		*newBoundBinaryOperator(syntax.PipePipeToken, LogicalOr, reflect.TypeOf(false), reflect.TypeOf(false), reflect.TypeOf(false)),
	}
)

type boundBinaryOperator struct {
	syntaxKind syntax.SyntaxKind
	kind       boundBinaryOperatorKind
	leftType   reflect.Type
	rightType  reflect.Type
	resultType reflect.Type
}

func newBoundBinaryOperatorWithSameType(syntaxKind syntax.SyntaxKind, kind boundBinaryOperatorKind, operatorType reflect.Type) *boundBinaryOperator {
	return newBoundBinaryOperator(syntaxKind, kind, operatorType, operatorType, operatorType)
}

func newBoundBinaryOperator(syntaxKind syntax.SyntaxKind, kind boundBinaryOperatorKind, leftType reflect.Type, rightType reflect.Type, resultType reflect.Type) *boundBinaryOperator {
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
