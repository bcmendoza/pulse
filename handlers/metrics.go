package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/bcmendoza/pulse/model"
)

type MetricsReqBody struct {
	Department string  `json:"department"`
	Patient    string  `json:"patient"`
	Metric     string  `json:"metric"`
	UnitType   string  `json:"unitType"`
	Lower      float64 `json:"lower"`
	Upper      float64 `json:"upper"`
}

func (hs *handlersState) addMetric() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if logger, ok := validateMethod("/metrics", r.Method, "POST", hs.logger, w); ok {
			var req MetricsReqBody
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&req)
			if err != nil {
				logger.Error().AnErr("json.NewDecoder", err).Msg("400 Bad Request")
				Report(ProblemDetail{
					StatusCode: http.StatusBadRequest,
					Detail:     "Could not unmarshall request JSON",
				}, w)
				return
			}

			// bank
			if req.Metric == "" || req.UnitType == "" {
				logger.Error().AnErr("addMetric()", errors.New("missing field(s)")).Msg("400 Bad Request")
				Report(ProblemDetail{
					StatusCode: http.StatusBadRequest,
					Detail:     "Metric, unitType, lower, or upper are empty",
				}, w)
				return
			}

			// exists -- too little time to validate dups
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
				return
			}

			// pattern matching would be nice here
			success := false
			if req.Patient == "" && req.Department == "" {
				hs.hospital.AddHospitalMetrics(req.Metric, req.UnitType, req.Lower, req.Upper)
				success = true
			}
			if req.Patient == "" && req.Department != "" {
				if _, ok := hs.hospital.Children[req.Department]; ok {
					hs.hospital.AddDepartmentMetrics(req.Department, req.Metric, req.UnitType, req.Lower, req.Upper)
					success = true
				}
			}
			if req.Patient != "" && req.Department != "" {
				if _, ok := hs.hospital.MetricKeys[model.MetricKey{
					Department: req.Department,
					Patient:    req.Patient,
					Metric:     req.Metric,
				}]; ok {
					hs.hospital.AddPatientMetric(req.Department, req.Patient, req.Metric, req.UnitType, req.Lower, req.Upper)
					success = true
				}
			}

			// catch all, nothing happened, doh
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
			jsonResp := fmt.Sprintf("{\"added\": \"%s\"}", req.Metric)
			if _, err := w.Write([]byte(jsonResp)); err != nil {
				logger.Error().AnErr("w.Write", err).Msg("500 Internal server error")
			} else {
				logger.Info().Msg("200 OK")
			}

		}
	}
}
