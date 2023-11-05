package http

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"
)

type Server struct {
	Server *http.Server
	Router *Router
}

func NewServer(router *Router, port string) *Server {
	engine := &http.Server{
		Addr:              port,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           router.engine,
	}

	return &Server{
		Router: router,
		Server: engine,
	}
}

func (s *Server) Init() error {
	go func() {
		log.Info().Msgf("Server is up and listening on port: %s", s.Server.Addr)
		if err := s.Server.ListenAndServe(); err != nil {
			log.Error().Msgf("Server failed to start: %v", err)
		}
	}()
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Server.Shutdown(ctx); err != nil {
		log.Error().Msgf("Server forced to shutdown: %v", err)
		return err
	}

	log.Info().Msg("Server exiting ")
	return nil
}
