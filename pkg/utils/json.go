package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func ResponseWithError(w http.ResponseWriter, code int, msg string) {
	type errResponse struct {
		Error string `json:"error"`
	}
	ResponseWithJson(w, code, errResponse{
		Error: msg,
	})
}

func ResponseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Fail to Marshal Json response: %v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
