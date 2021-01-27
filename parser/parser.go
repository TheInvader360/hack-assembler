package parser

import (
	"bufio"
	"bytes"
	"strings"

	"github.com/TheInvader360/hack-assembler/encoder"
	"github.com/TheInvader360/hack-assembler/symboltable"
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

// PopulateSymbolTable - populate symbol table with all "labels" (i.e "(xxx)" symbols)
func (p *Parser) PopulateSymbolTable(st *symboltable.SymbolTable) {
	value := 0
	for _, command := range p.SourceLines {
		if command[0] == '(' {
			label := strings.TrimLeft(command, "(")
			label = strings.TrimRight(label, ")")
			st.Put(label, value)
		} else {
			value++
		}
	}
}

// Translate - populate BinaryLines with translated SourceLines
func (p *Parser) Translate(encoder *encoder.Encoder) {
	for _, command := range p.SourceLines {
		if command[0] == '@' {
			//			p.BinaryLines = append(p.BinaryLines, encoder.EncodeAddressCommand(command))
		} else {
			//			p.BinaryLines = append(p.BinaryLines, encoder.EncodeComputeCommand(command))
		}
	}
}
