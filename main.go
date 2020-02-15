package main

import (
	"fmt"
	"strconv"
)

//Student Struct
type Student struct {
	StudentID     string
	Name          string
	FinalScore    float32
	Grade         string
	MidScore      int
	SemesterScore int
	Attendance    int
}

func validateNumber(param string) (valid bool, value int) {
	result, err := strconv.Atoi(param)
	if err != nil {
		valid = false
		value = 0
		return
	}
	valid = true
	value = result
	return
}

func calculateFinalScore(midScore, semesterScore, attendanceScore int) float32 {
	return (0.2 * float32(attendanceScore)) + (0.4 * float32(midScore)) + (0.4 * float32(semesterScore))
}

func calculateGrade(finalScore float32) string {
	if finalScore >= 85 {
		return "A"
	} else if finalScore >= 76 {
		return "B"
	} else if finalScore >= 61 {
		return "C"
	} else if finalScore >= 46 {
		return "D"
	} else {
		return "E"
	}
}

func printHeader() {
	fmt.Println("\n\n\n==========================================================================")
	fmt.Print("No. \t")
	fmt.Print("Student ID. \t")
	fmt.Print("Name. \t\t")
	fmt.Print("Final Score. \t")
	fmt.Println("Grade")
	fmt.Println("==========================================================================")
}

func printBody(students []Student) {
	for i := 0; i < len(students); i++ {
		fmt.Printf("%d \t", i+1)
		fmt.Printf("%s \t\t", students[i].StudentID)
		fmt.Printf("%s \t\t", students[i].Name)
		fmt.Printf("%.0f \t\t", students[i].FinalScore)
		fmt.Printf("%s \t", students[i].Grade)
		fmt.Println()
	}
}

func printFooter(studentCount, studentPass, studentFail int) {
	fmt.Println("==========================================================================")
	fmt.Printf("Number of Students\t\t: %d\n", studentCount)
	fmt.Printf("Number of Passing Students\t: %d\n", studentPass)
	fmt.Printf("Number of Failed Students\t: %d\n", studentFail)
	fmt.Println("==========================================================================")
}

func scanNumber() int {
	var inputNumber string
	var result int
	for {
		fmt.Scan(&inputNumber)
		valid, value := validateNumber(inputNumber)
		if !valid {
			fmt.Println("invalid number, please re-input: ")
		} else {
			result = value
			break
		}
	}
	return result
}

func main() {

	var studentCount int
	fmt.Print("Input the number of students:")
	studentCount = scanNumber()

	var students = make([]Student, 0)
	var studentPass int = 0
	var studentFail int = 0

	for i := 0; i < studentCount; i++ {
		var (
			studentID       string
			name            string
			finalScore      float32
			grade           string
			midScore        int
			semesterScore   int
			attendanceScore int
		)

		fmt.Printf("\n\n=========Student %d=========\n", i+1)
		fmt.Print("Student Id: ")
		fmt.Scan(&studentID)
		fmt.Print("Name: ")
		fmt.Scan(&name)
		fmt.Print("Mid Term Test Score: ")
		midScore = scanNumber()
		fmt.Print("Semester Test Score: ")
		semesterScore = scanNumber()
		fmt.Print("Attendance: ")
		attendanceScore = scanNumber()

		finalScore = calculateFinalScore(midScore, semesterScore, attendanceScore)
		if finalScore < 61 {
			studentFail++
		} else {
			studentPass++
		}
		grade = calculateGrade(finalScore)
		students = append(students, Student{StudentID: studentID, Name: name, MidScore: midScore, SemesterScore: semesterScore, Attendance: attendanceScore, Grade: grade, FinalScore: finalScore})
	}

	printHeader()
	printBody(students)
	printFooter(studentCount, studentPass, studentFail)
}
