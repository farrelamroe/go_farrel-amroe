package main

import "fmt"

type Student struct {
	Name         string
	MathScore    int
	ScienceScore int
	EnglishScore int
}

func getInfo(students []Student) {

	bestMathScore := -1
	bestScienceScore := -1
	bestEnglishScore := -1
	bestAverage := -1.0
	bestStudentMath := ""
	bestStudentScience := ""
	bestStudentEnglish := ""
	bestStudentAverage := ""

	for _, student := range students {

		if student.MathScore > bestMathScore {
			bestMathScore = student.MathScore
			bestStudentMath = student.Name
		}

		if student.ScienceScore > bestScienceScore {
			bestScienceScore = student.ScienceScore
			bestStudentScience = student.Name
		}

		if student.EnglishScore > bestEnglishScore {
			bestEnglishScore = student.EnglishScore
			bestStudentEnglish = student.Name
		}

		average := float64(student.MathScore+student.ScienceScore+student.EnglishScore) / 3.0

		if average > bestAverage {
			bestAverage = average
			bestStudentAverage = student.Name
		}
	}

	fmt.Printf("Best student in math: %s with %d\n", bestStudentMath, bestMathScore)
	fmt.Printf("Best student in science: %s with %d\n", bestStudentScience, bestScienceScore)
	fmt.Printf("Best student in english: %s with %d\n", bestStudentEnglish, bestEnglishScore)
	fmt.Printf("Best student in average: %s with %.2f\n", bestStudentAverage, bestAverage)
}

func main() {
	students := []Student{
		{"jamie", 67, 88, 90},
		{"michael", 77, 77, 90},
		{"aziz", 100, 65, 88},
		{"jamal", 50, 80, 75},
		{"eric", 70, 80, 65},
		{"john", 80, 76, 68},
	}

	getInfo(students)
}
