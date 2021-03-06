package syntax

func getUnaryOperatorPrecedence(kind SyntaxKind) int {
	switch kind {
	case PlusToken, MinusToken, BangToken:
		return 6
	default:
		return 0
	}
}

func getBinaryOperatorPrecedence(kind SyntaxKind) int {
	switch kind {
	case StarToken, SlashToken:
		return 5
	case PlusToken, MinusToken:
		return 4
	case EqualsEqualsToken, BangEqualsToken:
		return 3
	case AmpersandAmpersandToken:
		return 2
	case PipePipeToken:
		return 1
	default:
		return 0
	}
}

func getKeywordKind(text string) SyntaxKind {
	switch text {
	case "true":
		return TrueKeyword
	case "false":
		return FalseKeyword
	default:
		return IdentifierToken
	}
}
