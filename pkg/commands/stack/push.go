package stack

import (
	"errors"
	"fmt"
	"github.com/tazya/vm-translator/pkg/commands"
	"github.com/tazya/vm-translator/pkg/memory_segments"
	"strconv"
)

const max15bitValue = 36767

type Push struct {
	segment string
	index   string
}

func NewPush(segment, index string) (commands.Command, error) {
	value, err := strconv.Atoi(index)

	if err != nil || value > max15bitValue {
		return nil, errors.New(fmt.Sprintf("value must be integer, max number: %d", max15bitValue))
	}

	pushCommand := &Push{
		segment: segment,
		index:   index,
	}

	return pushCommand, nil
}

func (p *Push) GetASMInstructions() ([]string, error) {
	if p.segment == "constant" {
		return p.getConstantInstructions(), nil
	}

	segmentLabel, err := memory_segments.GetSegmentLabel(p.segment)

	if err != nil {
		return []string{}, err
	}

	// push local 2
	// @LCL
	// D=M
	// @2
	// A=D+A
	// D=M
	// @SP
	// A=M
	// M=D
	// @SP
	// M=M+1
	return []string{
		fmt.Sprintf("// push %s %s", p.segment, p.index),
		segmentLabel,
		"D=M",
		fmt.Sprintf("@%s", p.index),
		"A=D+A",
		"D=M",
		"@SP",
		"A=M",
		"M=D",
		"@SP",
		"M=M+1",
	}, nil
}

func (p *Push) getConstantInstructions() []string {
	// push constant 17
	// @17
	// D=A
	// @SP
	// A=M
	// M=D
	// @SP
	// M=M+1
	return []string{
		fmt.Sprintf("// push %s %s", p.segment, p.index),
		fmt.Sprintf("@%s", p.index),
		"D=A",
		"@SP",
		"A=M",
		"M=D",
		"@SP",
		"M=M+1",
	}
}
