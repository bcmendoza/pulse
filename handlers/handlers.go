package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/bcmendoza/pulse/model"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

type handlersState struct {
	hospital *model.Hospital
	logger   zerolog.Logger
}

type RequestBody struct {
	Department string `json:"department"`
	Patient    string `json:"patient"`
	Metric     string `json:"metric"`
	UnitType   string `json:"unitType"`
}

func Handlers(hospital *model.Hospital, logger zerolog.Logger) http.Handler {
	hs := handlersState{hospital, logger}
	r := mux.NewRouter()
	r.HandleFunc("/summary", hs.getSummary())
	r.HandleFunc("/departments", hs.addDepartment())
	r.HandleFunc("/patients", hs.addPatient())
	r.HandleFunc("/metrics", hs.addMetric())
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("/app/docs")))
	return r
}

func (hs *handlersState) getSummary() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if logger, ok := validateMethod("/summary", r.Method, "GET", hs.logger, w); ok {
			jsonResp, err := json.Marshal(hs.hospital)
			if err != nil {
				logger.Error().AnErr("json.Marshal", err).Msg("Could not marshall response into JSON")
				Report(ProblemDetail{
					StatusCode: http.StatusInternalServerError,
					Detail:     "Could not marshall response into JSON",
				}, w)
			} else {
				w.WriteHeader(http.StatusOK)
				w.Header().Set("Content-Type", "application/json")
				if _, err := w.Write([]byte(jsonResp)); err != nil {
					logger.Error().AnErr("w.Write", err).Msg("500 Internal server error")
				} else {
					logger.Info().Msg("200 OK")
				}
			}
		}
	}
}
