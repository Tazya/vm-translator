package arithmetic

import (
	"errors"
	"github.com/tazya/vm-translator/pkg/commands"
)

var arithmeticCommands = map[string]func() commands.Command{
	"add": NewAdd,
}

func IsArithmeticCommand(commandName string) bool {
	_, isExist := arithmeticCommands[commandName]

	return isExist
}

func GetCommand(fields []string) (commands.Command, error) {
	if len(fields) != 1 {
		return nil, errors.New("syntax error. Stack command must not have operands")
	}

	commandName := fields[0]
	commandConstructor, isExist := arithmeticCommands[commandName]

	if !isExist {
		return nil, errors.New("command not found")
	}

	return commandConstructor(), nil
}
