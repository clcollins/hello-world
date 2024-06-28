package main

import (
	"fmt"
	"log/slog"
	"net/http"
)

var port = "8080"

func format(r *http.Request) []any {
	return []any{
		"remoteAddr", r.RemoteAddr,
		"client", "-",
		"username", "-",
		"request", fmt.Sprintf("%v %v %v", r.Method, r.URL.Path, r.Proto),
		"status", http.StatusOK,
		"bytes", r.ContentLength,
		"referer", r.Referer(),
		"userAgent", r.UserAgent(),
	}
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("received request", format(r)...)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte("Hello, world!"))
	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("received request", format(r)...)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"status\": \"ok\"}"))
	})

	http.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("received request", format(r)...)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"status\": \"ok\"}"))
	})

	slog.Info("Starting server", "port", port)
	http.ListenAndServe(":"+port, nil)

}
