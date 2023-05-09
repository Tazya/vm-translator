package logical

import (
	"github.com/tazya/vm-translator/pkg/labels"
	"reflect"
	"testing"
)

func TestNot_GetASMInstructions(t *testing.T) {
	l := labels.NewLabels()
	expected := []string{
		"// not",
		"@SP",
		"M=M-1",
		"A=M",
		"M=!M",
		"@SP",
		"M=M+1",
	}

	t.Run("LT command", func(t *testing.T) {
		a := &Not{}
		instructions, _ := a.GetASMInstructions(l)

		if !reflect.DeepEqual(instructions, expected) {
			t.Errorf("GetASMInstructions() got = %v, want %v", instructions, expected)
		}
	})
}
