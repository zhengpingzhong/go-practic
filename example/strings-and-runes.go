package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "nihao我是嘴皮子"
	fmt.Println("len:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x", s[i])
	}
	fmt.Println()
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	for idx, runValue := range s {
		fmt.Printf("%#u starts at %d\n", runValue, idx)
	}

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width
		examineRune(runeValue)
	}
	fmt.Println(s[5:])
}
func examineRune(r rune) {
	if r == 'n' {
		fmt.Println("found n")
	} else if r == '我' {
		fmt.Println("found 我")
	}
}
