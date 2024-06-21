package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func (cfg *apiConfig) handlerValidateChirp(w http.ResponseWriter, req *http.Request) {
	var params struct {
		Body string `json:"body"`
	}

	if err := json.NewDecoder(req.Body).Decode(&params); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if len(params.Body) > 140 {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long")
		return
	}

	cleanedBody := censorProfaneWords(params.Body)
	response := map[string]string{"cleaned_body": cleanedBody}
	respondWithJSON(w, http.StatusOK, response)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func censorProfaneWords(text string) string {
	profaneList := []string{"kerfuffle", "sharbert", "fornax"}
	wordList := strings.Split(text, " ")
	var cleanedWords []string

	for _, word := range wordList {
		isProfane := false
		for _, profane := range profaneList {
			if strings.ToLower(word) == profane {
				cleanedWords = append(cleanedWords, "****")
				isProfane = true
			}
		}

		if !isProfane {
			cleanedWords = append(cleanedWords, word)
		}
	}

	result := strings.Join(cleanedWords, " ")
	return result
}
