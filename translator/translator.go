package translator

import (
	cfg "myreader"
	"strings"
)

type CFGTable interface {
	GetApplyCon(string, string) (uint16, string)
	GhowTable()
}

type Translator struct {
	reader            *cfg.Reader
	VisitedFirst      map[string][]string
	VisitedFollow     map[string][]string
	nonterminalsEmpty map[string]bool
	Empty             map[string]bool
	table             map[string][]cfg.Rule
}

func NewTranslator() *Translator {
	translator := &Translator{
		reader:        nil,
		VisitedFirst:  make(map[string][]string, 1),
		VisitedFollow: make(map[string][]string, 1),

		Empty: make(map[string]bool, 1),
		table: make(map[string][]cfg.Rule, 1),
	}
	return translator
}

func (t *Translator) SetReader(r *cfg.Reader) {
	t.reader = r
}

func (t *Translator) IsTerminal(token string) bool {
	return cfg.Contains(t.reader.Terminals, token)
}

func (t *Translator) GetApplyCon(leftToken string, rightToken string) (uint16, string) {
	return 0, ""
}

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

func unique(input []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range input {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
