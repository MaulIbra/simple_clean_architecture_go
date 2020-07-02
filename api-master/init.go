package api_master

import (
	"database/sql"
	"github.com/gorilla/mux"
	"maulibra/enigma_school/api-master/domain/student"
	"maulibra/enigma_school/api-master/domain/subject"
	"maulibra/enigma_school/api-master/domain/teacher"
)

func Init(db *sql.DB, router *mux.Router) {

	//student
	studentRepository := student.NewStudentRepository(db)
	studentUsecase := student.NewStudentUsecase(studentRepository)
	studentHandler := student.NewStudentHandler(studentUsecase)
	studentHandler.Students(router)

	//teacher
	teacherRepository := teacher.NewTeacherRepository(db)
	teacherUsecase := teacher.NewTeacherUsecase(teacherRepository)
	teacherHandler := teacher.NewTeacherHandler(teacherUsecase)
	teacherHandler.Teacher(router)

	//subject
	subjectRepository := subject.NewSubjectRepository(db)
	subjectUsecase := subject.NewSubjectUsecase(subjectRepository)
	subjectHandler := subject.NewSubjectHandler(subjectUsecase)
	subjectHandler.Subject(router)

}
