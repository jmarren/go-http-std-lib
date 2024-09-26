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

	mux := http.NewServeMux()

	defer cancel()

	return nil
}

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Stdout, os.Stdin, os.Getenv, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
