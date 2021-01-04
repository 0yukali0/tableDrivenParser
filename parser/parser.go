package parser

import (
	"fmt"
	"strconv"
	"strings"
	trans "translator"
)

var (
	symbals    []string
	userinput  []string
	applyRules []string
	ruleNums   []string
	accepted   bool
	isLamda    bool
	ruleNum    string
	syntex     string
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

	for !accepted {
		//POP()
		isLamda = false
		n := len(startSymbal) - 1
		top := symbals[n]

		if p.translator.IsTerminal(top) {
			//call match()
			inP := userinput[0]
			ismatch := strings.Compare(top, inP)
			if ismatch == 0 {
				userinput = userinput[1:]
			} else {
				syntex = "Error(Expected " + top + ")"
				break
			}

			if strings.Compare(top, "$") == 0 {
				accepted = true
				syntex = "Accept"
			}

			//call POP()
			symbals = symbals[:n]
			//end pop
		} else {
			x, llTable := p.translator.GetApplyCon(top, userinput[0])
			if strings.Compare(llTable, "L") == 0 {
				isLamda = true
			}
			fmt.Println(x)
			y := uint64(x)
			ruleNum = strconv.FormatUint(y, 10)

			if x == 0 {
				syntex = "Error(" + top + "vs. " + userinput[0]
				break
			} else {
				ruleNums = append(ruleNums, ruleNum)
				fmt.Println("call Apply()")
				if !isLamda {
					applyRules = strings.Split(llTable, " ")
				}
				p.Apply(applyRules)
			}

		}

	}

	fmt.Print(ruleNums)
	fmt.Print(syntex)
}

//procedure APPLY(p: A -> X1...Xm)

func (p *Parser) Apply(productionOfA []string) {
	//call POP()
	fmt.Println("in Apply()")
	n := len(symbals) - 1
	symbals = symbals[:n]
	//end pop

	//for i = m downto 1 do
	//	call PUSH(Xi)
	if !isLamda {
		fmt.Printf("apply rule: %v \n", productionOfA)
		for i := len(productionOfA) - 1; i >= 0; i-- {
			symbals = append(symbals, productionOfA[i])
		}
	}
}
