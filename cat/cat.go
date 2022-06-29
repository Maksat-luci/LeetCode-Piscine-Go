package main

import (
	"io"
	"io/ioutil"
	"os"
	"github.com/01-edu/z01"
)

func main() {
	file_name := os.Args[1:]
	length := 0
	var fileContent []byte
	var err error
	for range file_name {
		length++
	}
	if length > 2 {
		return
	}
	if length == 0 {
		io.Copy(os.Stdout, os.Stdin)
	}
	for i, _ := range file_name {
		fileContent, err = ioutil.ReadFile(file_name[i])
		if err != nil {
			for _, x := range string(err.Error()) {
				z01.PrintRune(x)
			}

		}
		for _, x := range fileContent {
			z01.PrintRune(rune(x))
		}
		z01.PrintRune('\n')
	}
}
