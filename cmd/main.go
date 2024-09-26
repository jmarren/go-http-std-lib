package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
)

func run(ctx context.Context, w io.Writer, r io.Reader, envfunc func(string) string, args []string) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	mux := http.NewServeMux()

	AddRoutes(mux)

	httpServer := &http.Server{
		Addr:    "7071",
		Handler: mux,
	}

	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
	}

	fmt.Fprintf(os.Stdout, "listeing on port 7071\n")

	return nil
}

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Stdout, os.Stdin, os.Getenv, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
