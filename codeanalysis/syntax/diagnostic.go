package syntax

import (
	"fmt"
	"reflect"

	"github.com/phlashdev/sherlock/codeanalysis/diagnostic"
)

func reportInvalidNumber(span diagnostic.TextSpan, text string, textType reflect.Type) diagnostic.Diagnostic {
	message := fmt.Sprintf("The number %s isn't valid %s.", text, textType)
	return *diagnostic.NewDiagnostic(span, message)
}

func reportBadCharacter(position int, char rune) diagnostic.Diagnostic {
	span := *diagnostic.NewTextSpan(position, 1)
	message := fmt.Sprintf("Bad character input: %q.", char)
	return *diagnostic.NewDiagnostic(span, message)
}

func reportUnexpectedToken(span diagnostic.TextSpan, actualKind SyntaxKind, expectedKind SyntaxKind) diagnostic.Diagnostic {
	message := fmt.Sprintf("Unexpected token <%v>, expected <%v>.", actualKind, expectedKind)
	return *diagnostic.NewDiagnostic(span, message)
}
