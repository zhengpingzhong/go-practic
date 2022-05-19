package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("base(p)", filepath.Base(p))

	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("C:/dir/file"))
	filename := "config.json"
	ext := filepath.Ext(filename)
	fmt.Println(ext)
	fmt.Println(strings.TrimSuffix(filename, ext))

}
