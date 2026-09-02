package university

import "fmt"

type ValidationError struct {
	Field string
	Msg   string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed on %s: message: %s", e.Field, e.Msg)
}
