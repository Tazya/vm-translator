package logical

import (
	"github.com/tazya/vm-translator/pkg/commands"
	"github.com/tazya/vm-translator/pkg/labels"
)

// And is x && y, when x is second element from stack, and y is first
type And struct {
}

func NewAnd() commands.Command {
	return &And{}
}

func (a *And) GetASMInstructions(labels *labels.Labels) ([]string, error) {
	return []string{
		"// and",
		"@SP",
		"M=M-1",
		"A=M",
		"D=M",
		"@SP",
		"M=M-1",
		"A=M",
		"M=M&D",
		"@SP",
		"M=M+1",
	}, nil
}
