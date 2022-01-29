package lexer

import (
	"fmt"
	"testing"
)

var (
	testNFALexer *Lexer
	testDFALexer *Lexer
)

func init() {
	initLexer()
}

func initLexer() {
	initTestNFA()
	initTestDFA()
	testNFALexer = NewLexer(testNFA)
	testDFALexer = NewLexer(testNFA)
}

func TestLexer_All(t *testing.T) {
	TestLexer_Lex(t)
}

func TestLexer_Lex(t *testing.T) {
	text := "type TestStruct struct {\n    dependency.FiniteAutomata\n    fieldA string `json:\"field_a\" middle_ware:\"field_a\"`\n    fieldB int64  `json:\"field_b,omitempty\"`\n    fieldC time.Time\n    *token.Token\n}"
	// text = "type TestStruct struct {\n    fieldC *map[time.Time]string\n}"

	structs := testNFALexer.Lex([]byte(text))
	fmt.Println("==========NFA==========")
	for _, tokens := range structs {
		for _, tok := range tokens {
			fmt.Println(tok.String())
		}
	}
	fmt.Println("==========NFA==========")

	// structs = testDFALexer.Lex(text)
	// fmt.Println("==========DFA==========")
	// for _, tokens := range structs {
	//     for _, tok := range tokens {
	//         fmt.Println(tok.String())
	//     }
	// }
	// fmt.Println("==========DFA==========")

}
