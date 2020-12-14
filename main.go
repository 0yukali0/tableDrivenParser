package main

import (
	"flag"
	"fmt"
)

var filePath = flag.String("FILEPATH", "", "Input your CFG path")

func main() {
	flag.Parse()
	for i := 0; i != flag.NArg(); i++ {
		fmt.Printf("arg[%d]=%s\n", i, flag.Arg(i))
	}

	fmt.Println("Path=", *filePath)
}
