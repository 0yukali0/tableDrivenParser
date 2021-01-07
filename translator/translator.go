package translator

import (
	cfg "myreader"
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
