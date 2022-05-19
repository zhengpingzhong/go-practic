package main

import (
	"flag"
	"fmt"
)

var (
	intflag    int
	boolflag   bool
	stringflag string
)

func init() {
	fmt.Println("first init func")
	flag.IntVar(&intflag, "intflag", 0, "int flag value")
	flag.BoolVar(&boolflag, "boolflag", false, "bool flag value")
	flag.StringVar(&stringflag, "stringflag", "default", "string flag value")
}

func init() {
	fmt.Println("second init func")
}

func init() {
	fmt.Println("third init func")
}

func main() {
	flag.Parse()
	fmt.Println("int flag", intflag)
	fmt.Println("bool flag", boolflag)
	fmt.Println("string flag", stringflag)
}
