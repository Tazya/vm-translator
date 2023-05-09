package logical

import (
	"fmt"
	"github.com/tazya/vm-translator/pkg/commands"
	"github.com/tazya/vm-translator/pkg/labels"
)

// Eq last number from stack
type Eq struct {
}

func NewEq() commands.Command {
	return &Eq{}
}

func (a *Eq) GetASMInstructions(labels *labels.Labels) ([]string, error) {
	labelPrefix := labels.GetNextLabel("eq")
	labelTrue := labelPrefix + "TRUE"
	labelEnd := labelPrefix + "END"

	return []string{
		"// eq",
		"@SP",
		"M=M-1",
		"A=M",
		"D=M",
		"@SP",
		"M=M-1",
		"A=M",
		"D=M-D",
		fmt.Sprintf("@%s", labelTrue),
		"D;JEQ",
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
