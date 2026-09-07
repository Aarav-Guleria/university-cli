package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Aarav-Guleria/UniversityCLI/university"
)

const dataFile = "university.json"

var reader = bufio.NewReader(os.Stdin)

func main() {
	db, err := university.Load(dataFile)

	if err != nil {
		db = university.NewUniversityDB()
		fmt.Println("Starting with a new database.")
	} else {
		fmt.Println("Loaded database from", dataFile)
	}

	for {
		printMenu()

		choice := readInt("Choose an option: ")

		fmt.Println()

		switch choice {
		case 1:
			addStudent(db)
		case 2:
			listStudents(db)
		case 3:
			findStudent(db)
		case 4:
			searchStudents(db)
		case 5:
			updateMarks(db)
		case 6:
			deleteStudent(db)
		case 7:
			showStatistics(db)
		case 8:
			addTeacher(db)
		case 9:
			addCourse(db)
		case 10:
			enrollStudent(db)
		case 11:
			listCourses(db)
		case 12:
			saveDatabase(db)
		case 0:
			saveDatabase(db)
			fmt.Println("END")
			return
		default:
			fmt.Println("Invalid choice")
		}
		fmt.Println()
	}
}

func printMenu() {
	fmt.Println("================================")
	fmt.Println("       UNIVERSITY SYSTEM")
	fmt.Println("================================")

	fmt.Println("1.  Add Student")
	fmt.Println("2.  List Students")
	fmt.Println("3.  Find Student")
	fmt.Println("4.  Search Students")
	fmt.Println("5.  Update Marks")
	fmt.Println("6.  Delete Student")
	fmt.Println("7.  Statistics")
	fmt.Println("8.  Add Teacher")
	fmt.Println("9.  Add Course")
	fmt.Println("10. Enroll Student")
	fmt.Println("11. List Courses")
	fmt.Println("12. Save Database")
	fmt.Println("0.  Exit")

	fmt.Println("================================")
}

func readString(prompt string) string {
	fmt.Print(prompt)

	input, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	return strings.TrimSpace(input)
}

func readInt(prompt string) int {
	for {
		input := readString(prompt)
		number, err := strconv.Atoi(input)
		if err == nil {
			return number
		}
		fmt.Println("Please enter a valid number.")
	}
}

func addStudent(db *university.UniversityDB) {
	fmt.Println("=====ADD STUDENT=====")

	name := readString("Name: ")
	age := readInt("Age: ")
	branch := readString("Branch: ")
	marks := readInt("Marks: ")

	student, err := university.NewStudent(
		0,
		name,
		age,
		branch,
		marks,
	)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	err = db.AddStudent(student)
	if err != nil {
		return
	}

	fmt.Println("Student added successfully.")
	fmt.Println("Assigned ID:", student.ID)
}

func listStudents(db *university.UniversityDB) {
	fmt.Println("=====STUDENTS=====")

	if len(db.Students) == 0 {
		fmt.Println("No students found")
		return
	}

	for _, student := range db.RankedStudents() {
		fmt.Printf(
			"ID: %d | %s | Age %d | %s | Marks: %d | Grade: %s\n",
			student.ID,
			student.Name,
			student.Age,
			student.Branch,
			student.Marks,
			student.Grade(),
		)
	}
}

func findStudent(db *university.UniversityDB) {
	id := readInt("Enter student ID: ")

	student, ok := db.FindStudent(id)

	if !ok {
		fmt.Println("Student not found.")
		return
	}

	fmt.Println("Student found:")
	fmt.Println("ID:", student.ID)
	fmt.Println("Name:", student.Name)
	fmt.Println("Age:", student.Age)
	fmt.Println("Branch:", student.Branch)
	fmt.Println("Marks:", student.Marks)
	fmt.Println("Grade:", student.Grade())
}

func searchStudents(db *university.UniversityDB) {
	query := readString("Search name: ")

	results := db.SearchStudentsByName(query)

	if len(results) == 0 {
		fmt.Println("No students found")
		return
	}

	fmt.Println("Search results:")

	for _, student := range results {
		fmt.Printf(
			"ID: %d | %s | %s %d\n",
			student.ID,
			student.Name,
			student.Branch,
			student.Marks,
		)
	}
}

func updateMarks(db *university.UniversityDB) {
	id := readInt("Student ID: ")

	student, ok := db.FindStudent(id)

	if !ok {
		fmt.Println("Student not found.")
		return
	}

	newMarks := readInt("New marks: ")

	err := student.UpdateMarks(newMarks)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Marks updated successfully")
}

func deleteStudent(db *university.UniversityDB) {
	id := readInt("Student ID: ")

	student, ok := db.FindStudent(id)

	if !ok {
		fmt.Println("Student not found")
		return
	}

	answer := strings.ToLower(
		readString(
			fmt.Sprintf(
				"Are you sure you want to delete %s (y/n)",
				student.Name,
			),
		),
	)

	if answer != "y" {
		fmt.Println("Delete cancelled")
		return
	}

	db.RemoveStudent(id)

	fmt.Println("Student deleted")
}

func showStatistics(db *university.UniversityDB) {
	fmt.Println("=====STATISTICS=====")

	if len(db.Students) == 0 {
		fmt.Println("No students available")
		return
	}

	fmt.Printf("Total students: %d\n", len(db.Students))

	fmt.Printf("Average marks: %.2f\n", db.AverageMarks())

	highest := db.HighestMarks()

	if highest != nil {
		fmt.Printf("Highest: %s (%d)\n", highest.Name, highest.Marks)
	}

	lowest := db.LowestMarks()

	if lowest != nil {
		fmt.Printf("Lowest: %s (%d)\n", lowest.Name, lowest.Marks)
	}

	fmt.Println("\nStudents per branch:")

	counts := db.StudentsByBranchCount()

	for branch, count := range counts {
		fmt.Printf("%s: %d\n", branch, count)
	}

	for branch, count := range counts {
		fmt.Printf("%s: %d\n", branch, count)
	}
}

func addTeacher(db *university.UniversityDB) {
	fmt.Println("=====ADD TEACHER=====")

	name := readString("Name: ")
	age := readInt("Age: ")
	department := readString("Department: ")
	salary := readInt("Salary: ")

	teacher, err := university.NewTeacher(0, name, age, department, salary)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = db.AddTeacher(teacher)
	if err != nil {
		return
	}

	fmt.Println("Teacher added successfully")
	fmt.Println("Assigned ID:", teacher.ID)
}

func addCourse(db *university.UniversityDB) {
	fmt.Println("=====ADD COURSE=====")

	code := readString("Course code: ")
	name := readString("Course name: ")

	var teacher *university.Teacher

	assignTeacher := strings.ToLower(readString("Assign Teacher (y/n):"))

	if assignTeacher == "y" {
		id := readInt("Teacher ID: ")
		foundTeacher, ok := db.FindTeacher(id)

		if !ok {
			fmt.Println("Teacher not found")
			return
		}
		teacher = foundTeacher
	}

	course, err := university.NewCourse(code, name, teacher)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	db.AddCourse(course)
	fmt.Println("Course added successfully")
}

func enrollStudent(db *university.UniversityDB) {
	fmt.Println("=====ENROLL STUDENT=====")

	courseCode := readString("Course code: ")
	course, ok := db.FindCourse(courseCode)
	if !ok {
		fmt.Println("Course not found")
		return
	}

	studentID := readInt("Student ID: ")
	student, ok := db.FindStudent(studentID)
	if !ok {
		fmt.Println("Student not found")
		return
	}

	err := course.AddStudent(student)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("Student enrolled successfully")
}

func listCourses(db *university.UniversityDB) {
	fmt.Println("=====COURSES=====")
	if len(db.Courses) == 0 {
		fmt.Println("No courses found")
		return
	}

	for _, course := range db.Courses {
		fmt.Println("___________________________")
		fmt.Println("Code:", course.Code)
		fmt.Println("Name", course.Name)

		if course.Teacher != nil {
			fmt.Println("Teacher:", course.Teacher.Name)
		} else {
			fmt.Println("Teacher: Unassigned")
		}

		fmt.Println("Students enrolled:", len(course.Students))

		for _, student := range course.Students {
			fmt.Printf(" :- %s (ID %d)\n", student.Name, student.ID)
		}
	}
}

func saveDatabase(db *university.UniversityDB) {
	err := db.Save(dataFile)
	if err != nil {
		fmt.Println("Failed to save:", err)
		return
	}
	fmt.Println("Database saved successfully")
}
