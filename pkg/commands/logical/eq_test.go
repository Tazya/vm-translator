package logical

import (
	"github.com/tazya/vm-translator/pkg/labels"
	"reflect"
	"testing"
)

func TestEq_GetASMInstructions(t *testing.T) {
	l := labels.NewLabels()
	expected := []string{
		"// eq",
		"@SP",
		"M=M-1",
		"A=M",
		"D=M",
		"@SP",
		"M=M-1",
		"A=M",
		"D=M-D",
		"@LABELeq1TRUE",
		"D;JEQ",
		"@SP",
		"A=M",
		"M=0",
		"@LABELeq1END",
		"0;JMP",
		"(LABELeq1TRUE)",
		"@SP",
		"A=M",
		"M=-1",
		"(LABELeq1END)",
		"@SP",
		"M=M+1",
	}

	t.Run("Eq command", func(t *testing.T) {
		a := &Eq{}
		instructions, _ := a.GetASMInstructions(l)

		if !reflect.DeepEqual(instructions, expected) {
			t.Errorf("GetASMInstructions() got = %v, want %v", instructions, expected)
		}
	})
}
