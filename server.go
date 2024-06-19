package main

import (
	"fmt"
	"log"
	"net/http"
)

type apiConfig struct {
	fileserverHits int
}

func Server() {
	const filepathRoot = "."
	const port = "8080"

	fileServerHandler := http.StripPrefix("/app/", http.FileServer(http.Dir(filepathRoot)))

	apiCfg := &apiConfig{}

	mux := http.NewServeMux()
	mux.Handle("/app/", apiCfg.middlewareMetricsInc(fileServerHandler))
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

func handlerHealth(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}

func (cfg *apiConfig) handlerMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hits: " + fmt.Sprint(cfg.fileserverHits)))
}

func (cfg *apiConfig) handlerAdminMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	htmlTemplate := `
		<html>
			<body>
				<h1>Welcome, Chirpy Admin</h1>
				<p>Chirpy has been visited %d times!</p>
			</body>
		</html>
		`

	w.Write([]byte(fmt.Sprintf(htmlTemplate, cfg.fileserverHits)))
}

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
	cfg.fileserverHits = 0
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hits reset to 0"))
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits++
		next.ServeHTTP(w, r)
	})
}
