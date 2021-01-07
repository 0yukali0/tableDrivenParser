package myreader

import (
	"bufio"
	"os"
	"strings"
)

func (r *Reader) ReadRule() error {
	in := bufio.NewReader(os.Stdin)
	input, _ := in.ReadString('\n')
	input = strings.Trim(input, "\t\r ")
	indexOfNumterminal := strings.Index(input, " ") + 1
	input = input[indexOfNumterminal:]

	return nil
}
