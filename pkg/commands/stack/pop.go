package stack

import (
	"errors"
	"fmt"
	"github.com/tazya/vm-translator/pkg/commands"
	"github.com/tazya/vm-translator/pkg/labels"
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

	if segment == "temp" && value > tempMaxIndex {
		return nil, errors.New(fmt.Sprintf("reach temp varibales limit: %d", tempMaxIndex))
	}

	if segment == "pointer" && (value != 0 && value != 1) {
		return nil, errors.New(fmt.Sprintf("value for pointer must be 0 or 1"))
	}

	command := &Pop{
		segment:   segment,
		index:     index,
		classname: classname,
	}

	return command, nil
}

func (p *Pop) GetASMInstructions(labels *labels.Labels) ([]string, error) {
	if p.segment == "constant" {
		return []string{}, errors.New("syntax error. can not write to constant")
	}

	if p.segment == "static" {
		return p.getStaticInstructions(), nil
	}

	if p.segment == "temp" {
		return p.getTempInstructions(), nil
	}

	if p.segment == "pointer" {
		return p.getPointerInstructions()
	}

	segmentLabel, err := memory_segments.GetSegmentLabel(p.segment, p.index)

	if err != nil {
		return []string{}, err
	}

	return []string{
		fmt.Sprintf("// pop %s %s", p.segment, p.index),
		segmentLabel, // temp = segment + index
		"D=M",
		fmt.Sprintf("@%s", p.index),
		"D=D+A",
		"@13",
		"M=D",
		"@SP", // @SP--
		"M=M-1",
		"A=M", // RAM[temp] = POP STACK
		"D=M",
		"@13",
		"A=M",
		"M=D",
	}, nil
}

func (p *Pop) getStaticInstructions() []string {
	return []string{
		fmt.Sprintf("// pop static %s", p.index),
		"@SP", // @SP--
		"M=M-1",
		"A=M",
		"D=M",
		fmt.Sprintf("@%s.%s", p.classname, p.index),
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

func (p *Pop) getPointerInstructions() ([]string, error) {
	label, err := memory_segments.GetSegmentLabel(p.segment, p.index)

	if err != nil {
		return []string{}, err
	}

	return []string{
		fmt.Sprintf("// pop pointer %s", p.index),
		"@SP", // @SP--
		"M=M-1",
		"A=M",
		"D=M",
		label,
		"M=D",
	}, nil
}
