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

func GetSegmentLabel(segmentName, value string) (string, error) {
	if segmentName == "pointer" {
		if value == "0" {
			segmentName = "this"
		} else {
			segmentName = "that"
		}
	}

	label, isExist := segmentLabels[segmentName]

	if !isExist {
		return "", errors.New(fmt.Sprintf("unknown memory segment '%s'", segmentName))
	}

	return label, nil
}
