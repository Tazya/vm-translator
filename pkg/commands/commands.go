package commands

type Command interface {
	GetASMInstructions() ([]string, error)
}
