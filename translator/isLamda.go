package translator

import (
	"strings"
)

var (
	// worklist = non-terminals that are discovered to derive lamda
	workList []string
	count    int
)

func (t *Translator) DerivesEmptyNonterminal() {
	//foreach A in NONTERMINALS() do
	//	SymbolDerivesEmpty(A) <- false

	for _, nonterminal := range t.reader.Nonterminals {
		//initialize as every nonterminal doesn't derives empty
		t.Empty[nonterminal] = false
	}

	t.nonterminalsEmpty = make(map[string]bool, 0)

	//foreach p in PRODUCTIONS() do

	for key, isEmpty := range t.nonterminalsEmpty {
		//key = string, rules = bool
		//	RuleDerivesEmpty(p) <- false

		// p = rule *********************************** caution
		t.nonterminalsEmpty[key] = false

		//	Count(p) <- 0
		count = 0
		//	foreach X in RHS(p)  do  Count(p) <- Count(p) + 1

		count = len(RHS(p))

		//	call CHECKFOREMPTY(p)
		t.checkForEmpty(p, key)

		/*
			for index, subRule := range rules {
				rules[index].Empty = false
				if subRule.Context == "L" {
					rules[index].Empty = true
					t.nonterminalsEmpty[key] = true
				}
			}
		*/
	}

	//	foreach X in WorkList do
	//	WorkList <- WorkList – {X}
	//	foreach x in OCCURRENCES(X) do
	//		p <- PRODUCTION(x)
	//		Count(p) <- Count(p) – 1
	//		call CHECKFOREMPTY(p)
	for i := range workList {
		workList = workList[1:]
		item := workList[i]
		// occurrences : return an iterator that visits each occerrence of item in the RHS of all rules
		for x := range occurrences(item) {
			p := production(x)
			count--
			t.checkForEmpty(p, item)

		}
	}

	/*
		for _, nonterminal := range t.reader.Nonterminals {
			if !t.nonterminalsEmpty[nonterminal] {
				t.nonterminalsEmpty[nonterminal] = t.isDeriveEmpty(nonterminal)
			}
		}
	*/
}
func (t *Translator) checkForEmpty(input string, key string) bool {

	//if Count(p) = 0
	//then
	//RuleDerivesEmpty(p) <- true
	//	A <- LHS(p)
	//	if not SymbolDerivesEmpty(A)
	//	then
	//		SymBolDerivesEmpty(A) <- true
	//		WorkList <- WorkList + {A}

	tokens := strings.Split(input, " ")

	if len(tokens) == 1 {
		if input == "L" {
			t.nonterminalsEmpty[key] = true
			nonterminal := LHS(input)

			if !t.Empty[nonterminal] {
				t.Empty[nonterminal] = true

				// worklist = non-terminals that are discovered to derive lamda
				workList = append(workList, nonterminal)
			}

		}

	}
}

/*
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
