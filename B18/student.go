package main

var students = []*Student{}

//Student is
type Student struct {
	Id    string
	Name  string
	Grade int32
}

//GetStudents is
func GetStudents() []*Student {
	return students
}

//SelectStudent is
func SelectStudent(id string) *Student {
	for _, each := range students {
		if each.Id == id {
			return each
		}
	}

	return nil
}

func init() {
	students = append(students, &Student{Id: "s001", Name: "bourne", Grade: 2})
	students = append(students, &Student{Id: "s002", Name: "ethan", Grade: 2})
	students = append(students, &Student{Id: "s003", Name: "wick", Grade: 3})
}
