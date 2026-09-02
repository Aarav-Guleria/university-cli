package university

import "fmt"

type Member interface {
	GetName() string
	GetRole() string
}

type Person struct {
	ID   int
	Name string
	Age  int
}

func (p Person) GetName() string {
	return p.Name
}

type Student struct {
	Person
	Branch string
	Marks  int
}

func NewStudent(id int, name string, age int, branch string, marks int) (*Student, error) {
	if name == "" {
		return nil, &ValidationError{Field: "name", Msg: "name cannot be empty"}
	}

	if age <= 0 {
		return nil, &ValidationError{Field: "age", Msg: "invalid age"}
	}

	if branch == "" {
		return nil, &ValidationError{Field: "branch", Msg: "branch cannot be empty"}
	}

	if marks < 0 || marks > 100 {
		return nil, &ValidationError{Field: "marks", Msg: "invalid marks"}
	}

	return &Student{
		Person: Person{
			ID:   id,
			Name: name,
			Age:  age,
		},
		Branch: branch,
		Marks:  marks,
	}, nil
}

func (s *Student) Grade() string {
	switch v := s.Marks; {
	case v >= 90:
		return "A"
	case v >= 80:
		return "B"
	case v >= 70:
		return "C"
	case v >= 60:
		return "D"
	default:
		return "F"
	}
}

func (s *Student) UpdateMarks(newMarks int) error {
	if newMarks <= 0 || newMarks >= 100 {
		return &ValidationError{Field: "marks", Msg: "invalid marks"}
	}

	s.Marks = newMarks
	return nil
}

func (s *Student) GetRole() string {
	return "student"
}

func (s *Student) String() string {
	return fmt.Sprintf("%s | %s | %d | Grade %s", s.Name, s.Branch, s.Marks, s.Grade())
}
