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
	demoChan chan<- struct{}
}

func Handlers(hospital *model.Hospital, logger zerolog.Logger) http.Handler {
	hs := handlersState{hospital: hospital, logger: logger}
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/streams", hs.getStreams())
	r.HandleFunc("/departments", hs.addDepartment())
	r.HandleFunc("/patients", hs.addPatient())
	r.HandleFunc("/metrics", hs.addMetric())
	r.HandleFunc("/pulses", hs.addMetricPulse())
	r.HandleFunc("/startDemo", hs.startDemo())
	r.HandleFunc("/stopDemo", hs.stopDemo())
	// r.HandleFunc("/doc").Handler(http.FileServer(http.Dir("/app/docs")))
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("/app/client")))
	return r
}

func (hs *handlersState) getStreams() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if logger, ok := validateMethod("/streams", r.Method, "GET", hs.logger, w); ok {
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

func (hs *handlersState) startDemo() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var jsonResp string
		if hs.demoChan != nil {
			demoChan := make(chan struct{}, 1)
			hs.demoChan = demoChan
			go hs.hospital.RunGenerator(demoChan)
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			jsonResp = "{\"started\": \"streaming for 5 minutes\"}"
		} else {
			jsonResp = "{\"already running\": \"hold yer horses pal\"}"
		}
		if _, err := w.Write([]byte(jsonResp)); err != nil {
			hs.logger.Error().AnErr("w.Write", err).Msg("500 Internal server error")
		} else {
			hs.logger.Info().Msg("200 OK")
		}
	}
}

func (hs *handlersState) stopDemo() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		hs.demoChan <- struct{}{}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		jsonResp := "{\"stopped\": \"no longer streaming\"}"
		if _, err := w.Write([]byte(jsonResp)); err != nil {
			hs.logger.Error().AnErr("w.Write", err).Msg("500 Internal server error")
		} else {
			hs.logger.Info().Msg("200 OK")
		}
	}
}
