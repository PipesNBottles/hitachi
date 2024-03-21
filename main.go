package main

import (
	"fmt"

	"github.com/PipesNBottles/hitachi/question1"
	"github.com/PipesNBottles/hitachi/question2"
	"github.com/PipesNBottles/hitachi/question3"
	"github.com/PipesNBottles/hitachi/question4"
)

func q4() {
	lands := []string{"AEEA", "EAAAAAE", "EAEAE"}
	for _, community := range lands {
		fmt.Println(question4.FindMinGroceries(community))
	}
}

func q3() {
	phoneNumbers := []string{
		"+1 (613)-995-0253",
		"613-995-0253",
		"1 613 995 0253",
		"613.995.0253",
		"+1 6(13)-995-0253",
		"+1 61(3)-995-0253",
		"613.99.50253",
		"61-3-995-0253",
	}
	for _, phoneNumber := range phoneNumbers {
		result, err := question3.Number(phoneNumber)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}
	}
}

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
	q3()
	q4()
}
