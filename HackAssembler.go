package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"github.com/TheInvader360/hack-assembler/encoder"
	"github.com/TheInvader360/hack-assembler/handler"
	"github.com/TheInvader360/hack-assembler/parser"
	"github.com/TheInvader360/hack-assembler/symboltable"

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

	st := symboltable.NewSymbolTable()
	parser.PopulateSymbolTable(st)

	keys := make([]string, 0, len(st.Map))
	for key := range st.Map {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool { return st.Map[keys[i]] < st.Map[keys[j]] })
	for _, key := range keys {
		fmt.Printf("%s, %d\n", key, st.Map[key])
	}

	encoder := encoder.NewEncoder()
	parser.Translate(encoder)

	outputFilename := strings.Replace(inputFilename, ".asm", ".hack", 1)
	output := []byte(strings.Join(parser.BinaryLines, "\n"))
	err = ioutil.WriteFile(outputFilename, output, 0777)
	handler.FatalError(errors.Wrap(err, fmt.Sprintf("Can't write file: %s", outputFilename)))
}
