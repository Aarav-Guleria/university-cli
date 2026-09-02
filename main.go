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
}
