package main

import (
	"context"
	"flag"
	"log"
	"os/signal"
	"syscall"

	"github.com/saas-templates/go-tailwind/server"
)

var addr = flag.String("addr", ":8080", "Server bind address")

func main() {
	flag.Parse()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	if err := server.Serve(ctx, *addr); err != nil {
		log.Fatalf("[FATAL] server exited with error: %v", err)
	}
}
