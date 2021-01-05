package myreader

import (
	"bufio"
	"errors"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

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

//Read file context into buffer
func (r *Reader) ReadFile(filepath string) error {
	f, err := os.Open(filepath)
	if err != nil {
		return errors.New("File path is invalid")
	}
	defer f.Close()

	fd, err := ioutil.ReadAll(f)
	if err != nil {
		return errors.New("Read to file fail")
	}
	r.Context = string(fd)
	r.initRules()
	return nil
}

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

func (r *Reader) ReadRule() error {
	in := bufio.NewReader(os.Stdin)
	input, _ := in.ReadString('\n')
	input = strings.Trim(input, "\t\r ")
	indexOfNumterminal := strings.Index(input, " ") + 1
	input = input[indexOfNumterminal:]

	return nil
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

func (r *Reader) Recount() {
	index := 1
	for _, nonterminal := range r.Nonterminals {
		rules := r.Rules[nonterminal]
		for subIndex, subRule := range rules {
			subRule.Num = uint(index)
			rules[subIndex] = subRule
			index++
		}
	}
}

func (r *Reader) initTreminals() {
	nonterminals := r.Nonterminals
	for _, keyRules := range r.Rules {
		for _, keyRule := range keyRules {
			tokens := strings.Split(keyRule.Context, " ")
			for _, token := range tokens {
				if !Contains(nonterminals, token) && !Contains(r.Terminals, token) {
					r.Terminals = append(r.Terminals, token)
				}
			}
		}
	}
	terminals := r.Terminals
	sort.SliceStable(terminals, func(i, j int) bool {
		if terminals[i] == "$" {
			return false
		}

		if terminals[j] == "$" {
			return true
		}

		result := strings.Compare(terminals[i], terminals[j]) * -1
		switch result {
		case 1:
			return true
		case -1:
			return false
		default:
			return true
		}
	})
	r.Terminals = terminals
}

func Contains(set []string, input string) bool {
	for _, value := range set {
		if input == value {
			return true
		}
	}
	return false
}
