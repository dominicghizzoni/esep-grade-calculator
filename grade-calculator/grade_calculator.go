package esepunittests

type GradeCalculator struct {
	grades []Grade
	passFailMode bool
}

type GradeType int

const (
	Assignment GradeType = iota
	Exam
	Essay
)

var gradeTypeName = map[GradeType]string{
	Assignment: "assignment",
	Exam:       "exam",
	Essay:      "essay",
}

func (gt GradeType) String() string {
	return gradeTypeName[gt]
}

type Grade struct {
	Name  string
	Grade int
	Type  GradeType
}

func NewGradeCalculator() *GradeCalculator {
	return &GradeCalculator{
		grades: make([]Grade, 0),
		passFailMode: false,
	}
}

func NewGradeCalculatorWithPassFail(passFail bool) *GradeCalculator {
	return &GradeCalculator{
		grades: make([]Grade, 0),
		passFailMode: passFail,
	}
}

func (gc *GradeCalculator) GetFinalGrade() string {
	numericalGrade := gc.calculateNumericalGrade()

	if gc.passFailMode {
		if numericalGrade >= 70 {
			return "Pass"
		}
		return "Fail"
	}

	if numericalGrade >= 90 {
		return "A"
	} else if numericalGrade >= 80 {
		return "B"
	} else if numericalGrade >= 70 {
		return "C"
	} else if numericalGrade >= 60 {
		return "D"
	}
	return "F"
}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
	gc.grades = append(gc.grades, Grade{
		Name:  name,
		Grade: grade,
		Type:  gradeType,
	})
}

func (gc *GradeCalculator) calculateNumericalGrade() int {
	assignmentGrades := []Grade{}
	examGrades := []Grade{}
	essayGrades := []Grade{}

	for _, grade := range gc.grades {
		switch grade.Type {
		case Assignment:
			assignmentGrades = append(assignmentGrades, grade)
		case Exam:
			examGrades = append(examGrades, grade)
		case Essay:
			essayGrades = append(essayGrades, grade)
		}
	}

	assignment_average := computeAverage(assignmentGrades)
	exam_average := computeAverage(examGrades)
	essay_average := computeAverage(essayGrades)

	weighted_grade := float64(assignment_average)*0.5 + float64(exam_average)*0.35 + float64(essay_average)*0.15

	return int(weighted_grade)
}

func computeAverage(grades []Grade) int {
	if len(grades) == 0 {
		return 0
	}
	sum := 0
	for _, grade := range grades {
		sum += grade.Grade
	}
	return sum / len(grades)
}