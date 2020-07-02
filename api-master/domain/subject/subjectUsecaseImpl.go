package subject

type subjectUsecase struct {
	repo SubjectRepository
}

func NewSubjectUsecase(repository SubjectRepository) SubjectUsecase {
	return &subjectUsecase{
		repo: repository,
	}
}

func (s *subjectUsecase) GetSubjects() ([]*Subject, error) {
	subjectList, err := s.repo.GetSubjects()
	if err != nil {
		return nil, err
	}
	return subjectList, nil
}

func (s *subjectUsecase) GetSubjectByID(id string) (*Subject, error) {
	subject, err := s.repo.GetSubjectByID(id)
	if err != nil {
		return nil, err
	}
	return subject, nil
}

func (s *subjectUsecase) PostSubject(subject *Subject) error {
	err := s.repo.PostSubject(subject)
	if err != nil {
		return err
	}
	return nil
}

func (s *subjectUsecase) UpdateSubject(subject *Subject) error {
	err := s.repo.UpdateSubject(subject)
	if err != nil {
		return err
	}
	return nil
}

func (s *subjectUsecase) DeleteSubject(id string) error {
	err := s.repo.DeleteSubject(id)
	if err != nil {
		return err
	}
	return nil
}
