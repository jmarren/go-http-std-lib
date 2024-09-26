package main

import (
	"net/http"

	"github.com/jmarren/go-http-std-lib/internal"
)

func AddRoutes(r *http.ServeMux) {
	r.HandleFunc("/", internal.IndexHandler)
}
