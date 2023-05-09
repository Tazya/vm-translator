package logical

import (
	"github.com/tazya/vm-translator/pkg/commands"
	"github.com/tazya/vm-translator/pkg/labels"
)

// Or is x || y, when x is second element from stack, and y is first
type Or struct {
}

func NewOr() commands.Command {
	return &Or{}
}

func (a *Or) GetASMInstructions(labels *labels.Labels) ([]string, error) {
	return []string{
		"// or",
		"@SP",
		"M=M-1",
		"A=M",
		"D=M",
		"@SP",
		"M=M-1",
		"A=M",
		"M=M|D",
		"@SP",
		"M=M+1",
	}, nil
}
