package main

import (
	"log"

	server "github.com/blankstatic/pdfcompose/pkg/rpc-server"
)

func main() {
	if err := server.RunRPCServer(); err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
}
