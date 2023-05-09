package logical

import (
	"fmt"
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
	labelPrefix := labels.GetNextLabel("or")
	labelTrue := labelPrefix + "TRUE"
	labelEnd := labelPrefix + "END"

	return []string{
		"// or",
		"@SP",
		"M=M-1",
		"A=M",
		"D=M",
		"@SP",
		"M=M-1",
		"A=M",
		"D=M|D",
		fmt.Sprintf("@%s", labelTrue),
		"!D;JLT",
		"@SP",
		"A=M",
		"M=0",
		fmt.Sprintf("@%s", labelEnd),
		"0;JMP",
		fmt.Sprintf("(%s)", labelTrue),
		"@SP",
		"A=M",
		"M=-1",
		fmt.Sprintf("(%s)", labelEnd),
		"@SP",
		"M=M+1",
	}, nil
}
