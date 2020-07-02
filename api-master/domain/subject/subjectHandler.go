/**
* Created by Maulana Ibrahim
* on 27 June 2020
 */

package subject

import (
	"errors"
	"github.com/gorilla/mux"
	"maulibra/enigma_school/api-master/utils"
	"net/http"
)

type SubjectHandler struct {
	usecase SubjectUsecase
}

func NewSubjectHandler(subjectUsecase SubjectUsecase) *SubjectHandler {
	return &SubjectHandler{subjectUsecase}
}

func (ms *SubjectHandler) Subject(r *mux.Router) {
	r.HandleFunc("/subjects", ms.readSubjects).Methods(http.MethodGet)
	r.HandleFunc("/subject/{id}", ms.readSubjectById).Methods(http.MethodGet)
	r.HandleFunc("/subject", ms.addSubject).Methods(http.MethodPost)
	r.HandleFunc("/subject/{id}", ms.editSubject).Methods(http.MethodPut)
	r.HandleFunc("/subject/{id}", ms.deleteSubject).Methods(http.MethodDelete)

}

func (sh *SubjectHandler) readSubjects(res http.ResponseWriter, req *http.Request) {
	listSubjects, err := sh.usecase.GetSubjects()
	if err != nil {
		utils.HandleRequest(res, http.StatusBadGateway)
	}
	utils.HandleResponse(res, http.StatusOK, listSubjects)
}

func (sh *SubjectHandler) readSubjectById(res http.ResponseWriter, req *http.Request) {
	id := utils.DecodePathVariabel("id", req)
	if len(id) > 0 {
		subject, err := sh.usecase.GetSubjectByID(id)
		if (Subject{}) == *subject && err == nil {
			utils.HandleResponse(res, http.StatusNoContent, utils.STATUS_NO_CONTENT)
		} else {
			utils.HandleResponse(res, http.StatusOK, subject)
		}
	} else {
		utils.HandleRequest(res, http.StatusBadRequest)
	}
}

func (sh *SubjectHandler) addSubject(res http.ResponseWriter, req *http.Request) {
	var subject Subject
	err := utils.JsonDecoder(&subject, req)
	if err != nil {
		utils.HandleRequest(res, http.StatusBadRequest)
	} else {
		err = sh.usecase.PostSubject(&subject)
		if err != nil {
			utils.HandleRequest(res, http.StatusBadGateway)
		} else {
			utils.HandleResponse(res, http.StatusCreated,subject)
		}
	}
}

func (sh *SubjectHandler) editSubject(res http.ResponseWriter, req *http.Request) {
	var subject Subject
	var errGetId error

	id := utils.DecodePathVariabel("id", req)
	if len(id) > 0 {
		subject.Id = id
	} else {
		errGetId = errors.New("Not found params")
	}

	err := utils.JsonDecoder(&subject, req)
	if err == nil && errGetId == nil {
		getSubject, _ := sh.usecase.GetSubjectByID(subject.Id)
		if getSubject.Id == subject.Id && err == nil {
			err := sh.usecase.UpdateSubject(&subject)
			if err != nil {
				utils.HandleRequest(res, http.StatusBadGateway)
			} else {
				utils.HandleResponse(res, http.StatusOK,subject)
			}
		} else {
			utils.HandleRequest(res, http.StatusNotFound)
		}
	} else {
		utils.HandleRequest(res, http.StatusBadRequest)
	}
}

func (sh *SubjectHandler) deleteSubject(res http.ResponseWriter, req *http.Request) {
	id := utils.DecodePathVariabel("id", req)
	if len(id) > 0 {
		err := sh.usecase.DeleteSubject(id)
		if err != nil {
			utils.HandleRequest(res, http.StatusNotFound)
		} else {
			utils.HandleRequest(res, http.StatusOK)
		}
	} else {
		utils.HandleRequest(res, http.StatusBadRequest)
	}
}
