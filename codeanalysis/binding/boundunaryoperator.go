package binding

import (
	"fmt"
	"reflect"

	"github.com/phlashdev/sherlock/codeanalysis/syntax"
)

var (
	unaryOperators []boundUnaryOperator = []boundUnaryOperator{
		*newBoundUnaryOperatorWithSameType(syntax.BangToken, LogicalNegation, reflect.TypeOf(true)),

		*newBoundUnaryOperatorWithSameType(syntax.PlusToken, Identity, reflect.TypeOf(0)),
		*newBoundUnaryOperatorWithSameType(syntax.MinusToken, Negation, reflect.TypeOf(0)),
	}
)

type boundUnaryOperator struct {
	syntaxKind  syntax.SyntaxKind
	kind        boundUnaryOperatorKind
	operandType reflect.Type
	resultType  reflect.Type
}

func newBoundUnaryOperatorWithSameType(syntaxKind syntax.SyntaxKind, kind boundUnaryOperatorKind, operandType reflect.Type) *boundUnaryOperator {
	return newBoundUnaryOperator(syntaxKind, kind, operandType, operandType)
}

func newBoundUnaryOperator(syntaxKind syntax.SyntaxKind, kind boundUnaryOperatorKind, operandType reflect.Type, resultType reflect.Type) *boundUnaryOperator {
	return &boundUnaryOperator{
		syntaxKind:  syntaxKind,
		kind:        kind,
		operandType: operandType,
		resultType:  resultType,
	}
}

func (b *boundUnaryOperator) SyntaxKind() syntax.SyntaxKind {
	return b.syntaxKind
}

func (b *boundUnaryOperator) Kind() boundUnaryOperatorKind {
	return b.kind
}

func (b *boundUnaryOperator) OperandType() reflect.Type {
	return b.operandType
}

func (b *boundUnaryOperator) ResultType() reflect.Type {
	return b.resultType
}

func bindUnaryOperator(syntaxKind syntax.SyntaxKind, operandType reflect.Type) (boundUnaryOperator, error) {
	for _, op := range unaryOperators {
		if op.syntaxKind == syntaxKind && op.operandType == operandType {
			return op, nil
		}
	}

	return boundUnaryOperator{}, fmt.Errorf("operator not found")
}
