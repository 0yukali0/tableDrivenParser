package myreader

import (
	"sort"
	"strings"
)

func (r *Reader) Recount() {
	index := 1
	for _, nonterminal := range r.Nonterminals {
		rules := r.Rules[nonterminal]
		for subIndex, subRule := range rules {
			subRule.Num = uint(index)
			rules[subIndex] = subRule
			index++
		}
	}
}

func (r *Reader) initTreminals() {
	nonterminals := r.Nonterminals
	for _, keyRules := range r.Rules {
		for _, keyRule := range keyRules {
			tokens := strings.Split(keyRule.Context, " ")
			for _, token := range tokens {
				if !Contains(nonterminals, token) && !Contains(r.Terminals, token) {
					r.Terminals = append(r.Terminals, token)
				}
			}
		}
	}
	terminals := r.Terminals
	sort.SliceStable(terminals, func(i, j int) bool {
		if terminals[i] == "$" {
			return false
		}

		if terminals[j] == "$" {
			return true
		}

		result := strings.Compare(terminals[i], terminals[j]) * -1
		switch result {
		case 1:
			return true
		case -1:
			return false
		default:
			return true
		}
	})
	r.Terminals = terminals
}
