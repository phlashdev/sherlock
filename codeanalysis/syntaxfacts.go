package codeanalysis

func getBinaryOperatorPrecedence(kind SyntaxKind) int {
	switch kind {
	case StarToken, SlashToken:
		return 2
	case PlusToken, MinusToken:
		return 1
	default:
		return 0
	}
}
