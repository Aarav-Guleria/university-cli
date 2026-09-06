package university

import "testing"

func TestGrade(t *testing.T) {
	tests := []struct {
		marks    int
		expected string
	}{
		{95, "A"},
		{85, "B"},
		{75, "C"},
		{65, "D"},
		{40, "F"},
	}

	for _, test := range tests {
		student := &Student{
			Marks: test.marks,
		}

		actual := student.Grade()

		if actual != test.expected {
			t.Errorf(
				"Grade(%d) = %s; expected %s",
				test.marks,
				actual,
				test.expected,
			)
		}
	}
}

func TestUpdateMarks(t *testing.T) {
	student := &Student{
		Marks: 50,
	}

	err := student.UpdateMarks(90)
	if err != nil {
		t.Fatal("expected valid marks to work")
	}

	if student.Marks != 90 {
		t.Fatalf(
			"expected marks to be 90 but got %d",
			student.Marks,
		)
	}

	err = student.UpdateMarks(150)
	if err == nil {
		t.Fatal("expected invalid marks to return an error")
	}
}
