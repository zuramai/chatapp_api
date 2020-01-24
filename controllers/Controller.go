package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/zuramai/whatsapp_api/models"
)

func writeJSON(w http.ResponseWriter, status int, message string, data interface{}) {

	var payload models.Payload
	if status == 200 {
		payload.Status = true
	} else {
		payload.Status = false
	}
	payload.Message = message
	payload.Data = data

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
