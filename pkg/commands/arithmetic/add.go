package arithmetic

import (
	"github.com/tazya/vm-translator/pkg/commands"
	"github.com/tazya/vm-translator/pkg/labels"
)

type Add struct {
}

func NewAdd() commands.Command {
	return &Add{}
}

func (a *Add) GetASMInstructions(labels *labels.Labels) ([]string, error) {
	return []string{
		"// add",
		"@SP",
		"M=M-1",
		"A=M",
		"D=M",
		"@SP",
		"M=M-1",
		"A=M",
		"M=D+M",
		"@SP",
		"M=M+1",
	}, nil
}
