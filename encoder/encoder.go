package encoder

import (
	"fmt"
	"strconv"
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
	return "C-Command" //TODO
}

func (e *Encoder) EncodeComp(component string) (string, error) {
	return "TODO", nil
}

func (e *Encoder) EncodeDest(component string) (string, error) {
	return "TODO", nil
}

func (e *Encoder) EncodeJump(component string) (string, error) {
	return "TODO", nil
}
