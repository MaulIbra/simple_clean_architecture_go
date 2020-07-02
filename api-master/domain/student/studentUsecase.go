package student

type StudentUsecase interface {
	GetStudents() ([]*Student, error)
	GetStudentByID(id string) (*Student, error)
	PostStudent(student *Student) error
	UpdateStudent(student *Student) error
	DeleteStudent(id string) error
}
