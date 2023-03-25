package memory_segments

import (
	"errors"
	"fmt"
)

var segmentLabels = map[string]string{
	"local":    "@LCL",
	"argument": "@ARG",
	"this":     "@THIS",
	"that":     "@THAT",
}

func GetSegmentLabel(segmentName string) (string, error) {
	label, isExist := segmentLabels[segmentName]

	if !isExist {
		return "", errors.New(fmt.Sprintf("unknown memory segment '%s'", segmentName))
	}

	return label, nil
}
