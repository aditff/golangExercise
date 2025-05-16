package main

import (
	"fmt"
	"strings"
)

type Student struct {
	No         int
	Name       string
	Exercise   float64
	Homework   float64
	Exam       float64
	FinalScore float64
	Grade      string
}

func nfinalScore(exercise, homework, exam float64) float64 {
	return (exercise * 0.2) + (homework * 0.3) + (exam * 0.5)
}

func checkGrade(finalScore float64) string {
	switch {
	case finalScore >= 85 && finalScore <= 100:
		return "A"
	case finalScore >= 75 && finalScore <= 84:
		return "B"
	case finalScore >= 65 && finalScore <= 74:
		return "C"
	case finalScore >= 50 && finalScore <= 64:
		return "D"
	default:
		return "E"
	}
}

func main() {
	students := []Student{
		{1, "Andi", 85, 90, 80, 0, ""},
		{2, "Rara", 90, 90, 92, 0, ""},
		{3, "Okta", 70, 80, 70, 0, ""},
	}

	for i := range students {
		students[i].FinalScore = nfinalScore(students[i].Exercise, students[i].Homework, students[i].Exam)
		students[i].Grade = checkGrade(students[i].FinalScore)
	}

	fmt.Println("Report: Matematika\n")
	

	header := []string{"Name", "Exercise", "Homework", "Exam", "Final Score", "Grade"}
	headerFormat := "%-15s | %-8s | %-8s | %-5s | %-10s | %-5s\n"
	line := strings.Repeat("-", 71)

	fmt.Printf(headerFormat, header[0], header[1], header[2], header[3], header[4], header[5])
	fmt.Println(line)

	for _, student := range students {
		data := []interface{}{student.Name, student.Exercise, student.Homework, student.Exam, student.FinalScore, student.Grade}
		dataFormat := "%-15s | %-8.0f | %-8.0f | %-5.0f | %-10.1f | %-5s\n"
		fmt.Printf(dataFormat, data...)
	}
}