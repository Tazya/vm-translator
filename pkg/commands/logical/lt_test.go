package logical

import (
	"github.com/tazya/vm-translator/pkg/labels"
	"reflect"
	"testing"
)

func TestLt_GetASMInstructions(t *testing.T) {
	l := labels.NewLabels()
	expected := []string{
		"// lt",
		"@SP",
		"M=M-1",
		"A=M",
		"D=M",
		"@SP",
		"M=M-1",
		"A=M",
		"D=M-D",
		"@LABELlt1TRUE",
		"D;JLT",
		"@SP",
		"A=M",
		"M=0",
		"@LABELlt1END",
		"0;JMP",
		"(LABELlt1TRUE)",
		"@SP",
		"A=M",
		"M=-1",
		"(LABELlt1END)",
		"@SP",
		"M=M+1",
	}

	t.Run("LT command", func(t *testing.T) {
		a := &Lt{}
		instructions, _ := a.GetASMInstructions(l)

		if !reflect.DeepEqual(instructions, expected) {
			t.Errorf("GetASMInstructions() got = %v, want %v", instructions, expected)
		}
	})
}
