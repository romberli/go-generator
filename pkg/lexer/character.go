package lexer

const (
	// ascii boundary
	asciiDigitStart             = 48
	asciiDigitEnd               = 57
	asciiAlphabetUpperCaseStart = 65
	asciiAlphabetUpperCaseEnd   = 90
	asciiAlphabetLowerCaseStart = 97
	asciiAlphabetLowerCaseEnd   = 122

	underBarRune     = '_'
	commaRune        = ','
	colonRune        = ':'
	backQuoteRune    = '`'
	doubleQuoteRune  = '"'
	spaceRune        = ' '
	EpsilonRune      = 'ε'
	dotRune          = '.'
	asteriskRune     = '*'
	leftBracketRune  = '['
	rightBracketRune = ']'
	mRune            = 'm'
	aRune            = 'a'
	pRune            = 'p'
)

type CharacterSet struct {
	Alphabets []rune
	Digits    []rune
	Symbols   []rune
}

// NewCharacterSet returns a new *CharacterSet
func NewCharacterSet(alphabets, digits, symbols []rune) *CharacterSet {
	return &CharacterSet{
		Alphabets: alphabets,
		Digits:    digits,
		Symbols:   symbols,
	}
}

// NewCharacterSetWithDefault returns a new *CharacterSet with default
func NewCharacterSetWithDefault() *CharacterSet {
	alphabets := []rune{underBarRune}

	for i := asciiAlphabetLowerCaseStart; i <= asciiAlphabetLowerCaseEnd; i++ {
		alphabets = append(alphabets, rune(i))
	}
	for i := asciiAlphabetUpperCaseStart; i <= asciiAlphabetUpperCaseEnd; i++ {
		alphabets = append(alphabets, rune(i))
	}

	var digits []rune
	for i := asciiDigitStart; i <= asciiDigitEnd; i++ {
		digits = append(digits, rune(i))
	}

	return NewCharacterSet(
		alphabets,
		digits,
		[]rune{colonRune, doubleQuoteRune, commaRune, spaceRune, dotRune},
	)
}

// GetAlphabets returns the alphabet runes
func (cs *CharacterSet) GetAlphabets() []rune {
	return cs.Alphabets
}

// GetAlphabets returns the digit runes
func (cs *CharacterSet) GetDigits() []rune {
	return cs.Digits
}

// GetSymbols returns the symbol runes
func (cs *CharacterSet) GetSymbols() []rune {
	return cs.Symbols
}
