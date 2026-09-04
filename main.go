package main

import (
	"fmt"

	"github.com/Aarav-Guleria/UniversityCLI/university"
)

func main() {
	db := university.NewUniversityDB()

	student, err := university.NewStudent(0, "Aarav", 21, "IT", 98)
	if err != nil {
		fmt.Println("error creating student:", err)
		return
	}
	db.AddStudent(student)

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
	fmt.Println(course)
	db.AddCoures(course)

	if err := course.AddStudent(student); err != nil {
		fmt.Println("error enrolling student:", err)
		return
	}

	found, ok := db.FindStudent(student.ID)
	if !ok {
		fmt.Println("student not found in db")
		return
	} else {
		fmt.Println("Found via DB:", found)
		fmt.Println()
		fmt.Println("Assigned student ID:", student.ID)
		fmt.Println("Student Name:", student.Name)
		fmt.Println()
		fmt.Println("Assigned teacher ID:", teacher.ID)
		fmt.Println("Teacher Name:", teacher.Name)
		fmt.Println()
	}
}
