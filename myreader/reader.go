package myreader

type Rule struct {
	Num     uint
	Context string
	Empty   bool
}

type Info struct {
	Num   uint
	Token string
}

type Reader struct {
	Context      string
	Terminals    []string
	Nonterminals []string
	RulesOwners  []Info
	RulesOwner   string
	Rules        map[string][]Rule
}

func NewReader() *Reader {
	fileReader := &Reader{
		Context:      "",
		Terminals:    make([]string, 1),
		Nonterminals: make([]string, 1),
		RulesOwners:  make([]Info, 1),
		RulesOwner:   "",
		Rules:        make(map[string][]Rule),
	}
	return fileReader
}
