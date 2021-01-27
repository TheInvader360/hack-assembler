package parser

import (
	"bufio"
	"bytes"
	"strings"
)

type Parser struct {
	CurrentLine int
	Lines       []string
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(data []byte) {
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		line := scanner.Text()
		lineNoComments := strings.Split(line, "//")[0]
		lineSanitized := strings.TrimSpace(lineNoComments)
		if len(lineSanitized) > 0 {
			p.Lines = append(p.Lines, lineSanitized)
		}
	}
}

func (p *Parser) HasMoreCommands() bool {
	return true //TODO
}

func (p *Parser) Advance() {
	//TODO
}
