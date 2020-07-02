package teacher

type teacherUsecase struct {
	repo TeacherRepository
}

func NewTeacherUsecase(repository TeacherRepository) TeacherUsecase {
	return &teacherUsecase{repository}
}

func (t *teacherUsecase) GetTeachers() ([]*Teacher, error) {
	teacherList, err := t.repo.GetTeachers()
	if err != nil {
		return nil, err
	}
	return teacherList, nil
}

func (t *teacherUsecase) GetTeacherByID(id string) (*Teacher, error) {
	teacher, err := t.repo.GetTeacherByID(id)
	if err != nil {
		return nil, err
	}
	return teacher, nil
}

func (t *teacherUsecase) PostTeacher(teacher *Teacher) error {
	err := t.repo.PostTeacher(teacher)
	if err != nil {
		return err
	}
	return nil
}

func (t *teacherUsecase) UpdateTeacher(teacher *Teacher) error {
	err := t.repo.UpdateTeacher(teacher)
	if err != nil {
		return err
	}
	return nil
}

func (t *teacherUsecase) DeleteTeacher(id string) error {
	err := t.repo.DeleteTeacher(id)
	if err != nil {
		return err
	}
	return nil
}
