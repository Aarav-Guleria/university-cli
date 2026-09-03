package university

import "errors"

type Course struct {
	Code string
	Name string

	Students []*Student
	Teacher  *Teacher
}

func NewCousre(code string, name string, teacher *Teacher) (*Course, error) {
	if code == "" {
		return nil, &ValidationError{Field: "code", Msg: "cannot be empty"}
	}

	if name == "" {
		return nil, &ValidationError{Field: "name", Msg: "cannot be empty"}
	}

	return &Course{
		Code:     code,
		Name:     name,
		Teacher:  teacher,
		Students: nil,
	}, nil
}

func (c *Course) AddStudent(student *Student) error {
	if student == nil {
		return errors.New("student cannot be nil")
	}

	for _, existing := range c.Students {
		if existing.ID == student.ID {
			return errors.New("student is already added")
		}
	}

	c.Students = append(c.Students, student)
	return nil
}

func (c *Course) RemoveStudent(studentID int) error {
	for i, student := range c.Students {
		if student.ID == studentID {
			c.Students = append(c.Students[:i], c.Students[i+1:]...)
			return nil
		}
	}
	return errors.New("student is not enrolled")
}
