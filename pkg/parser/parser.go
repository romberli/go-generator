package parser

import (
	"strings"

	"github.com/pingcap/errors"
	"github.com/romberli/go-generator/pkg/lexer"
	"github.com/romberli/go-generator/pkg/token"
	"github.com/romberli/go-generator/pkg/util"
	"github.com/romberli/go-util/constant"
)

type Parser struct {
	lexer  *lexer.Lexer
	next   int
	fields []*util.Field
}

func NewParser(lexer *lexer.Lexer) *Parser {
	return &Parser{lexer: lexer}
}

func (p *Parser) GetLexer() *lexer.Lexer {
	return p.lexer
}

func (p *Parser) Parse(bytes []byte) ([]*util.Struct, error) {
	var structs []*util.Struct

	tokensList := p.GetLexer().Lex(bytes)
	for _, tokens := range tokensList {
		s, err := p.ParseStruct(tokens)
		if err != nil {
			return nil, err
		}

		structs = append(structs, s)
	}

	return structs, nil
}

func (p *Parser) ParseStruct(tokens []*token.Token) (*util.Struct, error) {
	p.next = constant.ZeroInt
	p.fields = nil
	s := util.NewEmptyStruct()

	// type keyword
	if tokens[p.next].Type == token.Typ {
		p.next++
	} else {
		return nil, errors.Errorf("expected type keyword, but got %s", tokens[p.next].String())
	}
	// identifier(struct name)
	if tokens[p.next].Type == token.Identifier {
		s.Name = tokens[p.next].Lexeme
		p.next++
	} else {
		return nil, errors.Errorf("expected identifier, but got %s", tokens[p.next].String())
	}
	// struct keyword
	if tokens[p.next].Type == token.Struct {
		p.next++
	} else {
		return nil, errors.Errorf("expected struct keyword, but got %s", tokens[p.next].String())
	}
	// left brace
	if tokens[p.next].Type == token.LeftBrace {
		p.next++
	} else {
		return nil, errors.Errorf("expected left brace, but got %s", tokens[p.next].String())
	}
	// field list
	err := p.ParseField(tokens)
	if err != nil {
		return nil, err
	}
	s.Fields = p.fields
	// right brace
	if tokens[p.next].Type != token.RightBrace {
		return nil, errors.Errorf("expected right brace, but got %s", tokens[p.next].String())
	}
	p.next++

	return s, nil
}

func (p *Parser) ParseField(tokens []*token.Token) error {
	save := p.next
	err := p.parseFieldOne(tokens)
	if err == nil {
		return nil
	}

	p.next = save
	err = p.parseFieldTwo(tokens)
	if err == nil {
		return nil
	}

	p.next = save
	err = p.parseFieldThree(tokens)
	if err == nil {
		return nil
	}
	// epsilon
	p.next = save

	return nil
}

func (p *Parser) parseFieldOne(tokens []*token.Token) error {
	field := util.NewFieldWithNum(1)
	// field
	if tokens[p.next].Type == token.Identifier || tokens[p.next].Type == token.FieldType {
		field.Type = tokens[p.next].Lexeme
		typeStrList := strings.Split(field.Type, constant.AsteriskString)
		typeStr := typeStrList[len(typeStrList)-1]
		strList := strings.Split(typeStr, constant.DotString)
		field.Name = strList[len(strList)-1]
		p.next++
	} else {
		return errors.Errorf("expected identifier or field type, but got %s", tokens[p.next].String())
	}

	if tokens[p.next].Type == token.NewLine {
		p.fields = append(p.fields, field)
		p.next++
	} else {
		return errors.Errorf("expected new line(\\n), but got %s", tokens[p.next].String())
	}

	return p.ParseField(tokens)
}

func (p *Parser) parseFieldTwo(tokens []*token.Token) error {
	field := util.NewFieldWithNum(2)

	// field name
	if tokens[p.next].Type == token.Identifier {
		field.Name = tokens[p.next].Lexeme
		p.next++
	} else {
		return errors.Errorf("expected identifier, but got %s", tokens[p.next].String())
	}
	// field type
	if tokens[p.next].Type == token.Identifier || tokens[p.next].Type == token.FieldType {
		field.Type = tokens[p.next].Lexeme
		p.next++
	} else {
		return errors.Errorf("expected identifier or field type, but got %s", tokens[p.next].String())
	}
	// new line
	if tokens[p.next].Type == token.NewLine {
		p.fields = append(p.fields, field)
		p.next++
	} else {
		return errors.Errorf("expected new line(\\n), but got %s", tokens[p.next].String())
	}

	return p.ParseField(tokens)
}

func (p *Parser) parseFieldThree(tokens []*token.Token) error {
	field := util.NewFieldWithNum(3)
	// field name
	if tokens[p.next].Type == token.Identifier {
		field.Name = tokens[p.next].Lexeme
		p.next++
	} else {
		return errors.Errorf("expected identifier, but got %s", tokens[p.next].String())
	}
	// field type
	if tokens[p.next].Type == token.Identifier || tokens[p.next].Type == token.FieldType {
		field.Type = tokens[p.next].Lexeme
		p.next++
	} else {
		return errors.Errorf("expected identifier or field type, but got %s", tokens[p.next].String())
	}
	// field tag
	if tokens[p.next].Type == token.FieldTag {
		field.Tag = tokens[p.next].Lexeme
		p.next++
	} else {
		return errors.Errorf("expected field tag, but got %s", tokens[p.next].String())
	}
	// new line
	if tokens[p.next].Type == token.NewLine {
		p.fields = append(p.fields, field)
		p.next++
	} else {
		return errors.Errorf("expected new line(\\n), but got %s", tokens[p.next].String())
	}

	return p.ParseField(tokens)
}
