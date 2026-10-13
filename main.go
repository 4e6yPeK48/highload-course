package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	server := &http.Server{
		Addr:              ":8080",
		Handler:           handler(),
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       30 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}

func handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		io.WriteString(w, "ok\n")
	})
	return mux
}
