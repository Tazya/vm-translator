package arithmetic

import (
	"github.com/tazya/vm-translator/pkg/labels"
	"reflect"
	"testing"
)

func TestAdd_GetASMInstructions(t *testing.T) {
	l := labels.NewLabels()
	expected := []string{
		"// add",
		"@SP",
		"M=M-1",
		"A=M",
		"D=M",
		"@SP",
		"M=M-1",
		"A=M",
		"M=D+M",
		"@SP",
		"M=M+1",
	}

	t.Run("ADD command", func(t *testing.T) {
		a := &Add{}
		instructions, _ := a.GetASMInstructions(l)

		if !reflect.DeepEqual(instructions, expected) {
			t.Errorf("GetASMInstructions() got = %v, want %v", instructions, expected)
		}
	})
}
