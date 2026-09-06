package university

import "strings"

func (db *UniversityDB) SearchStudentsByName(query string) []*Student {
	query = strings.ToLower(query)

	var results []*Student

	for _, student := range db.Students {
		name := strings.ToLower(student.Name)

		if strings.Contains(name, query) {
			results = append(results, student)
		}
	}
	return results
}

func (db *UniversityDB) StudentsByBranch(branch string) []*Student {
	var results []*Student

	for _, student := range db.Students {
		if strings.EqualFold(student.Branch, branch) {
			results = append(results, student)
		}
	}
	return results
}
