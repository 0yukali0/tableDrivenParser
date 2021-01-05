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
	for way != "exit" {
		fmt.Println("input,file,exit")
		fmt.Scanln(&way)
		if way == "file" {
			var filePath string
			fmt.Scanln(&filePath)
			reader.ReadFile(filePath)
		}
	}
}

func reset() {
	reader = r.NewReader()
}
