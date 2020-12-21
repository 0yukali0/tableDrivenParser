package parser

import (
	"strings"
	trans "translator"
)

var (
	symbals   []string
	userinput []string
	accepted  bool
)

type Parser struct {
	translator *trans.Translator
}

func NewParser() *Parser {
	parser := &Parser{
		translator: nil,
	}
	return parser
}

func Input(remainInput string) {
	userinput = strings.Split(remainInput, " ")

}

func (p *Parser) SetTranslator(t *trans.Translator) {
	p.translator = t
}

func (p *Parser) LLParser() {
	//call PUSH(S)
	startSymbal := p.translator.GetStartSymbal()
	symbals = append(symbals, startSymbal)
	accepted = false

	for !accepted {
		n := len(startSymbal) - 1
		top = symbals[n]
		if p.translator.IsTerminal(symbals[n]) {
			for i, inP := range userinput {
				ismatch := strings.Compare(symbals[n], inP)

			}

		} else {
			p := translator.GetApplyCon(a)
		}

	}
	//	if TOS() belongs to Sigma
	//	then
	//		call MATCH(ts, TOS())
	//		if TOS() = $
	//		then  accepted  true
	//		call POP()
	//	else
	//		p <- LLtable[TOS(), ts.PEEK()]
	//		if p = 0
	//		then  call ERROR (syntax error)
	//	else  call APPLY(p)

}

//procedure APPLY(p: A -> X1...Xm)

func (p *Parser) Apply() {
	//call POP()
	n := len(symbals) - 1
	a := symbals[n] //top element
	startSymbal = startSymbal[:n]
	//end pop

	//for i = m downto 1 do
	//	call PUSH(Xi)

	productionOfA := p.trans.GetApplyCon(a)
	for i := len(productionOfA) - 1; i >= 0; i-- {
		startSymbal = append(startSymbal, productionOfA[i])
	}
}
