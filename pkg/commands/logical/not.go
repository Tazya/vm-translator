package logical

import (
	"github.com/tazya/vm-translator/pkg/commands"
	"github.com/tazya/vm-translator/pkg/labels"
)

// Not is !x, false == !true
type Not struct {
}

func NewNot() commands.Command {
	return &Not{}
}

func (a *Not) GetASMInstructions(_ *labels.Labels) ([]string, error) {
	return []string{
		"// not",
		"@SP",
		"M=M-1",
		"A=M",
		"M=!M",
		"@SP",
		"M=M+1",
	}, nil
}
