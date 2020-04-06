package main

var students = []*Student{}

//Student is
type Student struct {
	ID    string
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
		if each.ID == id {
			return each
		}
	}

	return nil
}

func init() {
	students = append(students, &Student{ID: "s001", Name: "bourne", Grade: 2})
	students = append(students, &Student{ID: "s002", Name: "ethan", Grade: 2})
	students = append(students, &Student{ID: "s003", Name: "wick", Grade: 3})
}
