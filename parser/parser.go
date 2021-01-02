package parser

import (
	"fmt"
	"strings"
	trans "translator"
)

var (
	symbals    []string
	userinput  []string
	applyRules []string
	ruleNums   []string
	accepted   bool
	ruleNum    int
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
		//POP()
		n := len(startSymbal) - 1
		top := symbals[n]

		if p.translator.IsTerminal(top) {

			inP := userinput[0]
			ismatch := strings.Compare(top, inP)
			if ismatch == 0 {
				userinput = userinput[1:]
			}

			if strings.Compare(top, "$") == 0 {
				accepted = true
			}

			//call POP()
			symbals = symbals[:n]
			//end pop
		} else {
			ruleNum, llTable := p.translator.GetApplyCon(top, userinput[0])
			applyRules = strings.Split(llTable, "")

			if ruleNum == 0 {
				fmt.Print("syntax error")
			} else {
				ruleNums = append(ruleNums, ruleNum)
				p.Apply(applyRules)
			}

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

func (p *Parser) Apply(productionOfA []string) {
	//call POP()
	n := len(symbals) - 1
	symbals = symbals[:n]
	//end pop

	//for i = m downto 1 do
	//	call PUSH(Xi)

	for i := len(productionOfA) - 1; i >= 0; i-- {
		symbals = append(symbals, productionOfA[i])
	}
}
