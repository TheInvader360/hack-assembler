package mock

import (
	"github.com/stretchr/testify/mock"
)

// Encoder - Mock struct (implements encoder.Encoder interface)
type Encoder struct {
	mock.Mock
}

// EncodeComputeCommand - Mock
func (e *Encoder) EncodeComputeCommand(command string) string {
	args := e.Called(command)
	return args.String(0)
}
