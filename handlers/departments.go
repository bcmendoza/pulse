package handlers

import (
	"errors"
	"fmt"
	"net/http"
)

func (hs *handlersState) addDepartment() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if logger, ok := validateMethod("/departments", r.Method, "POST", hs.logger, w); ok {
			req, ok := validateRequestFields(r.Body, logger, w)
			if ok {
				if req.Department == "" {
					logger.Error().AnErr("addDepartment()", errors.New("empty field(s)")).Msg("400 Bad Request")
					Report(ProblemDetail{
						StatusCode: http.StatusBadRequest,
						Detail:     "Department name is empty",
					}, w)
					return
				}

				if _, ok := hs.hospital.Children[req.Department]; ok {
					logger.Error().AnErr("addDepartment()", errors.New("empty field(s)")).Msg("400 Bad Request")
					Report(ProblemDetail{
						StatusCode: http.StatusBadRequest,
						Detail:     "Department already exists",
					}, w)
					return
				}
				hs.hospital.AddDepartment(req.Department)
				w.WriteHeader(http.StatusOK)
				w.Header().Set("Content-Type", "application/json")
				jsonResp := fmt.Sprintf("{\"added\": \"%s\"}", req.Department)
				if _, err := w.Write([]byte(jsonResp)); err != nil {
					logger.Error().AnErr("w.Write", err).Msg("500 Internal server error")
				} else {
					logger.Info().Msg("200 OK")
				}
			}
		}
	}
}
