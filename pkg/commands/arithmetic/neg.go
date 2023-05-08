package arithmetic

import (
	"github.com/tazya/vm-translator/pkg/commands"
	"github.com/tazya/vm-translator/pkg/labels"
)

// Neg last number from stack
type Neg struct {
}

func NewNeg() commands.Command {
	return &Neg{}
}

func (a *Neg) GetASMInstructions(labels *labels.Labels) ([]string, error) {
	return []string{
		"// neg",
		"@SP",
		"M=M-1",
		"A=M",
		"M=-M",
		"@SP",
		"M=M+1",
	}, nil
}
