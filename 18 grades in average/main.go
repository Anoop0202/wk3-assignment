package main

import "fmt"

func main() {
	var writtenTest, labExams, assignments float64
	fmt.Println("enter the score for the weitten test, lab exams, and assignments(out of 100)")
	fmt.Scanln(&writtenTest, &labExams, &assignments)
	overallGrade := (writtenTest * 0.7) + (labExams * 0.2) + (assignments * 0.1)
	fmt.Printf("the overall grade is %.2f\n", overallGrade)
}
