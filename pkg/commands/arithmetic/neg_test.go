package arithmetic

import (
	"github.com/tazya/vm-translator/pkg/labels"
	"reflect"
	"testing"
)

func TestNeg_GetASMInstructions(t *testing.T) {
	l := labels.NewLabels()
	expected := []string{
		"// neg",
		"@SP",
		"M=M-1",
		"A=M",
		"M=-M",
		"@SP",
		"M=M+1",
	}

	t.Run("NEG command", func(t *testing.T) {
		s := &Neg{}
		instructions, _ := s.GetASMInstructions(l)

		if !reflect.DeepEqual(instructions, expected) {
			t.Errorf("GetASMInstructions() got = %v, want %v", instructions, expected)
		}
	})
}
