package main

import (
	"log"
	"net/http"
)

type apiConfig struct {
	fileserverHits int
}

func Server() {
	const filepathRoot = "."
	const port = "8080"

	apiCfg := &apiConfig{
		fileserverHits: 0,
	}

	mux := http.NewServeMux()
	fsHandler := http.StripPrefix("/app/", http.FileServer(http.Dir(filepathRoot)))
	mux.Handle("/app/", apiCfg.middlewareMetricsInc(fsHandler))

	mux.HandleFunc("GET /api/healthz", handlerHealth)
	mux.HandleFunc("GET /api/metrics", apiCfg.handlerMetrics)
	mux.HandleFunc("GET /admin/metrics", apiCfg.handlerAdminMetrics)
	mux.HandleFunc("/api/reset", apiCfg.handlerReset)

	server := &http.Server{
		Addr:    "localhost:" + port,
		Handler: mux,
	}

	log.Printf("Serving on port: %s\n", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Could not listen on port %s\n", port)
	}
}
