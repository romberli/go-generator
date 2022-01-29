package token

import (
	"fmt"

	"github.com/romberli/go-util/constant"
)

type Type int

const (
	// keyword
	Typ Type = iota + 1
	Struct
	// identifier
	Identifier
	// struct
	StructName
	FieldName
	FieldType
	FieldTag
	// comparison operator
	LeftBrace
	RightBrace
	// separator
	NewLine
	Colon
	DoubleQuote
	// white space
	WhiteSpace
	// error
	Error
)

var (
	// epsilon
	Epsilon     rune = constant.ZeroInt
	KeywordList      = []Type{Typ, Struct}
)

// String returns the string representation of the token type
func (t Type) String() string {
	switch t {
	case Typ:
		return "TypeKeyword"
	case Struct:
		return "StructKeyword"
	case Identifier:
		return "Identifier"
	case LeftBrace:
		return "LeftBrace"
	case RightBrace:
		return "RightBrace"
	case StructName:
		return "StructName"
	case FieldName:
		return "FieldName"
	case FieldType:
		return "FieldType"
	case FieldTag:
		return "FieldTag"
	case NewLine:
		return "NewLine"
	case Colon:
		return "Colon"
	case DoubleQuote:
		return "DoubleQuote"
	case WhiteSpace:
		return "WhiteSpace"
	case Error:
		return "Error"
	default:
		return "Unknown"
	}
}

// IsKeyword returns if the token type is a keyword
func (t Type) IsKeyword() bool {
	for _, keyword := range KeywordList {
		if t == keyword {
			return true
		}
	}

	return false
}

type Token struct {
	Type   Type
	Lexeme string
}

// NewToken returns a new *Token
func NewToken(tokenType Type, lexeme string) *Token {
	return &Token{
		Type:   tokenType,
		Lexeme: lexeme,
	}
}

// String returns the string representation of the token
func (t *Token) String() string {
	return fmt.Sprintf("{tokenType: %s, lexeme: %s}", t.Type.String(), t.Lexeme)
}
