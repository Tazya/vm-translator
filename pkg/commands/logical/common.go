package logical

import (
	"errors"
	"github.com/tazya/vm-translator/pkg/commands"
)

var logicalCommands = map[string]func() commands.Command{
	"eq":  NewEq, // Equals
	"gt":  NewGt, // Greater than
	"lt":  NewLt, // Less than
	"and": NewAnd,
	"or":  NewOr,
	"not": NewNot,
}

func IsLogicalCommand(commandName string) bool {
	_, isExist := logicalCommands[commandName]

	return isExist
}

func GetCommand(fields []string) (commands.Command, error) {
	if len(fields) != 1 {
		return nil, errors.New("syntax error. Stack command must not have operands")
	}

	commandName := fields[0]
	commandConstructor, isExist := logicalCommands[commandName]

	if !isExist {
		return nil, errors.New("command not found")
	}

	return commandConstructor(), nil
}
