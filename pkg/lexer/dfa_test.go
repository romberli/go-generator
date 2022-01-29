package lexer

import (
	"fmt"
	"testing"
)

var (
	testDFA *DFA
)

func init() {
	initTestDFA()
}

func initTestDFA() {
	testDFA = NewDFAWithDefault()
}

func TestDFA_All(t *testing.T) {
	TestDFA_Print(t)
	TestDFA_Match(t)
}

func TestDFA_Print(t *testing.T) {
	testDFA.Print()
}

func TestDFA_Match(t *testing.T) {
	strList := []string{"type", "struct", "{", "SName", "`json:\"field_a\" middleware:\"field_a\"`", "fieldC"}

	for _, str := range strList {
		token := testDFA.Match([]rune(str))
		fmt.Println(token.String())
	}
}
