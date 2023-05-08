package labels

import "fmt"

type Labels struct {
	counters map[string]int
}

func NewLabels() *Labels {
	l := Labels{
		counters: map[string]int{},
	}

	return &l
}

func (l *Labels) GetNextLabel(labelType string) string {
	counter, hasLabel := l.counters[labelType]

	if hasLabel {
		counter++
		l.counters[labelType] = counter

		return fmt.Sprintf("LABEL%s%d", labelType, counter)
	}

	l.counters[labelType] = 1

	return fmt.Sprintf("LABEL%s%d", labelType, 1)
}
