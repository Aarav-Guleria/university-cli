package university

import (
	"slices"
)

func (db UniversityDB) AverageMarks() float64 {
	if len(db.Students) == 0 {
		return 0
	}

	total := 0
	for _, student := range db.Students {
		total += student.Marks
	}
	return float64(total) / float64(len(db.Students))
}

func (db *UniversityDB) HighestMarks() *Student {
	if len(db.Students) == 0 {
		return nil
	}

	var highest *Student

	for _, student := range db.Students {
		if highest == nil || student.Marks > highest.Marks {
			highest = student
		}
	}

	return highest
}

func (db *UniversityDB) LowestMarks() *Student {
	if len(db.Students) == 0 {
		return nil
	}

	var lowest *Student

	for _, student := range db.Students {
		if lowest == nil || student.Marks < lowest.Marks {
			lowest = student
		}
	}

	return lowest
}

func (db *UniversityDB) StudentsByBranchCount() map[string]int {
	counts := make(map[string]int)

	for _, student := range db.Students {
		counts[student.Branch]++
	}
	return counts
}

func (db *UniversityDB) RankedStudents() []*Student {
	var students []*Student

	for _, student := range db.Students {
		students = append(students, student)
	}
	slices.SortFunc(students, func(a, b *Student) int {
		return b.Marks - a.Marks
	})
	return students
}
