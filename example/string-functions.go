package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {
	p("Contains: ", s.Contains("test", "es"))
	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foooo", "o", "0", -1))
	p("Replace:   ", s.Replace("foooo", "o", "0", 2))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
}

//Contains:  true
//Contains:   true
//Count:      2
//HasPrefix:  true
//HasSuffix:  true
//Index:      1
//Join:       a-b
//Repeat:     aaaaa
//Replace:    f0000
//Replace:    f00oo
//Split:      [a b c d e]
//ToLower:    test
//ToUpper:    TEST
