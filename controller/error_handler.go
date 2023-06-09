package controller

import (
	"Tugas-Explorasi-1-PBP-Framework-API/model"
	"encoding/json"
	"log"
	"net/http"
)

func sendErrorResponse(w http.ResponseWriter, message string) {
	var response model.ErrorResponse
	response.Status = 400
	response.Message = message
	w.Header().Set("Content=Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println(err)
	}
}
