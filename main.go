package main

import (
	"fmt"

	"github.com/Aarav-Guleria/UniversityCLI/university"
)

func main() {
	student, err := university.NewStudent(1, "Aarav", 21, "IT", 98)
	if err != nil {
		fmt.Println("error creating student:", err)
		return
	}
	fmt.Println(student)

	teacher, err := university.NewTeacher(1, "Dr. Rathore", 45, "IT", 120000)
	if err != nil {
		fmt.Println("error creating teacher:", err)
		return
	}
	fmt.Println(teacher)

	course, err := university.NewCousre("G01", "GoLang", teacher)
	if err != nil {
		fmt.Println("error creating course", err)
	}
	// check
	fmt.Println(course)

	if err := course.AddStudent(student); err != nil {
		fmt.Println("error enrolling student:", err)
		return
	}
	fmt.Println("\nCourse:", course.Name)
	fmt.Println("Teacher:", course.Teacher.Name)
	fmt.Println("Enrolled:")
	for _, s := range course.Students {
		fmt.Println(" :-", s.Name)
	}
}
