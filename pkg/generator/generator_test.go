package generator

import (
	"testing"

	"github.com/romberli/go-generator/pkg/lexer"
	"github.com/romberli/go-generator/pkg/parser"
)

var (
	testGenerator *Generator
)

func init() {
	testGenerator = NewGenerator(parser.NewParser(lexer.NewLexer(lexer.NewNFAWithDefault())))
}

func TestAll(testing *testing.T) {

}

func TestGenerator_GenerateGetter(t *testing.T) {
	// asst := assert.New(t)

}
