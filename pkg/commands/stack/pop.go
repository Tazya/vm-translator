package stack

import (
	"errors"
	"fmt"
	"github.com/tazya/vm-translator/pkg/commands"
	"github.com/tazya/vm-translator/pkg/memory_segments"
	"strconv"
)

type Pop struct {
	segment   string
	index     string
	classname string
}

func NewPop(segment, index, classname string) (commands.Command, error) {
	value, err := strconv.Atoi(index)

	if err != nil || value > max15bitValue {
		return nil, errors.New(fmt.Sprintf("value must be integer, max number: %d", max15bitValue))
	}

	if segment == "static" && value > staticVariablesLimit {
		return nil, errors.New(fmt.Sprintf("reach static varibales limit: %d", staticVariablesLimit))
	}

	command := &Pop{
		segment:   segment,
		index:     index,
		classname: classname,
	}

	return command, nil
}

func (p *Pop) GetASMInstructions() ([]string, error) {
	if p.segment == "constant" {
		return []string{}, errors.New("syntax error. can not write to constant")
	}

	if p.segment == "static" {
		return p.getStaticInstructions(), nil
	}

	if p.segment == "temp" {
		return p.getTempInstructions(), nil
	}

	segmentLabel, err := memory_segments.GetSegmentLabel(p.segment, p.index)

	if err != nil {
		return []string{}, err
	}

	// TODO @5 can collide with pop temp command. To fix we need to find a safety memory address to hold pointer
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

func (p *Pop) getStaticInstructions() []string {
	return []string{
		fmt.Sprintf("// pop static %s", p.index),
		fmt.Sprintf("@%s.%s", p.classname, p.index),
		"D=M",
		"@SP", // @SP--
		"M=M-1",
		"A=M",
		"M=D",
	}
}

func (p *Pop) getTempInstructions() []string {
	intIndex, _ := strconv.Atoi(p.index)

	return []string{
		fmt.Sprintf("// pop temp %s", p.index),
		"@SP", // @SP--
		"M=M-1",
		"A=M",
		"D=M",
		fmt.Sprintf("@%d", tempBaseAddress+intIndex),
		"M=D",
	}
}
