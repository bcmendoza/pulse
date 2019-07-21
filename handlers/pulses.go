package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type PulsesReqBody struct {
	Department string  `json:"department"`
	Patient    string  `json:"patient"`
	Metric     string  `json:"metric"`
	Value      float64 `json:"value"`
}

func (hs *handlersState) addMetricPulse() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if logger, ok := validateMethod("/pulses", r.Method, "POST", hs.logger, w); ok {

			var req PulsesReqBody
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

			if req.Department == "" || req.Patient == "" || req.Metric == "" {
				logger.Error().AnErr("addMetricPulse()", errors.New("missing field(s)")).Msg("400 Bad Request")
				Report(ProblemDetail{
					StatusCode: http.StatusBadRequest,
					Detail:     "Fields are empty",
				}, w)
			}

			hs.hospital.AddMetricPulse(req.Department, req.Patient, req.Metric, req.Value)
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			jsonResp := fmt.Sprintf("{\"added\": %f}", req.Value)
			if _, err := w.Write([]byte(jsonResp)); err != nil {
				logger.Error().AnErr("w.Write", err).Msg("500 Internal server error")
			} else {
				logger.Info().Msg("200 OK")
			}

		}
	}
}
