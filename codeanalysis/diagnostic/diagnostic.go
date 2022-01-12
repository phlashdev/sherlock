package diagnostic

type Diagnostic struct {
	span    TextSpan
	message string
}

func NewDiagnostic(span TextSpan, message string) *Diagnostic {
	return &Diagnostic{
		span:    span,
		message: message,
	}
}

func (d *Diagnostic) Span() TextSpan {
	return d.span
}

func (d *Diagnostic) Message() string {
	return d.message
}

func (d *Diagnostic) String() string {
	return d.message
}

func Report(diagnostics []Diagnostic, span TextSpan, message string) []Diagnostic {
	diagnostic := NewDiagnostic(span, message)
	return append(diagnostics, *diagnostic)
}
