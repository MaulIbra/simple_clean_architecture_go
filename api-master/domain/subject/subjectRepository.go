package subject

type SubjectRepository interface {
	GetSubjects() ([]*Subject, error)
	GetSubjectByID(id string) (*Subject, error)
	PostSubject(subject *Subject) error
	UpdateSubject(subject *Subject) error
	DeleteSubject(id string) error
}
