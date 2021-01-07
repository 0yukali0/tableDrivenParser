package translator

/*
func (t *Translator) DerivesEmptyNonterminal() bool {
	for _, nonterminal := range t.reader.Nonterminals {
		t.Empty[nonterminal] = false
	}
	t.nonterminalsEmpty = make(map[string]bool, 0)
	for key, rules := range t.nonterminalsEmpty {
		t.nonterminalsEmpty[key] = false
		for index, subRule := range rules {
			rules[index].Empty = false
			if subRule.Context == "L" {
				rules[index].Empty = true
				t.nonterminalsEmpty[key] = true
			}
		}
	}
	for _, nonterminal := range t.reader.Nonterminals {
		if !t.nonterminalsEmpty[nonterminal] {
			t.nonterminalsEmpty[nonterminal] = t.isDeriveEmpty(nonterminal)
		}
	}
}

func (t *Translator) isDeriveEmpty(input string) bool {
	tokens := strings.Split(input, " ")
	if len(tokens) == 1 {
		if input == "L" {
			return true
		}
		if cfg.Contains(t.reader.Nonterminals, input) {
			if t.nonterminalsEmpty[input] {
				return true
			}
			return t.isDeriveEmpty()
		} else {
			return false
		}
	}
	return false
}
*/
