package dependency

import (
	"github.com/romberli/go-generator/pkg/token"
)

type FiniteAutomata interface {
	// Print prints all the states/sets of the finite automata
	Print()
	// Match matches the given runes and returns proper token
	Match(runes []rune) *token.Token
}
