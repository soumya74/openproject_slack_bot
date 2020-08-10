package handlers

import (
	"github.com/gorilla/mux"
)

func HandleRoutes(r *mux.Router) {
	r.HandleFunc("/webhooks", webhookHandler).Methods("POST")
	r.HandleFunc("/ping", healthTest)
}
