package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/phil-mansfield/glox/glox1/scanner"
	"github.com/phil-mansfield/glox/glox1/error"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Fprintf(os.Stderr, "Usage: glox1 [script]\n")
		os.Exit(64)
	} else if len(os.Args) == 2 {
		RunFile(os.Args[1])
	} else {
		RunPrompt()
	}
}

func RunPrompt() {
	scanner := bufio.NewScanner(os.Stdin)
	
	for {
		fmt.Print("> ")
		if !scanner.Scan() { break }
		Run(string(scanner.Bytes()))
		error.HadError = false
	}
}

func RunFile(file string) {
	text, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not open target file: %s", err.Error())
		os.Exit(64)
	}

	Run(string(text))

	if error.HadError { os.Exit(65) }
}

func Run(text string) {
	scanner := scanner.NewScanner(text)
	tokens := scanner.Scan()

	for i := range tokens {
		fmt.Println(tokens[i])
	}
}
