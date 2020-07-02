package subject

import (
	"database/sql"
	"errors"
	guuid "github.com/google/uuid"
	"maulibra/enigma_school/api-master/utils"
)

type subjectDB struct {
	db *sql.DB
}

func NewSubjectRepository(db *sql.DB) SubjectRepository {
	return &subjectDB{db}
}

func (s *subjectDB) GetSubjects() ([]*Subject, error) {
	subjectList := []*Subject{}
	stmt, err := s.db.Prepare(utils.SELECT_SUBJECTS)
	if err != nil {
		return subjectList, err
	}
	rows, err := stmt.Query()
	if err != nil {
		return subjectList, err
	}

	defer stmt.Close()

	for rows.Next() {
		s := Subject{}
		err := rows.Scan(&s.Id, &s.SubjectName)
		if err != nil {
			return subjectList, err
		}
		subjectList = append(subjectList, &s)
	}
	return subjectList, nil
}

func (s *subjectDB) GetSubjectByID(id string) (*Subject, error) {
	subject := Subject{}
	stmt, err := s.db.Prepare(utils.SELECT_SUBJECT_BY_ID)
	if err != nil {
		return &subject, err
	}
	err = stmt.QueryRow(id).Scan(&subject.Id, &subject.SubjectName)
	if err != nil {
		return &subject, err
	}

	defer stmt.Close()

	return &subject, nil
}

func (s *subjectDB) PostSubject(subject *Subject) error {
	id := guuid.New()
	subject.Id = id.String()
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(utils.INSERT_SUBJECT)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = stmt.Exec(id, subject.SubjectName)
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	return tx.Commit()
}

func (s *subjectDB) UpdateSubject(subject *Subject) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(utils.UPDATE_SUBJECT)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = stmt.Exec(subject.SubjectName, subject.Id)

	defer stmt.Close()

	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (s *subjectDB) DeleteSubject(id string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(utils.DELETE_SUBJECT)

	if err != nil {
		tx.Rollback()
		return err
	}
	res, err := stmt.Exec(id)
	if err != nil {
		tx.Rollback()
		return err
	}

	defer stmt.Close()

	count, err := res.RowsAffected()
	if count == 0 {
		return errors.New("gagal delete")
	}
	return tx.Commit()
}
