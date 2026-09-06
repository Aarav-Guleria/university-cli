package university

import (
	"encoding/json"
	"os"
)

type studentFile struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Branch string `json:"branch"`
	Marks  int    `json:"marks"`
}

type teacherFile struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Department string `json:"department"`
	Salary     int    `json:"salary"`
}

type courseFile struct {
	Code       string `json:"code"`
	Name       string `json:"name"`
	TeacherID  int    `json:"teacher_id"`
	StudentIDs []int  `json:"student_ids"`
}

type databaseFile struct {
	Students []studentFile `json:"students"`
	Teachers []teacherFile `json:"teachers"`
	Courses  []courseFile  `json:"courses"`
	NextID   int           `json:"next_id"`
}

func (db *UniversityDB) Save(filename string) error {
	var file databaseFile

	for _, student := range db.Students {
		file.Students = append(file.Students, studentFile{
			ID:     student.ID,
			Name:   student.Name,
			Age:    student.Age,
			Branch: student.Branch,
			Marks:  student.Marks,
		})
	}

	for _, teacher := range db.Teachers {
		file.Teachers = append(file.Teachers, teacherFile{
			ID:         teacher.ID,
			Name:       teacher.Name,
			Age:        teacher.Age,
			Department: teacher.Department,
			Salary:     teacher.Salary,
		})
	}

	for _, course := range db.Courses {
		cf := courseFile{
			Code: course.Code,
			Name: course.Name,
		}

		if course.Teacher != nil {
			cf.TeacherID = course.Teacher.ID
		}

		for _, student := range course.Students {
			cf.StudentIDs = append(cf.StudentIDs, student.ID)
		}

		file.Courses = append(file.Courses, cf)
	}

	data, err := json.MarshalIndent(file, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0o644)
}

func Load(filename string) (*UniversityDB, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var file databaseFile

	if err := json.Unmarshal(data, &file); err != nil {
		return nil, err
	}

	db := NewUniversityDB()

	db.nextID = file.NextID

	for _, saved := range file.Students {
		student := &Student{
			Person: Person{
				ID:   saved.ID,
				Name: saved.Name,
				Age:  saved.Age,
			},
			Branch: saved.Branch,
			Marks:  saved.Marks,
		}

		db.Students[student.ID] = student
	}

	for _, saved := range file.Teachers {
		teacher := &Teacher{
			Person: Person{
				ID:   saved.ID,
				Name: saved.Name,
				Age:  saved.Age,
			},
			Department: saved.Department,
			Salary:     saved.Salary,
		}
		db.Teachers[teacher.ID] = teacher
	}

	for _, saved := range file.Courses {
		course := &Course{
			Code: saved.Code,
			Name: saved.Name,
		}

		if saved.TeacherID != 0 {
			if teacher, ok := db.Teachers[saved.TeacherID]; ok {
				course.Teacher = teacher
			}
		}

		for _, studentID := range saved.StudentIDs {
			if student, ok := db.Students[studentID]; ok {
				course.Students = append(course.Students, student)
			}
		}
		db.Courses[course.Code] = course
	}
	return db, nil
}
