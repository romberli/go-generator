package lexer

import (
	"fmt"
	"testing"
	"time"

	"github.com/romberli/go-generator/pkg/dependency"
	"github.com/romberli/go-generator/pkg/token"
)

var (
	testNFA *NFA
)

type TestStruct struct {
	dependency.FiniteAutomata
	fieldA string `json:"field_a" middle_ware:"field_a"`
	fieldB int64  `json:"field_b,omitempty"`
	fieldC *map[time.Time]string
	*token.Token
	fieldD **[][5]map[*time.Time][]***[10]**[0]map[int64]map[**string]**token.Token
}

func init() {
	initTestNFA()
}

func initTestNFA() {
	testNFA = NewNFAWithDefault()
}

func TestNFA_All(t *testing.T) {
	TestNFA_Print(t)
	TestNFA_Match(t)
}

func TestNFA_Print(t *testing.T) {
	testNFA.Print()
}

func TestNFA_Match(t *testing.T) {
	strList := []string{"type", "struct", "{", "SName", "`json:\"field_a\" middleware:\"field_a\"`", "fieldC", "**[][5]map[*time.Time][]***[10]**[0]map[int64]map[**string]**token.Token"}
	strList = []string{"*map[time.Time]string"}
	for _, str := range strList {
		token := testNFA.Match([]rune(str))
		fmt.Println(token.String())
	}
}
