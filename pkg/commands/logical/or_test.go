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
		"M=M|D",
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
