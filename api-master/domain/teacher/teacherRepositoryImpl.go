package teacher

import (
	"database/sql"
	"errors"
	guuid "github.com/google/uuid"
	"maulibra/enigma_school/api-master/utils"
)

type teacherDB struct {
	db *sql.DB
}

func NewTeacherRepository(db *sql.DB) TeacherRepository {
	return &teacherDB{db: db}
}

func (t *teacherDB) GetTeachers() ([]*Teacher, error) {
	teacherList := []*Teacher{}
	stmt, err := t.db.Prepare(utils.SELECT_TEACHERS)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		t := Teacher{}
		err := rows.Scan(&t.Id, &t.FirstName, &t.LastName, &t.Email)
		if err != nil {
			return teacherList, err
		}
		teacherList = append(teacherList, &t)
	}
	return teacherList, nil
}

func (t *teacherDB) GetTeacherByID(id string) (*Teacher, error) {
	teacher := Teacher{}
	stmt, err := t.db.Prepare(utils.SELECT_TEACHERS_BY_ID)
	if err != nil {
		return &teacher, err
	}
	defer stmt.Close()
	stmt.QueryRow(id).Scan(&teacher.Id, &teacher.FirstName, &teacher.LastName, &teacher.Email)
	if err != nil {
		return &teacher, err
	}
	return &teacher, nil
}

func (t *teacherDB) PostTeacher(teacher *Teacher) error {
	id := guuid.New()
	teacher.Id = id.String()
	tx, err := t.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(utils.INSERT_TEACHER)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = stmt.Exec(id, teacher.FirstName, teacher.LastName, teacher.Email)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (t *teacherDB) UpdateTeacher(teacher *Teacher) error {
	tx, err := t.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(utils.UPDATE_TEACHER)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = stmt.Exec(teacher.FirstName, teacher.LastName, teacher.Email, teacher.Id)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (t *teacherDB) DeleteTeacher(id string) error {
	tx, err := t.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(utils.DELETE_TEACHER)
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
