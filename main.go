package main

import (
	"flag"
	"fmt"
)

//定义flags
var inputName = flag.String("name", "XuChao", "Input your name")
var inputAge = flag.Int("age", 25, "Input your age")
var inputGender = flag.String("gender", "boy", "Input your gender")

func main() {
	flag.Parse() //flag解析
	for i := 0; i != flag.NArg(); i++ {
		fmt.Printf("arg[%d]=%s\n", i, flag.Arg(i))
	}

	fmt.Println("name=", *inputName)
	fmt.Println("age=", *inputAge)
	fmt.Println("gender=", *inputGender)
}
