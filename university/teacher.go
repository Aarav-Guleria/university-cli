package university

import "fmt"

type Teacher struct {
	Person
	Department string
	Salary     int
}

func NewTeacher(id int, name string, age int, department string, salary int) (*Teacher, error) {
	if name == "" {
		return nil, &ValidationError{Field: "name", Msg: "cannot be empty"}
	}

	if age <= 0 {
		return nil, &ValidationError{Field: "age", Msg: "must be positive"}
	}

	if department == "" {
		return nil, &ValidationError{Field: "department", Msg: "cannot be empty"}
	}

	if salary < 0 {
		return nil, &ValidationError{Field: "salary", Msg: "cannot be negative"}
	}

	return &Teacher{
		Person: Person{
			ID:   id,
			Name: name,
			Age:  age,
		},
		Department: department,
		Salary:     salary,
	}, nil
}

func (t *Teacher) GetRole() string {
	return "teacher"
}

func (t *Teacher) String() string {
	return fmt.Sprintf("%s | %s dept", t.Name, t.Department)
}
