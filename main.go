package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/osag1e/logstack/service/middleware"
)

func main() {
	router := chi.NewRouter()
	listenAddr := ":8080"

	logStack := middleware.LogStack(
		middleware.LogRequest,
		middleware.LogResponse,
		middleware.RecoverPanic,
	)

	if err := http.ListenAndServe(listenAddr, logStack(router)); err != nil {
		log.Fatal("HTTP server error:", err)
	}
}
