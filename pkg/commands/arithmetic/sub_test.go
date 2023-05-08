package arithmetic

import (
	"github.com/tazya/vm-translator/pkg/labels"
	"reflect"
	"testing"
)

func TestSub_GetASMInstructions(t *testing.T) {
	l := labels.NewLabels()
	expected := []string{
		"// sub",
		"@SP",
		"M=M-1",
		"A=M",
		"D=M",
		"@SP",
		"M=M-1",
		"A=M",
		"M=M-D",
		"@SP",
		"M=M+1",
	}

	t.Run("SUB command", func(t *testing.T) {
		s := &Sub{}
		instructions, _ := s.GetASMInstructions(l)

		if !reflect.DeepEqual(instructions, expected) {
			t.Errorf("GetASMInstructions() got = %v, want %v", instructions, expected)
		}
	})
}
