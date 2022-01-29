package lexer

const (
	// char
	UnderBarRune = '_'
	// separator
	CommaRune      = ','
	LeftBraceRune  = '{'
	RightBraceRune = '}'
	// white space
	SpaceRune   = ' '
	TabRune     = '\t'
	ReturnRune  = '\r'
	NewLineRune = '\n'

	NewLineString = "\\n"
)

// IsAlphabet returns if the given rune is an alphabet
func IsAlphabet(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == UnderBarRune || c == dotRune
}

// IsDigit returns if the given rune is a digit
func IsDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

// IsWhiteSpace returns if the given rune is a white space
func IsWhiteSpace(c rune) bool {
	return c == SpaceRune || c == TabRune || c == ReturnRune || c == NewLineRune
}

// IsWhiteSpace returns if the given rune is either an alphabet or a digit
func IsAlphabetOrDigit(c rune) bool {
	return IsAlphabet(c) || IsDigit(c)
}

func IsBracket(c rune) bool {
	return c == leftBracketRune || c == rightBracketRune
}

func IsAsterisk(c rune) bool {
	return c == asteriskRune
}
