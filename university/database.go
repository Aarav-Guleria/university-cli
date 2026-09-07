package university

import "fmt"

type UniversityDB struct {
	Students map[int]*Student
	Teachers map[int]*Teacher
	Courses  map[string]*Course

	nextID int
}

func NewUniversityDB() *UniversityDB {
	return &UniversityDB{
		Students: make(map[int]*Student),
		Teachers: make(map[int]*Teacher),
		Courses:  make(map[string]*Course),
		nextID:   1,
	}
}

type DuplicateIDError struct {
	ID int
}

func (e *DuplicateIDError) Error() string {
	return fmt.Sprintf("ID %d already exists", e.ID)
}

func (db *UniversityDB) generateID() int {
	id := db.nextID
	db.nextID++
	return id
}

func (db *UniversityDB) AddStudent(student *Student) error {
	if student.ID == 0 {
		student.ID = db.generateID()
	} else {
		if _, existing := db.Students[student.ID]; existing {
			return &DuplicateIDError{ID: student.ID}
		}
	}
	db.Students[student.ID] = student
	return nil
}

func (db *UniversityDB) AddTeacher(teacher *Teacher) error {
	if teacher.ID == 0 {
		teacher.ID = db.generateID()
	} else {
		if _, existing := db.Teachers[teacher.ID]; existing {
			return &DuplicateIDError{ID: teacher.ID}
		}
	}
	db.Teachers[teacher.ID] = teacher
	return nil
}

func (db *UniversityDB) AddCourse(course *Course) {
	db.Courses[course.Code] = course
}

func (db *UniversityDB) FindStudent(id int) (*Student, bool) {
	student, ok := db.Students[id]
	return student, ok
}

func (db *UniversityDB) FindTeacher(id int) (*Teacher, bool) {
	teacher, ok := db.Teachers[id]
	return teacher, ok
}

func (db *UniversityDB) FindCourse(code string) (*Course, bool) {
	course, ok := db.Courses[code]
	return course, ok
}

func (db *UniversityDB) RemoveStudent(id int) {
	delete(db.Students, id)
}

func (db *UniversityDB) RemoveTeacher(id int) {
	delete(db.Teachers, id)
}

func (db *UniversityDB) RemoveCourse(code string) {
	delete(db.Courses, code)
}
