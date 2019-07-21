package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bcmendoza/pulse/handlers"
	"github.com/bcmendoza/pulse/model"

	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

func main() {
	var err error

	viper.AutomaticEnv()
	viper.SetDefault("PORT", 8080)

	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.Stamp}).
		With().Timestamp().Logger()
	logger.Info().Msg("Startup")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	logger.Info().Msg("Watch OS")

	// context
	_, cancelFunc := context.WithCancel(context.Background())

	hospital := model.New()

	// demo mode
	hospital.LoadTestSchemas()

	// REST server
	serverLogger := logger.With().Str("package", "handlers").Logger()
	server := http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", viper.GetInt("PORT")),
		Handler: handlers.Handlers(hospital, serverLogger),
	}
	go func() {
		serverLogger.Info().Msg("Startup REST server")
		if err = server.ListenAndServe(); err != nil && err.Error() != "http: Server closed" {
			serverLogger.Error().AnErr("server.ListenAndServe()", err).Msg("REST server error")
		}
	}()

	// shutdown
	s := <-sigChan
	cancelFunc()
	if err = server.Close(); err != nil {
		logger.Error().AnErr("server.Close()", err).Msg("REST server shutdown error")
	} else {
		logger.Info().Msg("Shutdown REST server")
	}
	logger.Info().Str("signal", s.String()).Msg("Shutdown")
}
