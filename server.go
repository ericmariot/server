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

	log.Printf("Serving on port: %s\n", port)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Could not listen on port %s\n", port)
	}
}
