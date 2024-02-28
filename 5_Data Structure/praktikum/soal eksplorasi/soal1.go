package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type CourseData struct {
	CourseName              string
	MentorSatisfactionScore int
	LearningSatisfactionScore int
}

func main() {
	file, err := os.Open("dataset.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var courses []CourseData

	totalMentorSatisfaction := 0
	totalLearningSatisfaction := 0

	for _, record := range records {
		mentorSatisfaction, err := strconv.Atoi(record[1])
		if err != nil {
			fmt.Println("Error parsing mentor satisfaction:", err)
			continue
		}
		learningSatisfaction, err := strconv.Atoi(record[2])
		if err != nil {
			fmt.Println("Error parsing learning satisfaction:", err)
			continue
		}

		totalMentorSatisfaction += mentorSatisfaction
		totalLearningSatisfaction += learningSatisfaction

		course := CourseData{
			CourseName:              record[0],
			MentorSatisfactionScore: mentorSatisfaction,
			LearningSatisfactionScore: learningSatisfaction,
		}
		courses = append(courses, course)
	}

	averageMentorSatisfaction := float64(totalMentorSatisfaction) / float64(len(records))
	averageLearningSatisfaction := float64(totalLearningSatisfaction) / float64(len(records))

	fmt.Printf("Rata-rata kepuasan mentor: %.2f\n", averageMentorSatisfaction)
	fmt.Printf("Rata-rata kepuasan belajar: %.2f\n", averageLearningSatisfaction)

	maxMentorSatisfaction := 0
	var topCourse string
	for _, course := range courses {
		if course.MentorSatisfactionScore > maxMentorSatisfaction {
			maxMentorSatisfaction = course.MentorSatisfactionScore
			topCourse = course.CourseName
		}
	}
	fmt.Printf("Kursus dengan rata-rata kepuasan mentor tertinggi: %s\n", topCourse)
}
