package encoder

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Encoder struct {
}

func NewEncoder() *Encoder {
	return &Encoder{}
}

func (e *Encoder) EncodeAddressCommand(command string) string {
	address, err := strconv.Atoi(command[1:]) // .asm files are ascii only, so getting the substring by this method is safe...
	if err != nil {
		fmt.Println("Invalid address:", command)
		panic(err)
	}
	return fmt.Sprintf("%016b", address)
}

func (e *Encoder) EncodeComputeCommand(command string) string {
	// Symbolic C-Command Syntax: dest=comp;jump
	// Binary C-Command Syntax: 111accccccdddjjj

	parts := strings.Split(command, ";")
	destAndComp := parts[0]
	jump := ""
	if len(parts) > 1 {
		jump = parts[1]
	}

	parts = strings.Split(destAndComp, "=")
	dest := parts[0]
	comp := parts[1]

	comp, _ = lookupComp(comp) // acccccc - TODO error handling
	dest, _ = lookupDest(dest) // ddd     - TODO error handling
	jump, _ = lookupJump(jump) // jjj     - TODO error handling

	return fmt.Sprintf("111%s%s%s", comp, dest, jump)
}

func lookupComp(comp string) (string, error) {
	return "xxxxxxx", nil //TODO
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
