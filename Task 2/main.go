package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Student struct {
	Name   string
	Grades []float64
}

func (s *Student) averageGrade() float64 {
	var sum float64
	for _, grade := range s.Grades {
		sum += grade
	}
	return sum / float64(len(s.Grades))
}

func main() {
	var student Student
	var numGrades int

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введіть ім'я та прізвище студента: ")
	name, _ := reader.ReadString('\n')
	student.Name = strings.TrimSpace(name)

	fmt.Print("Введіть кількість оцінок: ")
	fmt.Scan(&numGrades)

	student.Grades = make([]float64, numGrades)
	for i := 0; i < numGrades; i++ {
		fmt.Printf("Введіть оцінку #%d: ", i+1)
		fmt.Scan(&student.Grades[i])
	}

	fmt.Printf("Середній бал студента %s: %.2f\n", student.Name, student.averageGrade())
}
