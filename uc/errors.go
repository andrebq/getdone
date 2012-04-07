package uc

import (
	"fmt"
)

// Used to inform that a call into a UseCase object wasn't valid because the internal state of 
// object didn't allow the call
type InvalidState string

func newInvalidState(format string, args ...interface{}) InvalidState {
	return InvalidState(fmt.Sprintf(format, args...))
}
func (i InvalidState) Error() string {
	return string(i)
}
