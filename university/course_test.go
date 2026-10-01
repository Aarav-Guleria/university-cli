package university

import (
	"errors"
	"testing"
)

func TestAddStudent(t *testing.T) {
	student := &Student{
		Person: Person{
			ID:   1,
			Name: "Aarav",
		},
	}

	course := &Course{
		Code: "G101",
		Name: "GoLang",
	}

	err := course.AddStudent(student)
	if err != nil {
		t.Fatalf("unexpected error %v,", err)
	}

	err = course.AddStudent(student)

	if !errors.Is(err, ErrStudentEnrolled) {
		t.Fatalf("expected ErrStudentEnrolled, got %v", err)
	}

	if len(course.Students) != 1 {
		t.Fatalf("expected 1 student, got %d", len(course.Students))
	}
}

func TestAddNilStudent(t *testing.T) {
	course := &Course{
		Code: "G101",
		Name: "GoLang",
	}

	err := course.AddStudent(nil)

	if !errors.Is(err, ErrStudentNil) {
		t.Fatalf("expected ErrStudentNil, got %v", err)
	}
}

func TestRemoveStudentNotEnrolled(t *testing.T) {
	course := &Course{
		Code: "G101",
		Name: "GoLang",
	}

	err := course.RemoveStudent(123)

	if !errors.Is(err, ErrStudentNotEnrolled) {
		t.Fatalf(
			"expected ErrStudentNotEnrolled, got %v",
			err,
		)
	}
}
