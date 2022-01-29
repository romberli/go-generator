package lexer

import (
	"github.com/romberli/go-generator/pkg/token"
	"github.com/romberli/go-util/constant"
)

const (
	// keyword
	TypeString   = "type"
	StructString = "struct"
)

var (
	MultiRuneMap = map[token.Type]string{
		// keyword
		token.Typ:    TypeString,
		token.Struct: StructString,
	}
	SingleRuneMap = map[token.Type]rune{
		// parenthesis
		token.LeftBrace:  LeftBraceRune,
		token.RightBrace: RightBraceRune,
	}
)

type NFA struct {
	CharacterSet *CharacterSet
	Index        int
	InitState    *State
}

// NewNFA returns a new *NFA
func NewNFA(cs *CharacterSet) *NFA {
	nfa := &NFA{
		CharacterSet: cs,
		Index:        -1,
	}

	nfa.init()

	return nfa
}

// NewNFAWithDefault returns a new *NFA with default
func NewNFAWithDefault() *NFA {
	cs := NewCharacterSetWithDefault()

	return NewNFA(cs)
}

// init initialize the NFA
func (nfa *NFA) init() {
	nfa.InitState = nfa.getNewState()

	nfa.initKeyword()
	nfa.initSingleRune()
	// identifier
	nfa.initIdentifiers()
	// field type
	nfa.initDataTypes()
	// field tag
	nfa.initFieldTag()
}

// initKeyword initialize the states that can recognize tokens which have multi runes
func (nfa *NFA) initKeyword() {
	for tokenType, tokenString := range MultiRuneMap {
		start := nfa.getNewState()
		// temporary state
		tempState := start
		nfa.InitState.AddNext(token.Epsilon, start)

		for _, c := range tokenString {
			s := nfa.getNewState()
			tempState.AddNext(c, s)
			tempState = s
		}

		final := nfa.getNewFinalState(tokenType)
		tempState.AddNext(token.Epsilon, final)
	}
}

func (nfa *NFA) initIdentifiers() {
	start, end := nfa.getIdentifierStates()
	nfa.InitState.AddNext(token.Epsilon, start)
	final := nfa.getNewFinalState(token.Identifier)
	end.AddNext(token.Epsilon, final)
}

func (nfa *NFA) initDataTypes() {
	start, end := nfa.getDataTypeStates()
	nfa.InitState.AddNext(token.Epsilon, start)
	final := nfa.getNewFinalState(token.FieldType)
	end.AddNext(token.Epsilon, final)
}

// getIdentifierStates gets the start and end states of the identifier
func (nfa *NFA) getIdentifierStates() (*State, *State) {
	start := nfa.getNewState()

	s := nfa.getNewState()
	for _, c := range nfa.CharacterSet.GetAlphabets() {
		start.AddNext(c, s)
	}

	for _, c := range nfa.CharacterSet.GetAlphabets() {
		s.AddNext(c, s)
	}
	for _, c := range nfa.CharacterSet.GetDigits() {
		s.AddNext(c, s)
	}

	end := nfa.getNewState()
	s.AddNext(token.Epsilon, end)

	return start, end
}

func (nfa *NFA) getDigitStates() (*State, *State) {
	start := nfa.getNewState()
	s := nfa.getNewState()
	start.AddNext(token.Epsilon, s)

	for _, c := range nfa.CharacterSet.GetDigits() {
		s.AddNext(c, s)
	}

	end := nfa.getNewState()
	s.AddNext(token.Epsilon, end)

	return start, end
}

func (nfa *NFA) getAsteriskStates() (*State, *State) {
	start := nfa.getNewState()
	s := nfa.getNewState()
	start.AddNext(token.Epsilon, s)
	s.AddNext(asteriskRune, s)
	end := nfa.getNewState()
	s.AddNext(token.Epsilon, end)

	return start, end
}

func (nfa *NFA) getSliceStates() (*State, *State) {
	start := nfa.getNewState()

	asteriskStart, asteriskEnd := nfa.getAsteriskStates()
	start.AddNext(token.Epsilon, asteriskStart)
	digitStart, digitEnd := nfa.getDigitStates()
	asteriskEnd.AddNext(leftBracketRune, digitStart)
	s := nfa.getNewState()
	digitEnd.AddNext(rightBracketRune, s)
	s.AddNext(token.Epsilon, asteriskStart)

	end := nfa.getNewState()
	start.AddNext(token.Epsilon, end)
	s.AddNext(token.Epsilon, end)

	return start, end
}

func (nfa *NFA) getSimpleDataTypeStates() (*State, *State) {
	start := nfa.getNewState()
	asteriskStart, asteriskEnd := nfa.getAsteriskStates()
	start.AddNext(token.Epsilon, asteriskStart)

	idStart, idEnd := nfa.getIdentifierStates()
	asteriskEnd.AddNext(token.Epsilon, idStart)

	idAfterDotStart, idAfterDotEnd := nfa.getIdentifierStates()
	idEnd.AddNext(dotRune, idAfterDotStart)

	end := nfa.getNewState()
	idEnd.AddNext(token.Epsilon, end)
	idAfterDotEnd.AddNext(token.Epsilon, end)

	return start, end
}

func (nfa *NFA) getDataTypeStates() (*State, *State) {
	start := nfa.getNewState()

	sliceStart, sliceEnd := nfa.getSliceStates()
	start.AddNext(token.Epsilon, sliceStart)

	asteriskStart, asteriskEnd := nfa.getAsteriskStates()
	sliceEnd.AddNext(token.Epsilon, asteriskStart)
	a := nfa.getNewState()
	asteriskEnd.AddNext(mRune, a)
	b := nfa.getNewState()
	a.AddNext(aRune, b)
	c := nfa.getNewState()
	b.AddNext(pRune, c)
	d := nfa.getNewState()
	c.AddNext(leftBracketRune, d)
	mapKeyStart, mapKeyEnd := nfa.getSimpleDataTypeStates()
	d.AddNext(token.Epsilon, mapKeyStart)

	simpleStart, simpleEnd := nfa.getSimpleDataTypeStates()
	sliceEnd.AddNext(token.Epsilon, simpleStart)

	end := nfa.getNewState()
	end.AddNext(token.Epsilon, sliceStart)
	mapKeyEnd.AddNext(rightBracketRune, end)
	simpleEnd.AddNext(token.Epsilon, end)

	return start, end
}

// initSingleRune initialize the states that can recognize tokens which have single rune
func (nfa *NFA) initSingleRune() {
	for tokenType, c := range SingleRuneMap {
		start := nfa.getNewState()
		nfa.InitState.AddNext(token.Epsilon, start)

		s := nfa.getNewState()
		start.AddNext(c, s)

		final := nfa.getNewFinalState(tokenType)
		s.AddNext(token.Epsilon, final)
	}
}

// initFieldTag initialize the states that can recognize field tag of a struct
func (nfa *NFA) initFieldTag() {
	start := nfa.getNewState()
	nfa.InitState.AddNext(token.Epsilon, start)

	openQuote := nfa.getNewState()
	start.AddNext(backQuoteRune, openQuote)

	for _, c := range nfa.CharacterSet.GetAlphabets() {
		openQuote.AddNext(c, openQuote)
	}
	for _, c := range nfa.CharacterSet.GetDigits() {
		openQuote.AddNext(c, openQuote)
	}
	for _, c := range nfa.CharacterSet.GetSymbols() {
		openQuote.AddNext(c, openQuote)
	}

	closeQuote := nfa.getNewState()
	openQuote.AddNext(backQuoteRune, closeQuote)

	final := nfa.getNewFinalState(token.FieldTag)
	closeQuote.AddNext(token.Epsilon, final)
}

// Print prints all the states
func (nfa *NFA) Print() {
	nfa.InitState.Print()
}

// Match matches the given runes and returns proper token
func (nfa *NFA) Match(runes []rune) *token.Token {
	return nfa.match(nfa.InitState, constant.ZeroInt, runes)
}

func (nfa *NFA) match(s *State, i int, runes []rune) *token.Token {
	if i == len(runes) {
		// all input runes are matched, check the result
		if s.IsFinal {
			// final state found
			return token.NewToken(s.TokenType, string(runes))
		}
		// this state is not a final state, need to check if there is any ε-move that can transit to the final state
		for _, ns := range s.EpsilonMove() {
			if ns.IsFinal {
				// final state found
				return token.NewToken(ns.TokenType, string(runes))
			}
		}
		// all input runes are matched, and there is no ε-move that can transit to the final state, return error token
		return token.NewToken(token.Error, string(runes))
	}

	nsList := s.Next[runes[i]]
	if nsList == nil {
		nsList = s.Next[token.Epsilon]
		if nsList == nil {
			//  can't transit to any other state, return error token
			return token.NewToken(token.Error, string(runes))
		}
	} else {
		// matched an input rune, increase the matching index
		i++
	}

	for _, ns := range nsList {
		// match next rune recursively
		t := nfa.match(ns, i, runes)
		// if returning token is not an error token, it means matched some token,
		// otherwise, means this is not a good path, need to try another one
		if t.Type != token.Error {
			return t
		}
	}

	return token.NewToken(token.Error, string(runes[:i]))
}

// getNewState gets a new state
func (nfa *NFA) getNewState() *State {
	nfa.Index++
	return NewState(nfa.Index)
}

// getNewFinalState gets a new final state
func (nfa *NFA) getNewFinalState(tokenType token.Type) *State {
	final := nfa.getNewState()
	final.IsFinal = true
	final.TokenType = tokenType

	return final
}
