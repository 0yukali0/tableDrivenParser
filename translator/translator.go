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

func GetApplyCon(left string, right string) (ruleNum uint16, terminals string) {
	return 0, ""
}
