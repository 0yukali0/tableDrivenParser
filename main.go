package main

import (
	"fmt"
	r "myreader"
)

var (
	reader *r.Reader
	trans  *t.translator
	way    string
)

func main() {
	way = ""
	for way != "exit" {
		input()
	}
}

func input() {
	fmt.Println("input,file,exit")
	fmt.Scanln(&way)
	switch {
	case way == "file":
		var filePath string
		fmt.Scanln(&filePath)
		reset()
		reader.ReadFile(filePath)
	case way == "input":
		reset()
		reader.ReadRule()
	}
}

func reset() {
	reader = r.NewReader()
}
