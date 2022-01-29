package lexer

import (
	"github.com/romberli/go-generator/pkg/dependency"
	"github.com/romberli/go-generator/pkg/token"
)

type Lexer struct {
	fa dependency.FiniteAutomata
}

// NewLexer returns a new *Lexer
func NewLexer(fa dependency.FiniteAutomata) *Lexer {
	return &Lexer{
		fa: fa,
	}
}

// GetFiniteAutomata returns the finite automata of the lexer
func (l *Lexer) GetFiniteAutomata() dependency.FiniteAutomata {
	return l.fa
}

// Lex scans the input string and returns a token list
func (l *Lexer) Lex(bytes []byte) [][]*token.Token {
	var (
		isFieldTag bool
		isBody     bool
		isField    bool
		runes      []rune
		tokens     []*token.Token
		structs    [][]*token.Token
	)

	textRunes := []rune(string(bytes))
	length := len(textRunes)

	for i, c := range textRunes {
		if c == backQuoteRune && !isFieldTag {
			// string literal starts
			isFieldTag = true
			runes = append(runes, c)
			continue
		}

		if c == backQuoteRune && isFieldTag {
			// string literal ends
			isFieldTag = false
			runes = append(runes, c)
			// match string literal token
			tokens = append(tokens, l.GetFiniteAutomata().Match(runes))
			runes = nil
			continue
		}

		if isFieldTag {
			runes = append(runes, c)
			if i == length-1 {
				// field tag does not end with single back quote
				tokens = append(tokens, token.NewToken(token.Error, string(runes)))
			}
			continue
		}

		switch c {
		case SpaceRune, TabRune, ReturnRune:
			continue
		case NewLineRune:
			if isField {
				tokens = append(tokens, token.NewToken(token.NewLine, NewLineString))
				isField = false
			}
		case LeftBraceRune:
			tokens = append(tokens, token.NewToken(token.LeftBrace, string(LeftBraceRune)))
			isBody = true
		case RightBraceRune:
			tokens = append(tokens, token.NewToken(token.RightBrace, string(RightBraceRune)))
			isBody = false
			structs = append(structs, tokens)
			tokens = nil
		default:
			runes = append(runes, c)
			nc := textRunes[i+1]
			if i >= length-1 || (!IsAlphabetOrDigit(nc) && !IsBracket(nc)) && !IsAsterisk(nc) {
				t := l.GetFiniteAutomata().Match(runes)
				if isBody && t.Type == token.Identifier || t.Type == token.FieldType || t.Type == token.FieldTag {
					isField = true
				}
				tokens = append(tokens, t)
				runes = nil
			}
		}
	}

	return structs
}
