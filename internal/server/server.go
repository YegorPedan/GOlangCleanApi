package server

import (
	"context"
	"github.com/labstack/echo/v4"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const SecondsToStopServer = 10

type Server struct {
	addr string
	app  *echo.Echo
}

func NewServer(addr string) *Server {
	return &Server{
		addr: addr,
		app:  echo.New(),
	}
}

func (s *Server) Run() error {
	// Start server in a goroutine
	go func() {
		log.Printf("Starting server on %s", s.addr)
		if err := s.app.Start(s.addr); err != nil {
			log.Fatalf("Server start failed: %v", err)
		}
	}()

	if err := s.MapHandlers(s.app); err != nil {
		log.Fatalf("An error has occured maping the handlers: %v", err)
	}

	// Handle graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	log.Println("Shutting down server...")

	// Server has 10 second to shut down all process
	ctx, cancel := context.WithTimeout(context.Background(), SecondsToStopServer*time.Second)
	defer cancel()

	if err := s.app.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown error: %v", err)
		return err
	}

	log.Println("Server stopped gracefully")
	return nil
}
