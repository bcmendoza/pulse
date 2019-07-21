package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bcmendoza/pulse/model"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

type handlersState struct {
	hospital *model.Hospital
	logger   zerolog.Logger
}

func Handlers(hospital *model.Hospital, logger zerolog.Logger) http.Handler {
	hs := handlersState{hospital, logger}
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/streams", hs.getStreams())
	r.HandleFunc("/departments", hs.addDepartment())
	r.HandleFunc("/patients", hs.addPatient())
	r.HandleFunc("/metrics", hs.addMetric())
	r.HandleFunc("/pulses", hs.addMetricPulse())
	// r.PathPrefix("/").Handler(http.FileServer(http.Dir("/app/client")))
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("/app/docs")))
	return r
}

func (hs *handlersState) getStreams() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if logger, ok := validateMethod("/streams", r.Method, "GET", hs.logger, w); ok {
			logger.Info().Msg(fmt.Sprintf("%+v", hs.hospital))
			jsonResp, err := json.Marshal(hs.hospital)
			if err != nil {
				logger.Error().AnErr("json.Marshal", err).Msg("Could not marshall response into JSON")
				Report(ProblemDetail{
					StatusCode: http.StatusInternalServerError,
					Detail:     "Could not marshall response into JSON",
				}, w)
				return
			}

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
