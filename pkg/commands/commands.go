package commands

import "github.com/tazya/vm-translator/pkg/labels"

type Command interface {
	GetASMInstructions(labels *labels.Labels) ([]string, error)
}
