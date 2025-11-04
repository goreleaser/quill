package event

import (
	"fmt"
)

type ErrBadPayload struct {
	Type  string
	Field string
	Value interface{}
}

func (e *ErrBadPayload) Error() string {
	return fmt.Sprintf("event='%s' has bad event payload field='%v': '%+v'", string(e.Type), e.Field, e.Value)
}

func newPayloadErr(t string, field string, value interface{}) error {
	return &ErrBadPayload{
		Type:  t,
		Field: field,
		Value: value,
	}
}

func checkEventType(actual, expected string) error {
	if actual != expected {
		return newPayloadErr(expected, "Type", actual)
	}
	return nil
}
