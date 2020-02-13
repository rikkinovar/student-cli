package main

import (
	"fmt"
	"strconv"
)

type Student struct {
	StudentID     string
	Name          string
	FinalScore    float32
	Grade         string
	MidScore      int32
	SemesterScore int32
	Attendance    int32
}

func validateNumber(param string) (valid bool, value int32) {
	result, err := strconv.Atoi(param)
	if err != nil {
		valid = false
		value = 0
		return
	}
	valid = true
	value = int32(result)
	return
}

func main() {

	// reader := bufio.NewReader(os.Stdin)
	var studentCountParam string
	var studentCount int32
	fmt.Println("Input the number of students:")
	for {
		fmt.Scan(&studentCountParam)
		valid, value := validateNumber(studentCountParam)
		if !valid {
			fmt.Println("invalid number, please re-input: ")
		} else {
			studentCount = value
			break
		}
	}

	// for i:=0;i<studencou

	// studentCount, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	fmt.Println(studentCount)
	// fmt.Println("Student Id:")
	// fmt.Println("Name:")
	// fmt.Println("Mid Term Test Score:")
	// fmt.Println("Semester Test Score:")
	// fmt.Println("Attendance Score:")
}
