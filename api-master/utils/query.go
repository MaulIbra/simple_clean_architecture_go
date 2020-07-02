/**
* Created by Maulana Ibrahim
* on 27 June 2020
 */

package utils

const (
	SELECT_STUDENTS      = `SELECT * FROM STUDENT`
	SELECT_STUDENT_BY_ID = `SELECT * FROM STUDENT where studentId=?`
	INSERT_STUDENT       = `INSERT INTO STUDENT VALUES(?,?,?,?) `
	UPDATE_STUDENT       = `UPDATE STUDENT set first_name=?,last_name=?,email=? where studentId = ?`
	DELETE_STUDENT       = `DELETE FROM student WHERE studentId=?`

	SELECT_SUBJECTS      = `SELECT * FROM SUBJECT`
	SELECT_SUBJECT_BY_ID = `SELECT * FROM SUBJECT where subjectId=?`
	INSERT_SUBJECT       = `INSERT INTO SUBJECT VALUES(?,?) `
	UPDATE_SUBJECT       = `UPDATE SUBJECT set subject_name=? where subjectId = ?`
	DELETE_SUBJECT       = `DELETE FROM subject WHERE subjectId=?`

	SELECT_TEACHERS       = `SELECT * FROM TEACHER`
	SELECT_TEACHERS_BY_ID = `SELECT * FROM TEACHER where teacherId=?`
	INSERT_TEACHER        = `INSERT INTO TEACHER VALUES(?,?,?,?) `
	UPDATE_TEACHER        = `UPDATE teacher set first_name=?,last_name=?,email=? where teacherId = ?`
	DELETE_TEACHER        = `DELETE FROM teacher WHERE teacherId=?`
)
