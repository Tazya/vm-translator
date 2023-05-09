package logical

import (
	"fmt"
	"github.com/tazya/vm-translator/pkg/commands"
	"github.com/tazya/vm-translator/pkg/labels"
)

// Lt is x < y, when x is second element from stack, and y is first
type Lt struct {
}

func NewLt() commands.Command {
	return &Lt{}
}

func (a *Lt) GetASMInstructions(labels *labels.Labels) ([]string, error) {
	labelPrefix := labels.GetNextLabel("lt")
	labelTrue := labelPrefix + "TRUE"
	labelEnd := labelPrefix + "END"

	return []string{
		"// lt",
		"@SP",
		"M=M-1",
		"A=M",
		"D=M",
		"@SP",
		"M=M-1",
		"A=M",
		"D=M-D",
		fmt.Sprintf("@%s", labelTrue),
		"D;JLT",
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
