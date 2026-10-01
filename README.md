# University CLI

A command-line university management system written in Go.

This is a learning project I built while learning Go and exploring practical software development concepts such as structs, methods, interfaces, packages, error handling, JSON persistence, and testing.

## Features

- Add, list, search, update, and delete students
- Manage student marks and grades
- Add and manage teachers
- Create courses and assign teachers
- Enroll students in courses
- View course enrollment
- Calculate student statistics
- Persist data locally using JSON
- Validate user input and domain data
- Use typed and sentinel errors with idiomatic Go error handling
- Unit tests for core functionality

## Project Structure

```text
UniversityCLI/
main.go
go.mod

university/
course.go
test.go
database.go
errors.go
search.go
stats.go
storage.go
student.go
student_test.go
teacher.go
university.json
```

### Package responsibilities

- `main.go` — CLI interface and user interaction
- `student.go` — student model, validation, grades, and marks
- `teacher.go` — teacher model and validation
- `course.go` — course management and enrollment
- `database.go` — in-memory university data management
- `search.go` — student search operations
- `stats.go` — student statistics and ranking
- `storage.go` — JSON serialization and persistence
- `errors.go` — application-specific errors
- `*_test.go` — unit tests for the core package

## Requirements

- Go 1.27 or later

## Run the Project

Clone the repository:

```bash
git clone https://github.com/Aarav-Guleria/university-cli.git
cd university-cli
```

Run the application:

```bash
go run .
```

The application stores its data in:

```text
university.json
```

If the database file does not exist, the application starts with a new database.

## Run Tests

Run all tests with:

```bash
go test ./...
```

You can also run the tests with verbose output:

```bash
go test -v ./...
```

## Error Handling

The project uses Go's standard error-handling patterns where they are useful.

Examples include:

- `ValidationError` for structured validation failures
- Sentinel errors for known domain conditions
- `errors.As` for extracting typed errors
- `errors.Is` for checking known errors
- `fmt.Errorf(... %w ...)` for preserving underlying errors while adding context

The goal is to practice idiomatic Go error handling rather than adding abstractions that are unnecessary for the project.

## Persistence

The application uses JSON for local persistence.

The in memory database contains students, teachers, and courses. When saving, the application converts the data into JSON friendly representations and reconstructs the relationships when loading.

No external database is required.

## Learning Goals

This project is primarily a practical Go learning project.

Some of the concepts explored so far include:

- Structs and methods
- Pointers
- Interfaces
- Packages
- Maps and slices
- JSON encoding and decoding
- File I/O
- Error handling
- Custom error types
- Sentinel errors
- Error wrapping
- `errors.Is`
- `errors.As`
- Unit testing
- Basic project organization

## Status

This project is actively being used to learn Go through implementation and incremental refactoring.

The code is intentionally simple and is not intended to represent a production-ready university management system.

## Future Ideas

Possible future improvements include:

- Better CLI argument handling
- More comprehensive test coverage
- Improved input validation
- Cleaner CLI/domain separation
- More robust persistence handling
- Additional course and teacher operations
- Improved documentation and examples
- Automated checks with GitHub Actions

## License

This project is currently a personal learning project.
