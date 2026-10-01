package university

import (
	"path/filepath"
	"testing"
)

func TestSaveLoadKeepsNextID(t *testing.T) {
	db := NewUniversityDB()

	student, err := NewStudent(
		0,
		"Aarav",
		20,
		"CSE",
		80,
	)
	if err != nil {
		t.Fatal(err)
	}

	if err := db.AddStudent(student); err != nil {
		t.Fatal(err)
	}

	path := filepath.Join(t.TempDir(), "database.json")

	if err := db.Save(path); err != nil {
		t.Fatal(err)
	}

	loaded, err := Load(path)
	if err != nil {
		t.Fatal(err)
	}

	nextStudent, err := NewStudent(
		0,
		"Rahul",
		21,
		"ECE",
		70,
	)
	if err != nil {
		t.Fatal(err)
	}

	if err := loaded.AddStudent(nextStudent); err != nil {
		t.Fatal(err)
	}

	if nextStudent.ID <= student.ID {
		t.Fatalf("expected new ID > %d, got %d", student.ID, nextStudent.ID)
	}
}
