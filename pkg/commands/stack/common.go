package stack

import (
	"errors"
	"github.com/tazya/vm-translator/pkg/commands"
)

var stackCommands = map[string]func(segment, index string) (commands.Command, error){
	"push": NewPush,
}

func IsStackCommand(commandName string) bool {
	_, isExist := stackCommands[commandName]

	return isExist
}

func GetCommand(fields []string) (commands.Command, error) {
	if len(fields) != 3 {
		return nil, errors.New("syntax error. Stack command must have 2 operands")
	}

	commandName := fields[0]
	commandConstructor, isExist := stackCommands[commandName]

	if !isExist {
		return nil, errors.New("command not found")
	}

	segment := fields[1]
	index := fields[2]

	return commandConstructor(segment, index)
}
