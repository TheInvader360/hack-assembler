package encoder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeComputeCommand(t *testing.T) {
	type test struct {
		input, expectedResult string
	}
	tests := []test{
		{input: "0", expectedResult: "1110101010000000"},
		{input: "M=1", expectedResult: "1110111111001000"},
		{input: "-1;JGT", expectedResult: "1110111010000001"},
		{input: "D", expectedResult: "1110001100000000"},
		{input: "D=A", expectedResult: "1110110000010000"},
		{input: "D=D+A", expectedResult: "1110000010010000"},
		{input: "!D;JEQ", expectedResult: "1110001101000010"},
		{input: "!A", expectedResult: "1110110001000000"},
		{input: "MD=-D", expectedResult: "1110001111011000"},
		{input: "-A;JGE", expectedResult: "1110110011000011"},
		{input: "D+1", expectedResult: "1110011111000000"},
		{input: "A=A+1", expectedResult: "1110110111100000"},
		{input: "D-1;JLT", expectedResult: "1110001110000100"},
		{input: "A-1", expectedResult: "1110110010000000"},
		{input: "AM=D+A", expectedResult: "1110000010101000"},
		{input: "D-A;JNE", expectedResult: "1110010011000101"},
		{input: "A-D", expectedResult: "1110000111000000"},
		{input: "AD=D&A", expectedResult: "1110000000110000"},
		{input: "D|A;JLE", expectedResult: "1110010101000110"},
		{input: "M", expectedResult: "1111110000000000"},
		{input: "AMD=!M", expectedResult: "1111110001111000"},
		{input: "-M;JMP", expectedResult: "1111110011000111"},
		{input: "M+1", expectedResult: "1111110111000000"},
		{input: "M=M-1;JGT", expectedResult: "1111110010001001"},
		{input: "M=D", expectedResult: "1110001100001000"},
		{input: "D+M", expectedResult: "1111000010000000"},
		{input: "D=D-M;JLE", expectedResult: "1111010011010110"},
		{input: "M-D", expectedResult: "1111000111000000"},
		{input: "AMD=D&M;JMP", expectedResult: "1111000000111111"},
		{input: "D|M;JMP", expectedResult: "1111010101000111"},
	}
	encoder := NewEncoder()
	for _, tc := range tests {
		result := encoder.EncodeComputeCommand(tc.input)
		assert.Equal(t, tc.expectedResult, result)
	}
}

func TestLookupCompInvalid(t *testing.T) {
	result, err := lookupComp("aaa")
	assert.Equal(t, "", result)
	assert.EqualError(t, err, "Invalid Comp: aaa")
}

func TestLookupDestInvalid(t *testing.T) {
	result, err := lookupDest("bbb")
	assert.Equal(t, "", result)
	assert.EqualError(t, err, "Invalid Dest: bbb")
}

func TestLookupCompJump(t *testing.T) {
	result, err := lookupJump("ccc")
	assert.Equal(t, "", result)
	assert.EqualError(t, err, "Invalid Jump: ccc")
}
