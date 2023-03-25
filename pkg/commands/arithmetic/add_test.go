package arithmetic

import (
	"reflect"
	"testing"
)

func TestAdd_GetASMInstructions(t *testing.T) {
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
		instructions, _ := a.GetASMInstructions()

		if !reflect.DeepEqual(instructions, expected) {
			t.Errorf("GetASMInstructions() got = %v, want %v", instructions, expected)
		}
	})
}
