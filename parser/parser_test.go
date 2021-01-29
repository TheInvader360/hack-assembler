package parser

import (
	"strings"
	"testing"

	encodermock "github.com/TheInvader360/hack-assembler/encoder/mock"
	symbolsmock "github.com/TheInvader360/hack-assembler/symbols/mock"

	"github.com/stretchr/testify/assert"
)

func TestSanitize(t *testing.T) {
	d := []byte("// comment\n@2\nD=A\n\n@3//\nD=D+A // comment comment comment\n@0\n\n\nM=D\n")
	p := NewParser()
	p.Sanitize(d)
	assert.Equal(t, "@2\nD=A\n@3\nD=D+A\n@0\nM=D", strings.Join(p.SourceLines, "\n"))
}

func TestPopulateSymbolsMapWithLables(t *testing.T) {
	p := NewParser()
	p.SourceLines = append(p.SourceLines, "A-NOLABEL")
	p.SourceLines = append(p.SourceLines, "(B-LABEL)")
	p.SourceLines = append(p.SourceLines, "C-NOLABEL")
	p.SourceLines = append(p.SourceLines, "(D-LABEL)")
	s := symbolsmock.Symbols{}
	s.On("Put", "B-LABEL", 1).Return()
	s.On("Put", "D-LABEL", 2).Return()
	p.PopulateSymbolsMapWithLables(&s)
	s.AssertExpectations(t)
	s.AssertNumberOfCalls(t, "Put", 2)
}

func TestTranslate(t *testing.T) {
	p := NewParser()
	p.SourceLines = append(p.SourceLines, "@25")
	p.SourceLines = append(p.SourceLines, "@YES")
	p.SourceLines = append(p.SourceLines, "@NO1")
	p.SourceLines = append(p.SourceLines, "@NO2")
	p.SourceLines = append(p.SourceLines, "(XY)")
	p.SourceLines = append(p.SourceLines, "D=D-M;JLE")
	s := symbolsmock.Symbols{}
	s.On("Contains", "YES").Return(true)
	s.On("Get", "YES").Return(1)
	s.On("Contains", "NO1").Return(false)
	s.On("Put", "NO1", 16).Return()
	s.On("Contains", "NO2").Return(false)
	s.On("Put", "NO2", 17).Return()
	e := encodermock.Encoder{}
	e.On("EncodeComputeCommand", "D=D-M;JLE").Return("1111010011010110")
	p.Translate(&s, &e)
	s.AssertExpectations(t)
	e.AssertExpectations(t)
	s.AssertNumberOfCalls(t, "Contains", 3)
	s.AssertNumberOfCalls(t, "Get", 1)
	s.AssertNumberOfCalls(t, "Put", 2)
	e.AssertNumberOfCalls(t, "EncodeComputeCommand", 1)
}
