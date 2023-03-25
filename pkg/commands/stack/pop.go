package stack

import (
	"errors"
	"fmt"
	"github.com/tazya/vm-translator/pkg/commands"
	"github.com/tazya/vm-translator/pkg/memory_segments"
	"strconv"
)

type Pop struct {
	segment string
	index   string
}

func NewPop(segment, index string) (commands.Command, error) {
	value, err := strconv.Atoi(index)

	if err != nil || value > max15bitValue {
		return nil, errors.New(fmt.Sprintf("value must be integer, max number: %d", max15bitValue))
	}

	command := &Pop{
		segment: segment,
		index:   index,
	}

	return command, nil
}

func (p *Pop) GetASMInstructions() ([]string, error) {
	if p.segment == "constant" {
		return []string{}, errors.New("syntax error. can not write to constant")
	}

	segmentLabel, err := memory_segments.GetSegmentLabel(p.segment)

	if err != nil {
		return []string{}, err
	}

	return []string{
		fmt.Sprintf("// pop %s %s", p.segment, p.index),
		segmentLabel, // temp = segment + index
		"D=A",
		fmt.Sprintf("@%s", p.index),
		"D=D+A",
		"@5",
		"M=D",
		"@SP", // @SP--
		"M=M-1",
		"A=M", // RAM[temp] = POP STACK
		"D=M",
		"@5",
		"A=M",
		"M=D",
	}, nil
}
