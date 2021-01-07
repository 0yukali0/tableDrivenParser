package translator

import (
	cfg "myreader"
	"strings"
)

func (t *Translator) First(input string) []string {
	for _, A := range t.reader.Nonterminals {
		if _, existed := t.VisitedFirst[A]; existed {
			delete(t.VisitedFirst, A)
		}
	}
	result := t.InternalFirst(input)
	result = unique(result)
	return result
}

func (t *Translator) InternalFirst(input string) (result []string) {
	//null
	if input == "L" {
		result = append(result, "L")
		return
	}
	//terminals
	b := strings.Split(input, " ")
	x := b[0]
	if len(b) > 1 {
		b = b[1:]
	} else {
		b = make([]string, 0)
	}
	if cfg.Contains(t.reader.Terminals, x) {
		result = append(result, x)
		return
	}
	result = make([]string, 1)
	if _, visited := t.VisitedFirst[x]; !visited {
		t.VisitedFirst[x] = make([]string, 1)
		rhss := t.reader.Rules[x]
		var rhs cfg.Rule
		for _, rhs = range rhss {
			tmp := t.InternalFirst(rhs.Context)
			result = append(result, tmp...)
		}
	}

	if t.Empty[x] {
		bString := strings.Join(b[1:], " ")
		tmp := t.InternalFirst(bString)
		result = append(result, tmp...)
	}
	return
}
