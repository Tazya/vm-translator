package parser

import (
	"errors"
	"fmt"
	"github.com/tazya/vm-translator/pkg/commands"
	"github.com/tazya/vm-translator/pkg/commands/arithmetic"
	"github.com/tazya/vm-translator/pkg/commands/stack"
	"strings"
)

func ParseLine(l string) (commands.Command, error) {
	trimmedLine := strings.Trim(l, " ")

	if isComment(trimmedLine) || trimmedLine == "" {
		return nil, nil
	}

	fields := strings.Fields(l)
	commandName := fields[0]

	if stack.IsStackCommand(commandName) {
		return stack.GetCommand(fields)
	}

	if arithmetic.IsArithmeticCommand(commandName) {
		return arithmetic.GetCommand(fields)
	}

	return nil, errors.New(fmt.Sprintf("Unknown command '%s'", l))
}

func isComment(s string) bool {
	return strings.HasPrefix(s, "//")
}
