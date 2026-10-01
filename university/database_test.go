package university

import (
	"errors"
	"testing"
)

func TestRemoveStudentRemovesEnrollment(t *testing.T) {
	db := NewUniversityDB()

	student, err := NewStudent(
		0,
		"Aarav",
		20,
		"CSE",
		80,
	)
	if err != nil {
		t.Fatal(err)
	}

	if err := db.AddStudent(student); err != nil {
		t.Fatal(err)
	}

	course, err := NewCourse(
		"CS101",
		"Go",
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	if err := course.AddStudent(student); err != nil {
		t.Fatal(err)
	}

	if err := db.AddCourse(course); err != nil {
		t.Fatal(err)
	}

	db.RemoveStudent(student.ID)

	if _, ok := db.FindStudent(student.ID); ok {
		t.Fatal("expected student to be removed from database")
	}

	if len(course.Students) != 0 {
		t.Fatalf(
			"expected student enrollment to be removed, got %d students",
			len(course.Students),
		)
	}
}

func TestAddDuplicateCourse(t *testing.T) {
	db := NewUniversityDB()

	first, err := NewCourse("CS101", "Go", nil)
	if err != nil {
		t.Fatal(err)
	}

	second, err := NewCourse("CS101", "Advanced Go", nil)
	if err != nil {
		t.Fatal(err)
	}

	if err := db.AddCourse(first); err != nil {
		t.Fatal(err)
	}

	err = db.AddCourse(second)

	var duplicateErr *DuplicateCourseError

	if !errors.As(err, &duplicateErr) {
		t.Fatalf(
			"expected DuplicateCourseError, got %v",
			err,
		)
	}

	if duplicateErr.Code != "CS101" {
		t.Fatalf(
			"expected course code CS101, got %s",
			duplicateErr.Code,
		)
	}

	course, ok := db.FindCourse("CS101")
	if !ok {
		t.Fatal("expected course to exist")
	}

	if course.Name != "Go" {
		t.Fatalf(
			"expected original course to remain, got %s",
			course.Name,
		)
	}
}
