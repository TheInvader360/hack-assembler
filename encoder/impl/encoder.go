package impl

import (
	"fmt"
	"strings"

	"github.com/TheInvader360/hack-assembler/handler"

	"github.com/pkg/errors"
)

// Encoder - struct (implements encoder.Encoder interface)
type Encoder struct {
}

// NewEncoder - returns a pointer to new encoder
func NewEncoder() *Encoder {
	return &Encoder{}
}

// EncodeComputeCommand - returns a binary encoded C-Command
func (e *Encoder) EncodeComputeCommand(command string) string {
	// Symbolic C-Command Syntax: dest=comp;jump
	// Binary C-Command Syntax: 111accccccdddjjj

	parts := strings.Split(command, ";")
	destAndComp := parts[0]
	jump := ""
	if len(parts) > 1 {
		jump = parts[1]
	}

	dest := ""
	comp := ""
	parts = strings.Split(destAndComp, "=")
	if len(parts) > 1 {
		dest = parts[0]
		comp = parts[1]
	} else {
		comp = parts[0]
	}

	comp, err := lookupComp(comp) // acccccc
	handler.FatalError(err)

	dest, err = lookupDest(dest) // ddd
	handler.FatalError(err)

	jump, err = lookupJump(jump) // jjj
	handler.FatalError(err)

	return fmt.Sprintf("111%s%s%s", comp, dest, jump)
}

func lookupComp(comp string) (string, error) {
	switch comp {
	case "0":
		return "0101010", nil
	case "1":
		return "0111111", nil
	case "-1":
		return "0111010", nil
	case "D":
		return "0001100", nil
	case "A":
		return "0110000", nil
	case "!D":
		return "0001101", nil
	case "!A":
		return "0110001", nil
	case "-D":
		return "0001111", nil
	case "-A":
		return "0110011", nil
	case "D+1":
		return "0011111", nil
	case "A+1":
		return "0110111", nil
	case "D-1":
		return "0001110", nil
	case "A-1":
		return "0110010", nil
	case "D+A":
		return "0000010", nil
	case "D-A":
		return "0010011", nil
	case "A-D":
		return "0000111", nil
	case "D&A":
		return "0000000", nil
	case "D|A":
		return "0010101", nil
	case "M":
		return "1110000", nil
	case "!M":
		return "1110001", nil
	case "-M":
		return "1110011", nil
	case "M+1":
		return "1110111", nil
	case "M-1":
		return "1110010", nil
	case "D+M":
		return "1000010", nil
	case "D-M":
		return "1010011", nil
	case "M-D":
		return "1000111", nil
	case "D&M":
		return "1000000", nil
	case "D|M":
		return "1010101", nil
	}
	return "", errors.New("Invalid Comp: " + comp)
}

func lookupDest(dest string) (string, error) {
	switch dest {
	case "":
		return "000", nil // value not stored
	case "M":
		return "001", nil // RAM[A]
	case "D":
		return "010", nil // D Register
	case "MD":
		return "011", nil // RAM[A] and D Register
	case "A":
		return "100", nil // A Register
	case "AM":
		return "101", nil // A Register and RAM[A]
	case "AD":
		return "110", nil // A Register and D Register
	case "AMD":
		return "111", nil // A Register and RAM[A] and D Register
	}
	return "", errors.New("Invalid Dest: " + dest)
}

func lookupJump(jump string) (string, error) {
	switch jump {
	case "":
		return "000", nil // no jump
	case "JGT":
		return "001", nil // if out > 0 jump
	case "JEQ":
		return "010", nil // if out = 0 jump
	case "JGE":
		return "011", nil // if out >= 0 jump
	case "JLT":
		return "100", nil // if out < 0 jump
	case "JNE":
		return "101", nil // if out != 0 jump
	case "JLE":
		return "110", nil // if out < 0 jump
	case "JMP":
		return "111", nil // unconditional jump
	}
	return "", errors.New("Invalid Jump: " + jump)
}
