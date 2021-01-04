package reader

import (
	"bufio"
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

type rule struct {
	num     uint
	context string
}

type info struct {
	num   uint
	token string
}

type Reader struct {
	context      string
	Terminals    []string
	Nonterminals []string
	RulesOwners  []info
	RulesOwner   string
	Rules        map[string][]rule
}

func NewReader() *Reader {
	fileReader := &Reader{
		context:      "",
		Terminals:    make([]string, 1),
		Nonterminals: make([]string, 1),
		RulesOwner:   "",
		Rules:        make(map[string][]rule),
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
	r.context = string(fd)
	r.initRules()
	return nil
}

func (r *Reader) initRules() {
	lines := strings.Split(r.context, "\n")
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
	r.context = result
	lines = strings.Split(r.context, ">")
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
		ownerInfo := info{
			num:   uint(index),
			token: r.RulesOwner,
		}
		r.Nonterminals = append(r.Nonterminals, owner)
		r.RulesOwners = append(r.RulesOwners, ownerInfo)
	}
}

func (r *Reader) setRules(rulesLine string) {
	var rulesInfo []rule = make([]rule, 1)
	rulesLine = strings.Trim(rulesLine, " \r\t\n")
	rules := strings.Split(rulesLine, " ")

	for _, oneRule := range rules {
		ruleInfo := rule{
			context: strings.Trim(oneRule, " \r\t\n"),
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

func Contains(set []string, input string) bool {
	for _, value := range set {
		if input == value {
			return true
		}
	}
	return false
}
