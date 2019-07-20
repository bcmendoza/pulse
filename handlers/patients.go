package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/bcmendoza/pulse/model"
	"github.com/bcmendoza/pulse/utils"
)

func (hs *handlersState) addPatient() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if logger, ok := validateMethod("/department", r.Method, "POST", hs.logger, w); ok {
			req, ok := validateRequestFields(r.Body, logger, w)
			if ok {
				if req.Department == "" {
					logger.Error().AnErr("addPatient()", errors.New("missing field(s)")).Msg("400 Bad Request")
					Report(ProblemDetail{
						StatusCode: http.StatusBadRequest,
						Detail:     "Department name to add patient to is empty",
					}, w)
					return
				}

				if _, ok := hs.hospital.Children[req.Department]; !ok {
					logger.Error().AnErr("addPatient()", errors.New("missing field(s)")).Msg("400 Bad Request")
					Report(ProblemDetail{
						StatusCode: http.StatusBadRequest,
						Detail:     "Department name to add patient to is invalid",
					}, w)
					return
				}

				var uuid string
				for uuid == "" {
					temp := utils.UUID()
					if _, ok := hs.hospital.PatientKeys[model.PatientKey{
						Department: req.Department,
						Patient:    temp,
					}]; !ok {
						uuid = temp
					}
				}
				hs.hospital.AddPatient(req.Department, utils.UUID())
				w.WriteHeader(http.StatusOK)
				w.Header().Set("Content-Type", "application/json")
				jsonResp := fmt.Sprintf("{\"created\": \"%s\"}", uuid)
				if _, err := w.Write([]byte(jsonResp)); err != nil {
					logger.Error().AnErr("w.Write", err).Msg("500 Internal server error")
				} else {
					logger.Info().Msg("200 OK")
				}
			}
		}
	}
}
