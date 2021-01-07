package myreader

import "strings"

func (r *Reader) initRules() {
	lines := strings.Split(r.Context, "\n")
	result := ""
	//delete number and reorder context
	for index, line := range lines {
		line = strings.Trim(line, " \r\t")
		indexOfNumterminal := strings.Index(line, " ") + 1
		line = line[indexOfNumterminal:]
		line = strings.Trim(line, " \r\t")
		if strings.Contains(line, ">") {
			if index != 0 {
				result += line + ">"
			} else {
				result += line
			}
		} else if strings.Contains(line, "|") {
			result += line
		}
	}
	r.Context = result
	lines = strings.Split(r.Context, ">")
	//init with owner > rule | rule |....> owner > ..
	for index := 0; len(lines) != 0; index++ {
		r.setRulesOwner(lines[0])
		r.addRulesOwner(uint(index), r.RulesOwner)
		r.setRules(lines[1])
		lines = lines[2:]
	}
}

func (r *Reader) setRulesOwner(owner string) {
	r.RulesOwner = strings.Trim(owner, " \r\t\n")
}

func (r Reader) addRulesOwner(index uint, owner string) {
	if !Contains(r.Nonterminals, owner) {
		r.Nonterminals = append(r.Nonterminals, owner)
		ownerInfo := Info{
			Num:   uint(index),
			Token: r.RulesOwner,
		}
		r.Nonterminals = append(r.Nonterminals, owner)
		r.RulesOwners = append(r.RulesOwners, ownerInfo)
	}
}

func (r *Reader) setRules(rulesLine string) {
	var rulesInfo []Rule = make([]Rule, 1)
	rulesLine = strings.Trim(rulesLine, " \r\t\n")
	rules := strings.Split(rulesLine, " ")

	for _, oneRule := range rules {
		ruleInfo := Rule{
			Context: strings.Trim(oneRule, " \r\t\n"),
		}
		rulesInfo = append(rulesInfo, ruleInfo)
	}
	_, ok := r.Rules[r.RulesOwner]
	if !ok {
		r.Rules[r.RulesOwner] = rulesInfo
	} else {
		r.Rules[r.RulesOwner] = append(r.Rules[r.RulesOwner], rulesInfo...)
	}
}
