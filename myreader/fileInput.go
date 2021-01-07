package myreader

import (
	"errors"
	"io/ioutil"
	"os"
)

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
