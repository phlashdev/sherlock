package codeanalysis

func getUnaryOperatorPrecedence(kind SyntaxKind) int {
	switch kind {
	case PlusToken, MinusToken:
		return 3
	default:
		return 0
	}
}

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
