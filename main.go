package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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
	if err != nil {
		fmt.Println("Can't read file:", inputFilename)
		panic(err)
	}
	fmt.Println(string(data))

	outputFilename := strings.Replace(inputFilename, ".asm", ".hack", 1)

	output := []byte("lines\nof\ncode\n")
	err = ioutil.WriteFile(outputFilename, output, 0777)
	if err != nil {
		fmt.Println("Can't write file:", outputFilename)
		panic(err)
	}
}
