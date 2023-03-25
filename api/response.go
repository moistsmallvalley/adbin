package api

import (
	"encoding/json"
	"net/http"

	"github.com/moistsmallvalley/adbin/log"
)

func writeOK(w http.ResponseWriter, res any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Warn(err.Error())
	}
}

func writeNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	writeErrorMessage(w, "not found")
}

func writeBadRequest(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusBadRequest)
	writeErrorMessage(w, msg)
}

func writeInternalServerError(w http.ResponseWriter, msg string, err error) {
	log.Error("%s: %+v", msg, err)
	w.WriteHeader(http.StatusInternalServerError)
	writeErrorMessage(w, msg)
}

func writeErrorMessage(w http.ResponseWriter, msg string) {
	if err := json.NewEncoder(w).Encode(map[string]string{"error": msg}); err != nil {
		log.Warn(err.Error())
	}
}
