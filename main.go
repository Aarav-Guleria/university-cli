package main

import (
	"fmt"

	"github.com/Aarav-Guleria/UniversityCLI/university"
)

func main() {
	db := university.NewUniversityDB()

	aarav, err := university.NewStudent(0, "Aarav", 21, "IT", 98)
	if err != nil {
		fmt.Println("error creating student:", err)
		return
	}
	db.AddStudent(aarav)

	jacob, err := university.NewStudent(0, "Jacob", 20, "CSE", 99)
	if err != nil {
		fmt.Println("error creating student:", err)
		return
	}
	db.AddStudent(jacob)

	teacher, err := university.NewTeacher(0, "Dr. Rathore", 45, "IT", 120000)
	if err != nil {
		fmt.Println("error creating teacher:", err)
		return
	}
	db.AddTeacher(teacher)

	course, err := university.NewCourse("G101", "GoLang", teacher)
	if err != nil {
		fmt.Println("error creating course:", err)
		return
	}
	db.AddCourse(course)

	if err := course.AddStudent(aarav); err != nil {
		fmt.Println("error enrolling student:", err)
		return
	}
	if err := course.AddStudent(jacob); err != nil {
		fmt.Println("error enrolling student:", err)
		return
	}

	found, ok := db.FindStudent(aarav.ID)
	if !ok {
		fmt.Println("student not found in db")
		return
	} else {
		fmt.Println("Found via DB:", found)
		fmt.Println()
		fmt.Println("Assigned student ID:", aarav.ID)
		fmt.Println("Student Name:", aarav.Name)
		fmt.Println()
		fmt.Println("Assigned teacher ID:", teacher.ID)
		fmt.Println("Teacher Name:", teacher.Name)
	}

	fmt.Println(db.SearchStudentsByName("Aa"))
	fmt.Println(db.AverageMarks())
	fmt.Println(db.HighestMarks().Name, db.HighestMarks().ID, db.HighestMarks().Person)
	fmt.Println(db.LowestMarks().ID)
	fmt.Println(db.RankedStudents())
	fmt.Println(db.StudentsByBranchCount())
	counts := db.StudentsByBranchCount()
	for branch, count := range counts {
		fmt.Printf("branch: %s, counts: %d\n", branch, count)
	}

	db.RemoveStudent(2)
	fmt.Println(db.Students)

	fmt.Println(db.Teachers)
	db.RemoveTeacher(3)
	fmt.Println(db.Teachers)

	fmt.Println(db.FindCourse("G101"))
	db.RemoveCourse("G101")
	fmt.Println(db.FindCourse("G101"))
}
