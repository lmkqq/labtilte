package main

import (
	"testing"
)

func TestAverageGrade(t *testing.T) {
	// Тест 1: Студент з 3 оцінками
	student := Student{
		Name:   "Іван Іванов",
		Grades: []float64{5.0, 4.0, 3.0},
	}

	expected := 4.0
	result := student.averageGrade()
	if result != expected {
		t.Errorf("expected %.2f, got %.2f", expected, result)
	}

	// Тест 2: Студент з 2 оцінками
	student = Student{
		Name:   "Марія Петрівна",
		Grades: []float64{4.5, 5.5},
	}

	expected = 5.0
	result = student.averageGrade()
	if result != expected {
		t.Errorf("expected %.2f, got %.2f", expected, result)
	}

	// Тест 3: Студент з однією оцінкою
	student = Student{
		Name:   "Артем Сидоров",
		Grades: []float64{3.0},
	}

	expected = 3.0
	result = student.averageGrade()
	if result != expected {
		t.Errorf("expected %.2f, got %.2f", expected, result)
	}

	// Тест 4: Студент з 5 оцінками
	student = Student{
		Name:   "Ольга Вікторівна",
		Grades: []float64{4.0, 4.5, 5.0, 3.5, 4.0},
	}

	expected = 4.2
	result = student.averageGrade()
	if result != expected {
		t.Errorf("expected %.2f, got %.2f", expected, result)
	}
}
