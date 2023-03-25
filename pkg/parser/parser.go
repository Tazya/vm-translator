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

	if isPush(trimmedLine) {
		return stack.NewPush(trimmedLine)
	}

	if isAdd(trimmedLine) {
		return &arithmetic.Add{}, nil
	}

	return nil, errors.New(fmt.Sprintf("Unknown command '%s'", l))
}

func isComment(s string) bool {
	return strings.HasPrefix(s, "//")
}

func isPush(s string) bool {
	return strings.HasPrefix(s, "push")
}

func isAdd(s string) bool {
	return strings.HasPrefix(s, "add")
}
