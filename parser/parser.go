package parser

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/TheInvader360/hack-assembler/encoder"
	"github.com/TheInvader360/hack-assembler/symbols"
)

// Parser - struct
type Parser struct {
	SourceLines []string
	BinaryLines []string
}

// NewParser - returns a pointer to new parser
func NewParser() *Parser {
	return &Parser{}
}

// Sanitize - populate SourceLines with sanitized source data (comments and whitespace removed)
func (p *Parser) Sanitize(data []byte) {
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Split(line, "//")[0]
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			p.SourceLines = append(p.SourceLines, line)
		}
	}
}

// PopulateSymbolsMapWithLables - populate symbols map with all "labels" (i.e "(xxx)" symbols)
func (p *Parser) PopulateSymbolsMapWithLables(s symbols.Symbols) {
	value := 0
	for _, command := range p.SourceLines {
		if command[0] == '(' {
			label := strings.TrimLeft(command, "(")
			label = strings.TrimRight(label, ")")
			s.Put(label, value)
			fmt.Println(label, ":", value)
		} else {
			value++
		}
	}
}

// Translate - populate BinaryLines with translated SourceLines
func (p *Parser) Translate(s symbols.Symbols, encoder encoder.Encoder) {
	n := 16
	for _, command := range p.SourceLines {
		if command[0] == '@' {
			value, err := strconv.Atoi(command[1:]) // .asm files are ascii only, so getting the substring by this method is safe...
			if err != nil {
				// couldn't parse as an int, so it's safe to assume that this is a @symbol command
				symbol := command[1:]
				if s.Contains(symbol) {
					value = s.Get(symbol)
				} else {
					s.Put(symbol, n)
					value = n
					n++
				}
				fmt.Println(symbol, value)
			}
			p.BinaryLines = append(p.BinaryLines, fmt.Sprintf("%016b", value))
		} else if command[0] == '(' {
			fmt.Println(command)
		} else {
			p.BinaryLines = append(p.BinaryLines, encoder.EncodeComputeCommand(command))
		}
	}
}
