package logical

import (
	"github.com/tazya/vm-translator/pkg/labels"
	"reflect"
	"testing"
)

func TestOr_GetASMInstructions(t *testing.T) {
	l := labels.NewLabels()
	expected := []string{
		"// or",
		"@SP",
		"M=M-1",
		"A=M",
		"D=M",
		"@SP",
		"M=M-1",
		"A=M",
		"D=M|D",
		"@LABELor1TRUE",
		"!D;JLT",
		"@SP",
		"A=M",
		"M=0",
		"@LABELor1END",
		"0;JMP",
		"(LABELor1TRUE)",
		"@SP",
		"A=M",
		"M=-1",
		"(LABELor1END)",
		"@SP",
		"M=M+1",
	}

	t.Run("LT command", func(t *testing.T) {
		a := &Or{}
		instructions, _ := a.GetASMInstructions(l)

		if !reflect.DeepEqual(instructions, expected) {
			t.Errorf("GetASMInstructions() got = %v, want %v", instructions, expected)
		}
	})
}
