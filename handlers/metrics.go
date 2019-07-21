package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/bcmendoza/pulse/model"
)

func (hs *handlersState) addMetric() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if logger, ok := validateMethod("/department", r.Method, "POST", hs.logger, w); ok {
			req, ok := validateRequestFields(r.Body, logger, w)
			if ok {
				if req.Metric == "" || req.UnitType == "" {
					logger.Error().AnErr("addMetric()", errors.New("missing field(s)")).Msg("400 Bad Request")
					Report(ProblemDetail{
						StatusCode: http.StatusBadRequest,
						Detail:     "Metric, and/or unitType are empty",
					}, w)
				}

				if _, ok := hs.hospital.MetricKeys[model.MetricKey{
					Department: req.Department,
					Patient:    req.Patient,
					Metric:     req.Metric,
				}]; ok {
					logger.Error().AnErr("addMetric()", errors.New("missing field(s)")).Msg("400 Bad Request")
					Report(ProblemDetail{
						StatusCode: http.StatusBadRequest,
						Detail:     "Metric already exists",
					}, w)
				}

				success := false
				if req.Patient == "" && req.Department == "" {
					hs.hospital.AddHospitalMetrics(req.Metric, req.UnitType)
					success = true
				} else if req.Patient == "" && req.Department != "" {
					if _, ok := hs.hospital.Children[req.Department]; ok {
						hs.hospital.AddDepartmentMetrics(req.Department, req.Metric, req.UnitType)
						success = true
					}
				} else if req.Patient != "" && req.Department != "" {
					if _, ok := hs.hospital.PatientKeys[model.PatientKey{
						Department: req.Department,
						Patient:    req.Patient,
					}]; ok {
						hs.hospital.AddPatientMetric(req.Department, req.Patient, req.Metric, req.UnitType)
						success = true
					}
				}

				if success == false {
					logger.Error().AnErr("addMetric()", errors.New("invalid field(s)")).Msg("400 Bad Request")
					Report(ProblemDetail{
						StatusCode: http.StatusBadRequest,
						Detail:     "Unable to add",
					}, w)
					return
				}

				w.WriteHeader(http.StatusOK)
				w.Header().Set("Content-Type", "application/json")
				jsonResp := fmt.Sprintf("{\"created\": \"%s\"}", req.Metric)
				if _, err := w.Write([]byte(jsonResp)); err != nil {
					logger.Error().AnErr("w.Write", err).Msg("500 Internal server error")
				} else {
					logger.Info().Msg("200 OK")
				}
			}
		}
	}
}
