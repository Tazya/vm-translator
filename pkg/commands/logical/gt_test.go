package logical

import (
	"github.com/tazya/vm-translator/pkg/labels"
	"reflect"
	"testing"
)

func TestGt_GetASMInstructions(t *testing.T) {
	l := labels.NewLabels()
	expected := []string{
		"// gt",
		"@SP",
		"M=M-1",
		"A=M",
		"D=M",
		"@SP",
		"M=M-1",
		"A=M",
		"D=M-D",
		"@LABELgt1TRUE",
		"D;JGT",
		"@SP",
		"A=M",
		"M=0",
		"@LABELgt1END",
		"0;JMP",
		"(LABELgt1TRUE)",
		"@SP",
		"A=M",
		"M=-1",
		"(LABELgt1END)",
		"@SP",
		"M=M+1",
	}

	t.Run("GT command", func(t *testing.T) {
		a := &Gt{}
		instructions, _ := a.GetASMInstructions(l)

		if !reflect.DeepEqual(instructions, expected) {
			t.Errorf("GetASMInstructions() got = %v, want %v", instructions, expected)
		}
	})
}
