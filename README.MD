# ESEP Grade Calculator

This is a simple Go library for calculating a student's final letter grade based on weighted averages of assignments, exams, and essays. The project is written in Go (Golang) and includes unit tests to verify functionality.

## Overview

The `GradeCalculator` library allows users to:
- Add grades for assignments, exams, and essays.
- Calculate a weighted numerical grade (50% assignments, 35% exams, 15% essays).
- Convert the numerical grade to a letter grade (A, B, C, D, or F).

## Tech Stack
- **Golang**: Version 1.24 or higher

## Installation

1. Ensure you have Go 1.24+ installed. You can check your Go version with:
   ```
   go version
   ```
   If you do not have Go installed you can do so here: https://go.dev/dl/
2. Clone or fork the repository:
   ```
   git clone <repository-url>
   ```
3. Navigate to the project directory:
   ```
   cd <directory_name>
   ```

## Running Tests

The repository includes unit tests in `grade_calculator_test.go`. To run them:
```
cd grade_calculator
go test ./...
```

To generate a code coverage report:
```
go test -coverprofile=coverage.txt ./...
```

## Project Structure

- **`grade_calculator.go`**: Core logic for the grade calculator, including structs and methods for managing and calculating grades.
- **`grade_calculator_test.go`**: Unit tests to verify the `GetFinalGrade` method.
- **`go.mod`**: Go module definition for dependency management.
- **`README.md`**: This file, providing project documentation.

## Resources
- [Official Go Documentation](https://go.dev/doc/)
- [Go by Example](https://gobyexample.com/)
- [Go Modules Reference](https://go.dev/ref/mod)