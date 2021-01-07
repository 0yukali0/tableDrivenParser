package translator

import "strings"

func (t *Translator) InitTable() {
	/*
		for _, nonterminal := range t.reader.Nonterminals {
			for _, subRule := range t.reader.Rules[nonterminal] {
				var (
					first  []string
					follow []string
					name   string
				)
			}
		}
	*/
}

func (t *Translator) ShowTable() {
	top := strings.Join(t.reader.Terminals, " ")

	for _, nonterminal := range t.reader.Nonterminals {
		for _, terminal := range t.reader.Terminals {
			name := nonterminal + " " + terminal
			if space, existed := t.table[name]; existed {

			} else {

			}
		}
	}
}
