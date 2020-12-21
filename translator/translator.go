package translator

import (
	cfg "reader"
)

type CFGTable interface {
	GetApplyCon(string, string) (uint16, string)
	GhowTable()
}

type Translator struct {
	reader *cfg.FileReader
}

func NewTranslator() *Translator {
	translator := &Translator{
		reader: nil,
	}
	return translator
}

func (t *Translator) SetReader(r *cfg.FileReader) {
	t.reader = r
}

func (t *Translator) IsTerminal(token string) bool {
	return true
}

func (t *Translator) GetApplyCon(leftToken string, rightToken string) (uint16, string) {
	return 0, ""
}

func (t *Translator) GetStartSymbal() string {
	return ""
}
