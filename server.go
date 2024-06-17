package main

import (
	"log"
	"net/http"
)

func Server() {
	const port = "8080"

	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    "localhost:" + port,
		Handler: mux,
	}
	mux.Handle("/", http.FileServer(http.Dir(".")))

	log.Printf("Serving on port: %s\n", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Could not listen on port %s\n", port)
	}
}
