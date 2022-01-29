package generator

import (
	"fmt"

	"github.com/romberli/go-generator/pkg/parser"
	"github.com/romberli/go-generator/pkg/util"
	"github.com/romberli/go-util/constant"
)

type Generator struct {
	parser *parser.Parser
}

func NewGenerator(parser *parser.Parser) *Generator {
	return &Generator{
		parser: parser,
	}
}

func (g *Generator) GetParser() *parser.Parser {
	return g.parser
}

func (g *Generator) GenerateGetter(source []byte) ([]byte, error) {
	var str string
	structs, err := g.GetParser().Parse(source)
	if err != nil {
		return nil, err
	}

	for _, s := range structs {
		str += s.Reveal() + constant.CRLFString
		str += g.getStructGetter(s)
	}

	return []byte(str), nil
}

func (g *Generator) getStructGetter(s *util.Struct) string {
	var str string

	for _, field := range s.Fields {
		str += g.getFieldGetter(s, field)
	}

	return str
}

func (g *Generator) getFieldGetter(s *util.Struct, field *util.Field) string {
	return fmt.Sprintf("//Get%s returns %s of %s\nfunc (%s %s) Get%s() %s {\n    return %s.%s\n}\n\n",
		field.Name, field.Name, s.Name, s.GetShortName(), s.Name, field.Name, field.Type, s.GetShortName(), field.Name)
}