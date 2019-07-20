package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/rs/zerolog"
)

func validateMethod(route, method, expectedMethod string, logger zerolog.Logger, w http.ResponseWriter) (zerolog.Logger, bool) {
	logger = logger.With().Str("request-type", fmt.Sprintf("%s %s", method, route)).Logger()
	if method != expectedMethod {
		logger.Warn().Msg("405 Method Not Allowed")
		Report(ProblemDetail{StatusCode: http.StatusMethodNotAllowed, Detail: method}, w)
		return logger, false
	}
	logger.Info().Msg("Receive request")
	return logger, true
}

func validateRequestFields(reqBody io.ReadCloser, logger zerolog.Logger, w http.ResponseWriter) (RequestBody, bool) {
	var req RequestBody
	decoder := json.NewDecoder(reqBody)
	err := decoder.Decode(&req)
	if err != nil {
		logger.Error().AnErr("json.NewDecoder", err).Msg("400 Bad Request")
		Report(ProblemDetail{
			StatusCode: http.StatusBadRequest,
			Detail:     "Could not unmarshall request JSON",
		}, w)
		return RequestBody{}, false
	}
	return req, true
}
