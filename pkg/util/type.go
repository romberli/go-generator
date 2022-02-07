package util

import (
	"fmt"
	"unicode"

	"github.com/romberli/go-util/constant"
)

type Struct struct {
	Name   string
	Fields []*Field
}

func NewEmptyStruct() *Struct {
	return &Struct{
		Name:   constant.EmptyString,
		Fields: nil,
	}
}

func (s *Struct) String() string {
	return fmt.Sprintf("{Name: %s, Fields: %s", s.Name, s.Fields)
}

func (s *Struct) Reveal() string {
	var fieldStr string
	str := fmt.Sprintf("type %s struct {\n", s.Name)
	for _, field := range s.Fields {
		switch field.Num {
		case 1:
			fieldStr = fmt.Sprintf("    %s\n", field.Type)
		case 2:
			fieldStr = fmt.Sprintf("    %s    %s\n", field.Name, field.Type)
		case 3:
			fieldStr = fmt.Sprintf("    %s    %s    %s\n", field.Name, field.Type, field.Tag)
		}

		str += fieldStr
	}

	str += constant.RightBraceString + constant.CRLFString

	return str
}

func (s *Struct) GetShortName() string {
	var (
		runes       []rune
		isUppercase bool
	)
	for i, c := range s.Name {
		if i == constant.ZeroInt {
			runes = append(runes, unicode.ToLower(c))
			continue
		}

		if c >= 'A' && c <= 'Z' && !isUppercase {
			isUppercase = true
			runes = append(runes, unicode.ToLower(c))
			continue
		}

		isUppercase = false
	}

	return string(runes)
}

type Field struct {
	Name string
	Type string
	Tag  string
	Num  int
}

func NewEmptyField() *Field {
	return &Field{}
}

func NewFieldWithNum(num int) *Field {
	return &Field{
		Num: num,
	}
}

func (f *Field) String() string {
	tag := f.Tag
	if tag == constant.EmptyString {
		tag = `""`
	}

	return fmt.Sprintf("{Name: %s, Type: %s, Tag: %s, Num: %d}", f.Name, f.Type, tag, f.Num)
}
