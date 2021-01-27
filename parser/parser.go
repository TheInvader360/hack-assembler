package parser

import (
	"bufio"
	"bytes"
	"strings"

	"github.com/TheInvader360/hack-assembler/encoder"
)

type Parser struct {
	SourceLines []string
	BinaryLines []string
}

func NewParser() *Parser {
	return &Parser{}
}

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

func (p *Parser) Translate(encoder *encoder.Encoder) {
	for _, command := range p.SourceLines {
		if command[0] == '@' {
			p.BinaryLines = append(p.BinaryLines, encoder.EncodeAddressCommand(command))
		} else {
			p.BinaryLines = append(p.BinaryLines, encoder.EncodeComputeCommand(command))
		}
	}
}
