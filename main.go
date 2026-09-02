package main

import (
	"bufio"
	"fmt"
	"os"

	"example.com/university-clitwo/university"
)

const dataFile = "university.json"

var reader = bufio.NewReader(os.Stdin)

func main() {
	stud, err := university.NewStudent(1, "aarav", 21, "IT", 99)
	if err != nil {
		return &university.ValidationError{Field: "name"}
	}
	fmt.Println(stud)
}
