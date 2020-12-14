package translator

import (
	"reader"
)

type CFGTable interface {
	GetApplyCon(string, string) (uint16, string)
	GhowTable()
}

type Translator struct {
	reader reader.FileReader
}

func NewTranslator() {

}

func (t *Translator) GetApplyCon(leftToken string, rightToken string) (ruleNum uint16, terminals string) {
	return 0, ""
}
func (t *Translator) IsTerminal(token string) bool {
	return true
}
