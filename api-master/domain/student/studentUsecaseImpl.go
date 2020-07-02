package student

type studentUsecase struct {
	repository StudentRepository
}

func NewStudentUsecase(repository StudentRepository) StudentUsecase {
	return &studentUsecase{repository: repository}
}

func (s *studentUsecase) GetStudents() ([]*Student, error) {
	studentList, err := s.repository.GetStudents()
	if err != nil {
		return nil, err
	}
	return studentList, nil
}

func (s *studentUsecase) GetStudentByID(id string) (*Student, error) {
	student, err := s.repository.GetStudentByID(id)
	if err != nil {
		return nil, err
	}
	return student, nil
}

func (s *studentUsecase) PostStudent(student *Student) error {
	err := s.repository.PostStudent(student)
	if err != nil {
		return err
	}
	return nil
}

func (s *studentUsecase) UpdateStudent(student *Student) error {
	err := s.repository.UpdateStudent(student)
	if err != nil {
		return err
	}
	return nil
}

func (s *studentUsecase) DeleteStudent(id string) error {
	err := s.repository.DeleteStudent(id)
	if err != nil {
		return err
	}
	return nil
}
