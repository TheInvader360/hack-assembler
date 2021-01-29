package encoder

// Encoder - interface
type Encoder interface {
	EncodeComputeCommand(command string) string
}
