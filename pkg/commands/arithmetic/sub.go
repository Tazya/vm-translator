package arithmetic

import (
	"github.com/tazya/vm-translator/pkg/commands"
	"github.com/tazya/vm-translator/pkg/labels"
)

// Sub two numbers from stack
type Sub struct {
}

func NewSub() commands.Command {
	return &Sub{}
}

func (a *Sub) GetASMInstructions(labels *labels.Labels) ([]string, error) {
	return []string{
		"// sub",
		"@SP",
		"M=M-1",
		"A=M",
		"D=M",
		"@SP",
		"M=M-1",
		"A=M",
		"M=M-D",
		"@SP",
		"M=M+1",
	}, nil
}
