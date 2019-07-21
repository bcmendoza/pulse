package handlers

import (
	"errors"
	"fmt"
	"net/http"
)

func (hs *handlersState) addMetricPulse() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if logger, ok := validateMethod("/pulses", r.Method, "POST", hs.logger, w); ok {
			req, ok := validateRequestFields(r.Body, logger, w)
			if ok {
				if req.Department == "" || req.Patient == "" || req.Metric == "" || req.Value == 0 {
					logger.Error().AnErr("addMetricPulse()", errors.New("missing field(s)")).Msg("400 Bad Request")
					Report(ProblemDetail{
						StatusCode: http.StatusBadRequest,
						Detail:     "Fields are empty",
					}, w)
				}

				hs.hospital.AddMetricPulse(req.Department, req.Patient, req.Metric, req.Value)
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
}
