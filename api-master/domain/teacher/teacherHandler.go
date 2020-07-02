/**
* Created by Maulana Ibrahim
* on 27 June 2020
 */

package teacher

import (
	"errors"
	"github.com/gorilla/mux"
	"maulibra/enigma_school/api-master/utils"
	"net/http"
)

type teacherHandler struct {
	usecase TeacherUsecase
}

func NewTeacherHandler(teacherUsecase TeacherUsecase) *teacherHandler {
	return &teacherHandler{teacherUsecase}
}
func (ms *teacherHandler) Teacher(r *mux.Router) {
	r.HandleFunc("/teachers", ms.readTeachers).Methods(http.MethodGet)
	r.HandleFunc("/teacher/{id}", ms.readTeacherById).Methods(http.MethodGet)
	r.HandleFunc("/teacher", ms.addTeacher).Methods(http.MethodPost)
	r.HandleFunc("/teacher/{id}", ms.editTeacher).Methods(http.MethodPut)
	r.HandleFunc("/teacher/{id}", ms.deleteTeacher).Methods(http.MethodDelete)
}

func (ts *teacherHandler) readTeachers(res http.ResponseWriter, req *http.Request) {
	listTeachers, err := ts.usecase.GetTeachers()
	if err != nil {
		utils.HandleRequest(res, http.StatusBadGateway)
	}
	utils.HandleResponse(res, http.StatusOK, listTeachers)
}

func (ts *teacherHandler) readTeacherById(res http.ResponseWriter, req *http.Request) {
	id := utils.DecodePathVariabel("id", req)
	if len(id) > 0 {
		teacher, err := ts.usecase.GetTeacherByID(id)
		if (Teacher{}) == *teacher && err == nil {
			utils.HandleResponse(res, http.StatusNoContent, utils.STATUS_NO_CONTENT)
		} else {
			utils.HandleResponse(res, http.StatusOK, teacher)
		}
	} else {
		utils.HandleRequest(res, http.StatusBadRequest)
	}
}

func (ts *teacherHandler) addTeacher(res http.ResponseWriter, req *http.Request) {
	var teacher Teacher
	err := utils.JsonDecoder(&teacher, req)
	if err != nil {
		utils.HandleRequest(res, http.StatusBadRequest)
	} else {
		err = ts.usecase.PostTeacher(&teacher)
		if err != nil {
			utils.HandleRequest(res,http.StatusBadGateway)
		} else {
			utils.HandleResponse(res,http.StatusCreated,teacher)
		}
	}
}

func (ts *teacherHandler) editTeacher(res http.ResponseWriter, req *http.Request) {
	var teacher Teacher
	var errGetId error

	id := utils.DecodePathVariabel("id", req)
	if len(id) > 0 {
		teacher.Id = id
	} else {
		errGetId = errors.New("Not found params")
	}

	err := utils.JsonDecoder(&teacher, req)
	if err == nil && errGetId == nil {
		getTeacher, _ := ts.usecase.GetTeacherByID(teacher.Id)
		if getTeacher.Id == teacher.Id && err == nil {
			err := ts.usecase.UpdateTeacher(&teacher)
			if err != nil {
				utils.HandleRequest(res, http.StatusBadGateway)
			} else {
				utils.HandleResponse(res, http.StatusOK,teacher)
			}
		} else {
			utils.HandleRequest(res, http.StatusNotFound)
		}
	} else {
		utils.HandleRequest(res, http.StatusBadRequest)
	}
}

func (ts *teacherHandler) deleteTeacher(res http.ResponseWriter, req *http.Request) {
	id := utils.DecodePathVariabel("id", req)
	if len(id) > 0 {
		err := ts.usecase.DeleteTeacher(id)
		if err != nil {
			utils.HandleRequest(res, http.StatusNotFound)
		} else {
			utils.HandleRequest(res, http.StatusOK)
		}
	} else {
		utils.HandleRequest(res, http.StatusBadRequest)
	}
}
