package main

import (
	"fmt"

	"github.com/PipesNBottles/hitachi/question1"
	"github.com/PipesNBottles/hitachi/question2"
)

func q2() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	nums = question2.AddValsAndRemoveVals(nums)
	fmt.Println(nums)
}

func q1() {
	students := []*question1.Student{{Name: "s1", Age: 11}, {Name: "s2", Age: 12}, {Name: "s3", Age: 12}, {Name: "s4", Age: 12}, {Name: "s5", Age: 13}}
	studentService := question1.CreateStudentService(students)
	studentService.Add("s6", 14)
	studentService.DelByName("s2")
	students = studentService.GetByAgeRange(10, 14)
	for _, student := range students {
		fmt.Printf(fmt.Sprintf("%+v ", student))
	}
}

func main() {
	q1()
	q2()
}
