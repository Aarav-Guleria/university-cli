package university

import (
	"errors"
	"fmt"
)

var (
	ErrStudentNil         = errors.New("student cannot be nil")
	ErrStudentEnrolled    = errors.New("student is already enrolled")
	ErrStudentNotEnrolled = errors.New("student is not enrolled")
)

type ValidationError struct {
	Field string
	Msg   string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed on %s: %s", e.Field, e.Msg)
}

type DuplicateCourseError struct {
	Code string
}

func (e *DuplicateCourseError) Error() string {
	return fmt.Sprintf("course %s already exists", e.Code)
}
