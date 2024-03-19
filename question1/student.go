package question1

import (
	"fmt"
	"strings"
)

type StudentService interface {
	Add(name string, age int)
	DelByName(name string)
	GetByAgeRange(minAge, maxAge int) []*Student
}

type Student struct {
	Name string
	Age  int
}

type StudentManager struct {
	StudentArray []*Student
}

func CreateStudentService(currentStudents []*Student) StudentService {
	return &StudentManager{
		StudentArray: currentStudents,
	}
}

func (s *StudentManager) Add(name string, age int) {
	student := Student{Name: name, Age: age}
	s.StudentArray = append(s.StudentArray, &student)
	fmt.Println(fmt.Sprintf("Added student %+v", student))
}

func (s *StudentManager) DelByName(name string) {
	if len(s.StudentArray) > 0 {
		index := 0
		found := false
		for i, student := range s.StudentArray {
			if strings.EqualFold(student.Name, name) {
				index = i
				found = true
				break
			}
		}
		if found {
			s.StudentArray = append(s.StudentArray[:index], s.StudentArray[index+1:]...)
			fmt.Println(fmt.Sprintf("Deleted student %s", name))
			return
		}
	}
	fmt.Println(fmt.Sprintf("No student %s", name))
}

func (s *StudentManager) GetByAgeRange(minAge, maxAge int) []*Student {
	// database call
	studentArray := []*Student{}
	for _, student := range s.StudentArray {
		if student.Age >= minAge && student.Age < maxAge {
			studentArray = append(studentArray, student)
		}
	}
	return studentArray
}
