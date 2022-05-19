package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	os.Setenv("foo", "1")
	fmt.Println("foo:", os.Getenv("foo"))
	fmt.Println("bar:", os.Getenv("bar"))

	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}
