package university

import "testing"

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

	if err == nil {
		t.Fatal("expected duplicate enrollment to fail")
	}

	if len(course.Students) != 1 {
		t.Fatalf(
			"expected 1 student, got %d",
			len(course.Students),
		)
	}
}
