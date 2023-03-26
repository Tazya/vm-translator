package stack

import (
	"errors"
	"github.com/tazya/vm-translator/pkg/commands"
)

const (
	max15bitValue        = 36767
	staticVariablesLimit = 249 // 16-255 memory registers
	tempIndexLimit       = 7
	tempBaseAddress      = 5
)

var stackCommands = map[string]func(segment, index, classname string) (commands.Command, error){
	"push": NewPush,
	"pop":  NewPop,
}

func IsStackCommand(commandName string) bool {
	_, isExist := stackCommands[commandName]

	return isExist
}

func GetCommand(fields []string, classname string) (commands.Command, error) {
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

	return commandConstructor(segment, index, classname)
}
