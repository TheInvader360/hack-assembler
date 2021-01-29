package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	encoderimpl "github.com/TheInvader360/hack-assembler/encoder/impl"
	"github.com/TheInvader360/hack-assembler/handler"
	"github.com/TheInvader360/hack-assembler/parser"
	symbolsimpl "github.com/TheInvader360/hack-assembler/symbols/impl"

	"github.com/pkg/errors"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Missing file parameter")
		return
	}
	inputFilename := os.Args[1]

	if !strings.HasSuffix(inputFilename, ".asm") {
		fmt.Println("Expected an asm file (*.asm)")
		os.Exit(1)
	}

	data, err := ioutil.ReadFile(inputFilename)
	handler.FatalError(errors.Wrap(err, fmt.Sprintf("Can't read file: %s", inputFilename)))

	parser := parser.NewParser()
	parser.Sanitize(data)

	s := symbolsimpl.NewSymbols()
	s.Seed()
	fmt.Println("First Pass Label Symbols:")
	parser.PopulateSymbolsMapWithLables(s)

	e := encoderimpl.NewEncoder()
	fmt.Println("\nSecond Pass Variable Symbols:")
	parser.Translate(s, e)

	outputFilename := strings.Replace(inputFilename, ".asm", ".hack", 1)
	output := []byte(strings.Join(parser.BinaryLines, "\n"))
	err = ioutil.WriteFile(outputFilename, output, 0777)
	handler.FatalError(errors.Wrap(err, fmt.Sprintf("Can't write file: %s", outputFilename)))
}
