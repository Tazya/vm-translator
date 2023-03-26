package stack

import (
	"errors"
	"fmt"
	"github.com/tazya/vm-translator/pkg/commands"
	"github.com/tazya/vm-translator/pkg/memory_segments"
	"strconv"
)

type Push struct {
	segment   string
	index     string
	classname string
}

func NewPush(segment, index, classname string) (commands.Command, error) {
	value, err := strconv.Atoi(index)

	if err != nil || value > max15bitValue {
		return nil, errors.New(fmt.Sprintf("value must be integer, max number: %d", max15bitValue))
	}

	if segment == "static" && value > staticVariablesLimit {
		return nil, errors.New(fmt.Sprintf("reach static varibales limit: %d", staticVariablesLimit))
	}

	pushCommand := &Push{
		segment:   segment,
		index:     index,
		classname: classname,
	}

	return pushCommand, nil
}

func (p *Push) GetASMInstructions() ([]string, error) {
	if p.segment == "constant" {
		return p.getConstantInstructions(), nil
	}

	if p.segment == "static" {
		return p.getStaticInstructions(), nil
	}

	segmentLabel, err := memory_segments.GetSegmentLabel(p.segment)

	if err != nil {
		return []string{}, err
	}

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
	return []string{
		fmt.Sprintf("// push constant %s", p.index),
		fmt.Sprintf("@%s", p.index),
		"D=A",
		"@SP",
		"A=M",
		"M=D",
		"@SP",
		"M=M+1",
	}
}

func (p *Push) getStaticInstructions() []string {
	return []string{
		fmt.Sprintf("// push static %s", p.index),
		fmt.Sprintf("@%s.%s", p.classname, p.index),
		"D=M",
		"@SP",
		"A=M",
		"M=D",
		"@SP",
		"M=M+1",
	}
}
