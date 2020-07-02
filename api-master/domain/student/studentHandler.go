/**
* Created by Maulana Ibrahim
* on 27 June 2020
 */

package student

import (
	"errors"
	"github.com/gorilla/mux"
	"maulibra/enigma_school/api-master/utils"
	"net/http"
)

type StudentHandler struct {
	usecase StudentUsecase
}

func NewStudentHandler(studentUsecase StudentUsecase) *StudentHandler {
	return &StudentHandler{usecase: studentUsecase}
}

func (ms *StudentHandler) Students(r *mux.Router) {
	r.HandleFunc("/students", ms.readStudents).Methods(http.MethodGet)
	r.HandleFunc("/student/{id}", ms.readStudentById).Methods(http.MethodGet)
	r.HandleFunc("/student", ms.addStudent).Methods(http.MethodPost)
	r.HandleFunc("/student/{id}", ms.editStudent).Methods(http.MethodPut)
	r.HandleFunc("/student/{id}", ms.deleteStudent).Methods(http.MethodDelete)
}

func (sh *StudentHandler) readStudents(res http.ResponseWriter, req *http.Request) {
	listStudents, err := sh.usecase.GetStudents()
	if err != nil {
		utils.HandleRequest(res, http.StatusBadGateway)
	}
	utils.HandleResponse(res, http.StatusOK, listStudents)
}

func (sh *StudentHandler) readStudentById(res http.ResponseWriter, req *http.Request) {
	id := utils.DecodePathVariabel("id", req)
	if len(id) > 0 {
		student, err := sh.usecase.GetStudentByID(id)
		if (Student{}) == *student && err == nil {
			utils.HandleResponse(res, http.StatusOK, utils.STATUS_NO_CONTENT)
		} else {
			utils.HandleResponse(res, http.StatusOK, student)
		}
	} else {
		utils.HandleRequest(res, http.StatusBadRequest)
	}
}

func (sh *StudentHandler) addStudent(res http.ResponseWriter, req *http.Request) {
	var student Student
	err := utils.JsonDecoder(&student, req)
	if err != nil {
		utils.HandleRequest(res, http.StatusBadRequest)
	} else {
		err = sh.usecase.PostStudent(&student)
		if err != nil {
			utils.HandleRequest(res, http.StatusBadGateway)
		} else {
			utils.HandleResponse(res, http.StatusCreated,student)
		}
	}
}

func (sh *StudentHandler) editStudent(res http.ResponseWriter, req *http.Request) {
	var student Student
	var errGetId error

	id := utils.DecodePathVariabel("id", req)
	if len(id) > 0 {
		student.Id = id
	} else {
		errGetId = errors.New("Not found params")
	}

	err := utils.JsonDecoder(&student, req)
	if err == nil && errGetId == nil {
		getStudent, _ := sh.usecase.GetStudentByID(student.Id)
		if getStudent.Id == student.Id && err == nil {
			err := sh.usecase.UpdateStudent(&student)
			if err != nil {
				utils.HandleRequest(res, http.StatusBadGateway)
			} else {
				utils.HandleResponse(res, http.StatusOK,student)
			}
		} else {
			utils.HandleRequest(res, http.StatusNotFound)
		}
	} else {
		utils.HandleRequest(res, http.StatusBadRequest)
	}
}

func (sh *StudentHandler) deleteStudent(res http.ResponseWriter, req *http.Request) {
	id := utils.DecodePathVariabel("id", req)
	if len(id) != 0 {
		err := sh.usecase.DeleteStudent(id)
		if err != nil {
			utils.HandleRequest(res, http.StatusNotFound)
		} else {
			utils.HandleRequest(res, http.StatusOK)
		}
	} else {
		utils.HandleRequest(res, http.StatusBadRequest)
	}
}
