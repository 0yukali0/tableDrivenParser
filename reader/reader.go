package reader

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

type CFGAnalysis interface {
	Read(string)
	SetLAndRHS()
	SetTerminals()
}

type FileReader struct {
	context      string
	err          error
	Nonterminals map[string]LHSInfo
	Terminals    map[string]RHSInfo
	Tokens       []RHSInfo
}

func (r *FileReader) Read(filepath string) {
	f, err := os.Open(filepath)
	if err != nil {
		r.err = errors.New("File path is invalid")
	}
	defer f.Close()

	fd, err := ioutil.ReadAll(f)
	if err != nil {
		r.err = errors.New("Read to file fail")
	}
	r.context = string(fd)
}

//<LHS>[ ]*>[ ]*<RHS>([\t\n ]<RHS>)
func (r *FileReader) SetLAndRHS() {
	lines := strings.Split(r.context, "\n")
	//var assignedLHS LHSInfo
	for _, context := range lines {
		//assignedLHSName, find := findLHS(context)
		_, find := FindLHS(context)
		if find {

		}
	}
	//r.SetTerminals()
}

func FindLHS(line string) (string, bool) {
	return "", false
}
