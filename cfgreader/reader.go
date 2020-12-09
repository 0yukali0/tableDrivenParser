package cfgreader

import (
	"fmt"
	"io/ioutil"
	"os"
)

type cfgAnalysis interface {
	Read(string) string
}

func Read3(filepath string) (fileContext string) {
	f, err := os.Open("file/test")
	if err != nil {
		fmt.Println("read file fail", err)
		return ""
	}
	defer f.Close()

	fd, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("read to fd fail", err)
		return ""
	}

	return string(fd)
}
