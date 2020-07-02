package teacher

type TeacherRepository interface {
	GetTeachers() ([]*Teacher, error)
	GetTeacherByID(id string) (*Teacher, error)
	PostTeacher(teacher *Teacher) error
	UpdateTeacher(teacher *Teacher) error
	DeleteTeacher(id string) error
}
