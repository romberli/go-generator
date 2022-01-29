package parser

import (
	"testing"

	"github.com/romberli/go-generator/pkg/lexer"
	"github.com/stretchr/testify/assert"
)

var (
	testParser *Parser
)

func init() {
	testParser = NewParser(lexer.NewLexer(lexer.NewNFAWithDefault()))
}

func TestAll(testing *testing.T) {

}

func TestParser_ParseStruct(t *testing.T) {
	asst := assert.New(t)

	text := "type TestStruct struct {\n    dependency.FiniteAutomata\n    fieldA string `json:\"field_a\" middle_ware:\"field_a\"`\n    fieldB int64  `json:\"field_b,omitempty\"`\n    fieldC *map[time.Time]string\n    *token.Token\n    fieldD **[][5]map[*time.Time][]***[10]**[0]map[int64]map[**string]**token.Token\n}"
	tokensList := testParser.lexer.Lex([]byte(text))

	for _, tokens := range tokensList {
		s, err := testParser.ParseStruct(tokens)
		asst.Nil(err, "test ParseStruct() failed")
		t.Log(s.String())
	}
}
func TestParser_Parse(t *testing.T) {
	asst := assert.New(t)

	text := "type TestStruct struct {\n    dependency.FiniteAutomata\n    fieldA string `json:\"field_a\" middle_ware:\"field_a\"`\n    fieldB int64  `json:\"field_b,omitempty\"`\n    fieldC *map[time.Time]string\n    *token.Token\n    fieldD **[][5]map[*time.Time][]***[10]**[0]map[int64]map[**string]**token.Token\n}"
	structs, err := testParser.Parse([]byte(text))
	asst.Nil(err, "test Parse() failed")
	for _, s := range structs {
		t.Log(s.String())
	}
}
