package main

import (
	"github.com/YegorPedan/GOlangCleanApi/internal/server"
	"log"
)

func main() {
	s := server.NewServer(":4567")

	if err := s.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
