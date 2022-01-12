package binding

import (
	"fmt"
	"reflect"

	"github.com/phlashdev/sherlock/codeanalysis/diagnostic"
)

func reportUndefinedUnaryOperator(span diagnostic.TextSpan, operatorText string, operandType reflect.Type) diagnostic.Diagnostic {
	message := fmt.Sprintf("Unary operator %q is not defined for type %v", operatorText, operandType)
	return *diagnostic.NewDiagnostic(span, message)
}

func reportUndefinedBinaryOperator(span diagnostic.TextSpan, operatorText string, leftType reflect.Type, rightType reflect.Type) diagnostic.Diagnostic {
	message := fmt.Sprintf("Binary operator %q is not defined for types %v and %v", operatorText, leftType, rightType)
	return *diagnostic.NewDiagnostic(span, message)
}
