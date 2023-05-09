package logical

import (
	"github.com/tazya/vm-translator/pkg/labels"
	"reflect"
	"testing"
)

func TestAnd_GetASMInstructions(t *testing.T) {
	l := labels.NewLabels()
	expected := []string{
		"// and",
		"@SP",
		"M=M-1",
		"A=M",
		"D=M",
		"@SP",
		"M=M-1",
		"A=M",
		"D=M&D",
		"@LABELand1TRUE",
		"!D;JLT",
		"@SP",
		"A=M",
		"M=0",
		"@LABELand1END",
		"0;JMP",
		"(LABELand1TRUE)",
		"@SP",
		"A=M",
		"M=-1",
		"(LABELand1END)",
		"@SP",
		"M=M+1",
	}

	t.Run("LT command", func(t *testing.T) {
		a := &And{}
		instructions, _ := a.GetASMInstructions(l)

		if !reflect.DeepEqual(instructions, expected) {
			t.Errorf("GetASMInstructions() got = %v, want %v", instructions, expected)
		}
	})
}
