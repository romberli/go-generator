package generator

import (
	"testing"

	"github.com/romberli/go-generator/pkg/lexer"
	"github.com/romberli/go-generator/pkg/parser"
	"github.com/stretchr/testify/assert"
)

const testStructString = "type TestStruct struct {\n    dependency.FiniteAutomata\n    fieldA string `json:\"field_a\" middle_ware:\"field_a\"`\n    fieldB int64  `json:\"field_b,omitempty\"`\n    fieldC *map[time.Time]string\n    *token.Token\n    fieldD **[][5]map[*time.Time][]***[10]**[0]map[int64]map[**string]**token.Token\n}"

var (
	testGenerator *Generator
)

func init() {
	testGenerator = NewGenerator(parser.NewParser(lexer.NewLexer(lexer.NewNFAWithDefault())))
}

func TestAll(t *testing.T) {
	TestGenerator_GenerateGetter(t)
}

func TestGenerator_GenerateGetter(t *testing.T) {
	asst := assert.New(t)

	bytes, err := testGenerator.GenerateGetter([]byte(testStructString))
	asst.Nil(err, "test GenerateGetter() failed")
	t.Log(string(bytes))
}
