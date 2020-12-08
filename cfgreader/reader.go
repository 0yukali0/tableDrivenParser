package cfgreader

import (
	"fmt"
	"io/ioutil"
	"os"
)

func Read3() string {
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
