package student

import (
	"database/sql"
	"errors"
	guuid "github.com/google/uuid"
	"maulibra/enigma_school/api-master/utils"
)

type studentDB struct {
	db *sql.DB
}

func NewStudentRepository(db *sql.DB) StudentRepository {
	return &studentDB{db}
}

func (sDB *studentDB) GetStudents() ([]*Student, error) {
	listStudent := []*Student{}
	stmt, err := sDB.db.Prepare(utils.SELECT_STUDENTS)
	if err != nil {
		return listStudent, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return listStudent, err
	}

	for rows.Next() {
		s := Student{}
		err := rows.Scan(&s.Id, &s.FirstName, &s.LastName, &s.Email)
		if err != nil {
			return listStudent, err
		}
		listStudent = append(listStudent, &s)
	}
	return listStudent, nil
}

func (sDB *studentDB) GetStudentByID(id string) (*Student, error) {
	student := Student{}
	stmt, err := sDB.db.Prepare(utils.SELECT_STUDENT_BY_ID)
	if err != nil {
		return &student, err
	}
	defer stmt.Close()
	stmt.QueryRow(id).Scan(&student.Id, &student.FirstName, &student.LastName, &student.Email)
	if err != nil {
		return &student, err
	}
	return &student, nil
}

func (sDB *studentDB) PostStudent(student *Student) error {
	id := guuid.New()
	student.Id = id.String()
	tx, err := sDB.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(utils.INSERT_STUDENT)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = stmt.Exec(id, student.FirstName, student.LastName, student.Email)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (sDB *studentDB) UpdateStudent(student *Student) error {
	tx, err := sDB.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(utils.UPDATE_STUDENT)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = stmt.Exec(student.FirstName, student.LastName, student.Email, student.Id)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (sDB *studentDB) DeleteStudent(id string) error {
	tx, err := sDB.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(utils.DELETE_STUDENT)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}
	res, err := stmt.Exec(id)
	if err != nil {
		tx.Rollback()
		return err
	}
	count, err := res.RowsAffected()
	if count == 0 {
		return errors.New("gagal delete")
	}
	return tx.Commit()
}
