package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type params struct {
	Body string `json:"body"`
}

type errorReturn struct {
	Error string `json:"error"`
}

type validReturn struct {
	Valid bool `json:"valid"`
}

func (cfg *apiConfig) handlerValidateChirp(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	p := params{}
	err := decoder.Decode(&p)
	if err != nil {
		log.Printf("Error decoding parameters: %s", err)
		w.WriteHeader(500)
		return
	}

	if len(p.Body) > 140 {
		respBody := errorReturn{
			Error: "Chirp is too long",
		}
		dat, err := json.Marshal(respBody)
		if err != nil {
			log.Printf("Error marshalling JSON: %s", err)
			w.WriteHeader(500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		w.Write(dat)
		return
	}

	respBody := validReturn{
		Valid: true,
	}
	dat, err := json.Marshal(respBody)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(dat)
}
