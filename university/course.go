package university

type Course struct {
	Code string `json:"code"`
	Name string `json:"name"`

	Students []*Student `json:"-"`
	Teacher  *Teacher   `json:"-"`
}

func NewCourse(code string, name string, teacher *Teacher) (*Course, error) {
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
		return ErrStudentNil
	}

	for _, existing := range c.Students {
		if existing.ID == student.ID {
			return ErrStudentEnrolled
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
	return ErrStudentNotEnrolled
}
